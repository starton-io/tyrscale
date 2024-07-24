# Default service name
SERVICE?=manager

# Docker registry
REGISTRY?=registry.gitlab.com/starton/blockchain/tyrscale

# Docker image tag
TAG?=latest
VERSION?=0.0.1

# Go build environment variables
GO_BUILD_ENV?=CGO_ENABLED=1
GOOS?=darwin
GOARCH?=arm64
ENV?=local
GO_MOD_TIDY ?= GO111MODULE=on go mod tidy
LDFLAGS=-ldflags "-s -w -X main.version=$(TAG)"

GO_LINT?=golangci-lint
GO_LINT_CONFIG?=.golangci.yml

SWITCH_DEPLOY_CONTEXT=cd ./deploy

# Test configuration
GO_TEST?=go test
GO_TEST_OPTS?=-v -cover
GO_TEST_COVERPROFILE ?= export RUNNING_MOD="test"; go test -v 2>&1 -coverprofile

.PHONY: run destroy build docker-build test clean help


install-deps:
	cd ./$(SERVICE) && $(GO_MOD_TIDY)

# Compile the application binary for a specific service
build:
	mkdir -p ./$(SERVICE)/bin
	$(GO_BUILD_ENV) GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o ./$(SERVICE)/bin/$(SERVICE) ./$(SERVICE)/cmd

# Build the Docker image for a specific service
docker-build:
	docker build -t $(REGISTRY)/$(SERVICE):$(TAG) -f ./$(SERVICE)/Dockerfile .

generate-proto:
	protoc --go_out ./manager/pkg/pb/network --go-grpc_out ./manager/pkg/pb/network ./manager/proto/network.proto
	protoc --go_out ./manager/pkg/pb/recommendation --go-grpc_out=require_unimplemented_servers=false:./manager/pkg/pb/recommendation ./manager/proto/recommendation.proto
	protoc --go_out ./manager/pkg/pb/rpc --go-grpc_out=require_unimplemented_servers=false:./manager/pkg/pb/rpc ./manager/proto/rpc.proto
	protoc --go_out ./manager/pkg/pb/route --go-grpc_out=require_unimplemented_servers=false:./manager/pkg/pb/route ./manager/proto/route.proto
	protoc --go_out ./manager/pkg/pb/upstream --go-grpc_out=require_unimplemented_servers=false:./manager/pkg/pb/upstream ./manager/proto/upstream.proto
	protoc-go-inject-tag -input=./manager/pkg/pb/rpc/rpc.pb.go
#protoc --go_out ./proto/gen/go/plugin --go-grpc_out ./proto/gen/go/plugin ./proto/plugin/plugin.proto

generate-proto-gateway:
	mkdir -p ./gateway/proto/gen/go/plugin
	protoc --go_out ./gateway/proto/gen/go/plugin --go-grpc_out=./gateway/proto/gen/go/plugin ./gateway/proto/plugin/plugin.proto


# Generate swagger docs
# and replace the package name in swagger.yaml/swagger.json with sed
#sed -i '' "s/github_com_starton-io_tyrscale_manager_api_//g" ./manager/docs/swagger.{yaml,json}
#sed -i '' "s/github_com_starton-io_tyrscale_manager_pkg_pb_//g" ./manager/docs/swagger.{yaml,json}

generate-all: generate-swagger generate-openapi generate-sdk-go generate-sdk-typescript

generate-swagger:
	cd ./manager && swag fmt && swag init --pdl 3 -g ./cmd/main.go  -o ./docs
	sed -i '' "s/github[^.]*starton*\.//g" ./manager/docs/{swagger.yaml,swagger.json,docs.go}
	sed -i '' "s/github_com_starton-io_[a-zA-Z0-9_]*_dto_//g" ./manager/docs/{swagger.yaml,swagger.json,docs.go}
	sed -i '' "s/github_com_starton-io_[a-zA-Z0-9_]*_dto\.//g" ./manager/docs/{swagger.yaml,swagger.json,docs.go}
	sed -i '' "s/github_com_starton-io_[a-zA-Z0-9_]*_pb_[a-zA-Z0-9]*\.//g" ./manager/docs/{swagger.yaml,swagger.json,docs.go}

