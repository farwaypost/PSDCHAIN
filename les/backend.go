// Copyright 2016 The go-psdchaineum Authors
// This file is part of the go-psdchaineum library.
//
// The go-psdchaineum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-psdchaineum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-psdchaineum library. If not, see <http://www.gnu.org/licenses/>.

// Package les implements the Light Psdchain Subprotocol.
package les

import (
	"fmt"
	"sync"
	"time"

	"github.com/psdchaineum/go-psdchaineum/accounts"
	"github.com/psdchaineum/go-psdchaineum/common"
	"github.com/psdchaineum/go-psdchaineum/common/hexutil"
	"github.com/psdchaineum/go-psdchaineum/consensus"
	"github.com/psdchaineum/go-psdchaineum/core"
	"github.com/psdchaineum/go-psdchaineum/core/bloombits"
	"github.com/psdchaineum/go-psdchaineum/core/rawdb"
	"github.com/psdchaineum/go-psdchaineum/core/types"
	"github.com/psdchaineum/go-psdchaineum/pch"
	"github.com/psdchaineum/go-psdchaineum/pch/downloader"
	"github.com/psdchaineum/go-psdchaineum/pch/filters"
	"github.com/psdchaineum/go-psdchaineum/pch/gasprice"
	"github.com/psdchaineum/go-psdchaineum/event"
	"github.com/psdchaineum/go-psdchaineum/internal/pchapi"
	"github.com/psdchaineum/go-psdchaineum/light"
	"github.com/psdchaineum/go-psdchaineum/log"
	"github.com/psdchaineum/go-psdchaineum/node"
	"github.com/psdchaineum/go-psdchaineum/p2p"
	"github.com/psdchaineum/go-psdchaineum/p2p/discv5"
	"github.com/psdchaineum/go-psdchaineum/params"
	rpc "github.com/psdchaineum/go-psdchaineum/rpc"
)

type LightPsdchain struct {
	lesCommons

	odr         *LesOdr
	relay       *LesTxRelay
	chainConfig *params.ChainConfig
	// Channel for shutting down the service
	shutdownChan chan bool

	// Handlers
	peers      *peerSet
	txPool     *light.TxPool
	blockchain *light.LightChain
	serverPool *serverPool
	reqDist    *requestDistributor
	retriever  *retrieveManager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer

	ApiBackend *LesApiBackend

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	networkId     uint64
	netRPCService *pchapi.PublicNetAPI

	wg sync.WaitGroup
}

func New(ctx *node.ServiceContext, config *pch.Config) (*LightPsdchain, error) {
	chainDb, err := pch.CreateDB(ctx, config, "lightchaindata")
	if err != nil {
		return nil, err
	}
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlockWithOverride(chainDb, config.Genesis, config.ConstantinopleOverride)
	if _, isCompat := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !isCompat {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	peers := newPeerSet()
	quitSync := make(chan struct{})

	lpch := &LightPsdchain{
		lesCommons: lesCommons{
			chainDb: chainDb,
			config:  config,
			iConfig: light.DefaultClientIndexerConfig,
		},
		chainConfig:    chainConfig,
		eventMux:       ctx.EventMux,
		peers:          peers,
		reqDist:        newRequestDistributor(peers, quitSync),
		accountManager: ctx.AccountManager,
		engine:         pch.CreateConsensusEngine(ctx, chainConfig, &config.Ethash, nil, false, chainDb),
		shutdownChan:   make(chan bool),
		networkId:      config.NetworkId,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   pch.NewBloomIndexer(chainDb, params.BloomBitsBlocksClient, params.HelperTrieConfirmations),
	}

	lpch.relay = NewLesTxRelay(peers, lpch.reqDist)
	lpch.serverPool = newServerPool(chainDb, quitSync, &lpch.wg)
	lpch.retriever = newRetrieveManager(peers, lpch.reqDist, lpch.serverPool)

	lpch.odr = NewLesOdr(chainDb, light.DefaultClientIndexerConfig, lpch.retriever)
	lpch.chtIndexer = light.NewChtIndexer(chainDb, lpch.odr, params.CHTFrequencyClient, params.HelperTrieConfirmations)
	lpch.bloomTrieIndexer = light.NewBloomTrieIndexer(chainDb, lpch.odr, params.BloomBitsBlocksClient, params.BloomTrieFrequency)
	lpch.odr.SetIndexers(lpch.chtIndexer, lpch.bloomTrieIndexer, lpch.bloomIndexer)

	// Note: NewLightChain adds the trusted checkpoint so it needs an ODR with
	// indexers already set but not started yet
	if lpch.blockchain, err = light.NewLightChain(lpch.odr, lpch.chainConfig, lpch.engine); err != nil {
		return nil, err
	}
	// Note: AddChildIndexer starts the update process for the child
	lpch.bloomIndexer.AddChildIndexer(lpch.bloomTrieIndexer)
	lpch.chtIndexer.Start(lpch.blockchain)
	lpch.bloomIndexer.Start(lpch.blockchain)

	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		lpch.blockchain.SetHead(compat.RewindTo)
		rawdb.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}

	lpch.txPool = light.NewTxPool(lpch.chainConfig, lpch.blockchain, lpch.relay)
	if lpch.protocolManager, err = NewProtocolManager(lpch.chainConfig, light.DefaultClientIndexerConfig, true, config.NetworkId, lpch.eventMux, lpch.engine, lpch.peers, lpch.blockchain, nil, chainDb, lpch.odr, lpch.relay, lpch.serverPool, quitSync, &lpch.wg); err != nil {
		return nil, err
	}
	lpch.ApiBackend = &LesApiBackend{lpch, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.MinerGasPrice
	}
	lpch.ApiBackend.gpo = gasprice.NewOracle(lpch.ApiBackend, gpoParams)
	return lpch, nil
}

