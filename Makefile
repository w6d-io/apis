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

ifeq (darwin,$(GOOS))
PROTOC_ZIP=https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-osx-x86_64.zip
else
PROTOC_ZIP=https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip
endif

export PATH := $(shell pwd)/bin:${PATH}

all: proto generate

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

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: fmt vet
	go test $(GOTAGS) -v -coverpkg=./... -coverprofile=cover.out ./...
	@go tool cover -func cover.out | grep total

.PHONY: protobuf
protobuf: protoc protoc-gen-go protoc-gen-go-grpc protoc-gen-doc #bin/protoc-gen-openapiv2

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
PROTOC             ?= $(LOCALBIN)/protoc
PROTOC_GEN_GO      ?= $(LOCALBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC ?= $(LOCALBIN)/protoc-gen-go-grpc
PROTOC_GEN_DOC     ?= $(LOCALBIN)/protoc-gen-doc
CONTROLLER_GEN     ?= $(LOCALBIN)/controller-gen
GITCHGLOG          ?= $(LOCALBIN)/git-chglog

## Tool Versions
CONTROLLER_TOOLS_VERSION ?= v0.9.2

.PHONY: protoc
protoc: $(PROTOC)
$(PROTOC): $(LOCALBIN) ## install protoc locally if necessary.
	@test -s $(LOCALBIN)/protoc || $(call install,$(PROTOC),$(PROTOC),$(PROTOC_ZIP))

.PHONY: protoc-gen-go
protoc-gen-go: $(PROTOC_GEN_GO)
$(PROTOC_GEN_GO): $(LOCALBIN)
	$(call go-get-tool,$(PROTOC_GEN_GO),google.golang.org/protobuf/cmd/protoc-gen-go)

.PHONY: protoc-gen-go-grpc
protoc-gen-go-grpc: $(PROTOC_GEN_GO_GRPC)
$(PROTOC_GEN_GO_GRPC): $(LOCALBIN)
	$(call go-get-tool,$(PROTOC_GEN_GO_GRPC),google.golang.org/grpc/cmd/protoc-gen-go-grpc)

.PHONY: protoc-gen-doc
protoc-gen-doc: $(PROTOC_GEN_DOC)
$(PROTOC_GEN_DOC): $(LOCALBIN)
	$(call go-get-tool,$(PROTOC_GEN_DOC),github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc)

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

.PHONY: chglog
chglog: $(GITCHGLOG) ## Download git-chglog locally if necessary
$(GITCHGLOG): $(LOCALBIN)
	test -s $(LOCALBIN)/git-chglog || GOBIN=$(LOCALBIN) go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

## genarate proto file
.PHONY: project
project: protobuf project/v1alpha1/project.pb.go project/v1alpha1/project_grpc.pb.go
project/v1alpha1/project.pb.go project/v1alpha1/project_grpc.pb.go: proto/project.proto
	$(PROTOC) -Iproto --go_out=. --go_opt=module=github.com/w6d-io/apis --go-grpc_opt=module=github.com/w6d-io/apis --go-grpc_out=. --doc_opt=.config/templates/grpc-md.tmpl,project.md --doc_out=docs/apis proto/project.proto

.PHONY: authz
authz: protobuf authz/v1alpha1/authz.pb.go #authz/v1alpha1/authz_grpc.pb.go
authz/v1alpha1/authz.pb.go authz/v1alpha1/authz_grpc.pb.go: proto/authz.proto
	$(PROTOC) -Iproto --go_out=. --go_opt=module=github.com/w6d-io/apis --go-grpc_opt=module=github.com/w6d-io/apis --go-grpc_out=. --doc_opt=.config/templates/grpc-md.tmpl,authz.md --doc_out=docs/apis proto/authz.proto

# Changelog
.PHONY: changelog
changelog: chglog
	$(GITCHGLOG) -o docs/CHANGELOG.md --next-tag $(NEXT_TAG)

##@ Build

.PHONY: proto
proto: project authz

##@ Script

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
