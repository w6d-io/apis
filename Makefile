export GO111MODULE  := on
export PATH         := ./bin:${PATH}
export NEXT_TAG     ?=

ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

ifeq (,$(shell go env GOOS))
GOOS       = $(shell echo $OS)
else
GOOS       = $(shell go env GOOS)
endif

ifeq (,$(shell go env GOARCH))
GOARCH     = $(shell echo uname -p)
else
GOARCH     = $(shell go env GOARCH)
endif

PROTOC = $(shell pwd)/bin/protoc
ifeq (darwin,$(GOOS))
PROTOC_ZIP=https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-osx-x86_64.zip
else
PROTOC_ZIP=https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip
endif

export PATH := $(shell pwd)/bin:${PATH}

all: proto

##@ General

REF        = $(shell git symbolic-ref --quiet HEAD 2> /dev/null)
VERSION   ?= $(shell basename /$(shell git symbolic-ref --quiet HEAD 2> /dev/null ) )
VCS_REF    = $(shell git rev-parse HEAD)
GOVERSION  = $(shell go env GOVERSION)
BUILD_DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: protobuf
protobuf: bin/protoc bin/protoc-gen-go bin/protoc-gen-go-grpc bin/protoc-gen-openapiv2 bin/protoc-gen-doc

bin/protoc: ## install protoc locally if necessary.
	$(call install,$(PROTOC),"bin/protoc",$(PROTOC_ZIP))


bin/protoc-gen-go:
	$(call go-get-tool,bin/protoc-gen-go,google.golang.org/protobuf/cmd/protoc-gen-go)

bin/protoc-gen-go-grpc:
	$(call go-get-tool,bin/protoc-gen-go-grpc,google.golang.org/grpc/cmd/protoc-gen-go-grpc)

bin/protoc-gen-doc:
	$(call go-get-tool,bin/protoc-gen-doc,github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc)

project/v1alpha1/project.pb.go project/v1alpha1/project_grpc.pb.go: bin/protoc bin/protoc-gen-go bin/protoc-gen-go-grpc bin/protoc-gen-doc proto/project.proto
	$(PROTOC) -Iproto --go_out=. --go_opt=module=github.com/w6d-io/apis --go-grpc_opt=module=github.com/w6d-io/apis --go-grpc_out=. --doc_opt=docs/templates/grpc-md.tmpl,project.md --doc_out=docs/apis proto/project.proto

.PHONY: proto
proto: project/v1alpha1/project.pb.go project/v1alpha1/project_grpc.pb.go

.PHONY: changelog
changelog:
	git-chglog -o docs/CHANGELOG.md --next-tag $(NEXT_TAG)

define install
@[ -f $(1) ] || { \
set -e;\
TMP_DIR=$$(mktemp -d);\
cd $$TMP_DIR ;\
wget -q $(3);\
unzip *.zip $(2);\
mv $(2) $(1);\
rm -rf $$TMP_DIR ;\
}
endef