func lesTopic(genesisHash common.Hash, protocolVersion uint) discv5.Topic {
	var name string
	switch protocolVersion {
	case lpv1:
		name = "LES"
	case lpv2:
		name = "LES2"
	default:
		panic(nil)
	}
	return discv5.Topic(name + "@" + common.Bytes2Hex(genesisHash.Bytes()[0:8]))
}

type LightDummyAPI struct{}

// Psdchainbase is the address that mining rewards will be send to
func (s *LightDummyAPI) Psdchainbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

// Coinbase is the address that mining rewards will be send to (alias for Psdchainbase)
func (s *LightDummyAPI) Coinbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

// Hashrate returns the POW hashrate
func (s *LightDummyAPI) Hashrate() hexutil.Uint {
	return 0
}

// Mining returns an indication if this node is currently mining.
func (s *LightDummyAPI) Mining() bool {
	return false
}

// APIs returns the collection of RPC services the psdchaineum package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *LightPsdchain) APIs() []rpc.API {
	return append(pchapi.GetAPIs(s.ApiBackend), []rpc.API{
		{
			Namespace: "pch",
			Version:   "1.0",
			Service:   &LightDummyAPI{},
			Public:    true,
		}, {
			Namespace: "pch",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "pch",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.ApiBackend, true),
			Public:    true,
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
	}...)
}

func (s *LightPsdchain) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *LightPsdchain) BlockChain() *light.LightChain      { return s.blockchain }
func (s *LightPsdchain) TxPool() *light.TxPool              { return s.txPool }
func (s *LightPsdchain) Engine() consensus.Engine           { return s.engine }
func (s *LightPsdchain) LesVersion() int                    { return int(ClientProtocolVersions[0]) }
func (s *LightPsdchain) Downloader() *downloader.Downloader { return s.protocolManager.downloader }
func (s *LightPsdchain) EventMux() *event.TypeMux           { return s.eventMux }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *LightPsdchain) Protocols() []p2p.Protocol {
	return s.makeProtocols(ClientProtocolVersions)
}

// Start implements node.Service, starting all internal goroutines needed by the
// Psdchain protocol implementation.
func (s *LightPsdchain) Start(srvr *p2p.Server) error {
	log.Warn("Light client mode is an experimental feature")
	s.startBloomHandlers(params.BloomBitsBlocksClient)
	s.netRPCService = pchapi.NewPublicNetAPI(srvr, s.networkId)
	// clients are searching for the first advertised protocol in the list
	protocolVersion := AdvertiseProtocolVersions[0]
	s.serverPool.start(srvr, lesTopic(s.blockchain.Genesis().Hash(), protocolVersion))
	s.protocolManager.Start(s.config.LightPeers)
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Psdchain protocol.
func (s *LightPsdchain) Stop() error {
	s.odr.Stop()
	s.bloomIndexer.Close()
	s.chtIndexer.Close()
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()
	s.engine.Close()

	s.eventMux.Stop()

	time.Sleep(time.Millisecond * 200)
	s.chainDb.Close()
	close(s.shutdownChan)

	return nil
}
