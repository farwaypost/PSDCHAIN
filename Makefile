# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gpch android ios gpch-cross swarm evm all test clean
.PHONY: gpch-linux gpch-linux-386 gpch-linux-amd64 gpch-linux-mips64 gpch-linux-mips64le
.PHONY: gpch-linux-arm gpch-linux-arm-5 gpch-linux-arm-6 gpch-linux-arm-7 gpch-linux-arm64
.PHONY: gpch-darwin gpch-darwin-386 gpch-darwin-amd64
.PHONY: gpch-windows gpch-windows-386 gpch-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gpch:
	build/env.sh go run build/ci.go install ./cmd/gpch
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gpch\" to launch gpch."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gpch.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Gpch.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

lint: ## Run linters.
	build/env.sh go run build/ci.go lint

clean:
	./build/clean_go_build_cache.sh
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

swarm-devtools:
	env GOBIN= go install ./cmd/swarm/mimegen

# Cross Compilation Targets (xgo)

gpch-cross: gpch-linux gpch-darwin gpch-windows gpch-android gpch-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gpch-*

gpch-linux: gpch-linux-386 gpch-linux-amd64 gpch-linux-arm gpch-linux-mips64 gpch-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-*

gpch-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gpch
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep 386

gpch-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gpch
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep amd64

gpch-linux-arm: gpch-linux-arm-5 gpch-linux-arm-6 gpch-linux-arm-7 gpch-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep arm

gpch-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gpch
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep arm-5

gpch-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gpch
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep arm-6

gpch-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gpch
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep arm-7

gpch-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gpch
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep arm64

gpch-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gpch
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep mips

gpch-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gpch
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep mipsle

gpch-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gpch
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep mips64

gpch-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gpch
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gpch-linux-* | grep mips64le

gpch-darwin: gpch-darwin-386 gpch-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gpch-darwin-*

gpch-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gpch
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-darwin-* | grep 386

gpch-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gpch
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-darwin-* | grep amd64

gpch-windows: gpch-windows-386 gpch-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gpch-windows-*

gpch-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gpch
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-windows-* | grep 386

gpch-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gpch
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gpch-windows-* | grep amd64
