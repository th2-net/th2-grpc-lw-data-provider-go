COMMON_SRC_MAIN_PROTO_DIR=src/main/proto
LWDP_SRC_MAIN_PROTO_DIR=grpc/src/main/proto
GITHUB_TH2=github.com/th2-net

TH2_GRPC_COMMON_VERSION=dev-version-4 # TODO: replace to a tag after solving https://nvd.nist.gov/vuln/detail/CVE-2025-24970
TH2_GRPC_COMMON=th2-grpc-common
TH2_GRPC_COMMON_URL=$(GITHUB_TH2)/$(TH2_GRPC_COMMON)@$(TH2_GRPC_COMMON_VERSION)
LW_DATA_PROVIDER_VERSION_VARIABLE=dev-version-2 # TODO: replace to a tag after release
TH2_LW_DATA_PROVIDER=th2-lw-data-provider
TH2_LW_DATA_PROVIDER_URL=$(GITHUB_TH2)/$(TH2_LW_DATA_PROVIDER)@$(LW_DATA_PROVIDER_VERSION_VARIABLE)

PROTOC_GEN_GO_VERSION=v1.36.5
PROTOC_GEN_GO_GRPC_VERSION=v1.5.1

PROTOC_DIR=.protoc
PROTOC_BASE_URL=https://github.com/protocolbuffers/protobuf/releases/download
PROTOC_VERSION=29.3
PROTOC_OS=$(shell uname -s)
PROTOC_ARCH=$(shell uname -m)
PROTOC_FILE=protoc-$(PROTOC_VERSION)-$(PROTOC_OS)-$(PROTOC_ARCH).zip
PROTOC_URL=$(PROTOC_BASE_URL)/v$(PROTOC_VERSION)/$(PROTOC_FILE)

init-work-space: install-protoc generate-grpc-files tidy

install-protoc:
	- rm -rf $(PROTOC_DIR)
	- mkdir $(PROTOC_DIR)

	wget -P $(PROTOC_DIR) $(PROTOC_URL)
	unzip $(PROTOC_DIR)/$(PROTOC_FILE) -d $(PROTOC_DIR)
	$(PROTOC_DIR)/bin/protoc --version

configure-grpc-generator:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

prepare-grpc-module:
	go get $(TH2_GRPC_COMMON_URL)
	go get $(TH2_LW_DATA_PROVIDER_URL)
	go get google.golang.org/protobuf@$(PROTOC_GEN_GO_VERSION)
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

generate-grpc-files: configure-grpc-generator prepare-grpc-module tidy
	$(eval $@_COMMON_PROTO_DIR := $(shell go list -m -f '{{.Dir}}' $(TH2_GRPC_COMMON_URL))/$(COMMON_SRC_MAIN_PROTO_DIR))
	$(eval $@_LWDP_PROTO_DIR := $(shell go list -m -f '{{.Dir}}' $(TH2_LW_DATA_PROVIDER_URL))/$(LWDP_SRC_MAIN_PROTO_DIR))
	$(PROTOC_DIR)/bin/protoc \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=module=github.com/th2-net/th2-grpc-lw-data-provider-go \
		--go-grpc_opt=module=github.com/th2-net/th2-grpc-lw-data-provider-go \
		--proto_path=$($@_COMMON_PROTO_DIR) \
		--proto_path=$($@_LWDP_PROTO_DIR) \
		$(shell find $($@_LWDP_PROTO_DIR) -name '*.proto' )

check-grpc-files: generate-grpc-files
	@git diff --exit-code ./*.go || { echo "Committed Protobuf files do not match the newly generated. Please, regenerate Protobuf and commit changes"; exit 1; }

tidy:
	go mod tidy -v

build: 
	go vet ./...
	go build -v ./...

run-test:
	go test -v -race ./...