generate-openapi:
	openapi-generator generate -p outputFile=openapi.yaml -i ./manager/docs/swagger.yaml -g openapi-yaml -o ./manager/docs/v3

generate-sdk-go:
	cd ./sdk/tyrscale-sdk-go && openapi-generator generate -i ../../manager/docs/v3/openapi.yaml -g go --additional-properties packageName=tyrscalesdkgo,packageVersion=$(VERSION),useTags=true,generateInterfaces=true --git-user-id starton-io --git-repo-id tyrscale/sdk/tyrscale-sdk-go

generate-sdk-typescript:
	cd ./sdk/tyrscale-sdk-typescript && openapi-generator generate -i ../../manager/docs/v3/openapi.yaml -g typescript-fetch --additional-properties=supportsES6=true,npmName="@starton/tyrscale-sdk-typescript",npmVersion=$(VERSION) --openapi-normalizer REFACTOR_ALLOF_WITH_PROPERTIES_ONLY=true --git-user-id starton-io --git-repo-id tyrscale/sdk/tyrscale-sdk-typescript

install:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/favadi/protoc-go-inject-tag@latest
	brew install openapi-generator
	brew install protobuf

run:
	${SWITCH_DEPLOY_CONTEXT} && docker-compose up --build

# Run tests for a specific service

test:
ifeq ($(ENV), local)
	cd ./$(SERVICE) && ${GO_MOD_TIDY}
	${GO_TEST_COVERPROFILE} coverage.txt  ./$(SERVICE)/...
else ifeq ($(ENV), ci)
	cd ./$(SERVICE) && ${GO_MOD_TIDY}
	go install github.com/jstemmer/go-junit-report@latest
	${GO_TEST_COVERPROFILE} ${CI_PROJECT_DIR}/${ARTIFACTS_DIR}/$(SERVICE)/.testCoverage.out  ./$(SERVICE)/... > ${CI_PROJECT_DIR}/${ARTIFACTS_DIR}/$(SERVICE)/tests.out
else
	$(error ENV ${ENV} not found)	
endif


test-lint:
	cd ./$(SERVICE) && ${GO_LINT} run -c ${GO_LINT_CONFIG}

test-report:
ifeq ($(ENV), ci)
	cat  ${CI_PROJECT_DIR}/${ARTIFACTS_DIR}/$(SERVICE)/tests.out | ${GOPATH}/bin/go-junit-report > ${CI_PROJECT_DIR}/${ARTIFACTS_DIR}/$(SERVICE)/report.xml
	${GO_TOOL_COVER} -func ${CI_PROJECT_DIR}/${ARTIFACTS_DIR}/$(SERVICE)/.testCoverage.out
	${GO_TOOL_COVER} -html=${CI_PROJECT_DIR}/${ARTIFACTS_DIR}/$(SERVICE)/.testCoverage.out -o ${CI_PROJECT_DIR}/${ARTIFACTS_DIR}/$(SERVICE)/index.html
endif

#test:
#	$(GO_TEST) $(GO_TEST_OPTS) ./$(SERVICE)/...

# Clean up the compiled binary and other artifacts for a specific service
clean:
	rm -f ./$(SERVICE)/bin/$(SERVICE)

destroy:
	${SWITCH_DEPLOY_CONTEXT} && docker-compose down

# Display help information
help:
	@echo "Available commands:"
	@echo "  build               - Compile the application binary for a specific service."
	@echo "  docker-build        - Build the Docker image for a specific service."
	@echo "  run                 - Start the service using docker-compose."
	@echo "  test                - Run tests for a specific service."
	@echo "  clean               - Clean up binaries and other artifacts for a specific service."
	@echo "  destroy             - Stop and remove containers, networks, images, and volumes."
	@echo "  generate-proto      - Generate protobuf files for the service."
	@echo "  generate-swagger    - Generate Swagger documentation for the service."
	@echo "  generate-openapi    - Generate OpenAPI specifications for the service."
	@echo "  generate-sdk-go     - Generate Go SDK using OpenAPI specs."
	@echo "  generate-sdk-typescript - Generate TypeScript SDK using OpenAPI specs."
	@echo "  install             - Install necessary tools for the project."
	@echo "Usage:"
	@echo "  make <command> [SERVICE=name] [REGISTRY=your-docker-repo] [TAG=tag] [VERSION=version]"