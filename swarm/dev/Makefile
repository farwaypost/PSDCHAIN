.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/psdchaineum/go-psdchaineum/cmd/swarm
	go build -o bin/gpch     github.com/psdchaineum/go-psdchaineum/cmd/gpch
	go build -o bin/bootnode github.com/psdchaineum/go-psdchaineum/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/psdchaineum/go-psdchaineum/swarm/...
