PROJECT_NAME := "go-rest-api"

GIT_COMMIT := $(shell git rev-parse --short HEAD)

CURR_DIR := $(shell pwd)

.PHONY: all build production-build help

all: build

lib-update: ## Build the binary file

build: ## Build the binary file
	@go build -ldflags="-X 'github.com/deepakhcu07/go-rest-api/api.BuildTime=$(TIMESTAMP)' -X 'github.com/deepakhcu07/go-rest-api/api.BuildVersion=$(GIT_COMMIT)'" -o $(CURR_DIR)/bin/$(PROJECT_NAME).o

run: ## Build the binary file
	@go build -o $(CURR_DIR)/bin/$(PROJECT_NAME).o
	./bin/$(PROJECT_NAME).o

production-build: ## Build in the Production Mode for Linux
	@GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/deepakhcu07/go-rest-api/api.BuildTime=$(TIMESTAMP)' -X 'github.com/deepakhcu07/go-rest-api/api.BuildVersion=$(GIT_COMMIT)'" -o $(CURR_DIR)/bin/$(PROJECT_NAME).o


alpine-build: ## Build in the Production Mode for Alpine Linux
	@@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags static_all,netgo -a -v -ldflags="-X 'github.com/deepakhcu07/go-rest-api/api.BuildTime=$(TIMESTAMP)' -X 'github.com/deepakhcu07/go-rest-api/api.BuildVersion=$(GIT_COMMIT)'" -o $(CURR_DIR)/bin/$(PROJECT_NAME)-release.o



clean: ## Remove previous build
	@rm -f $(CURR_DIR)/bin/*.o

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'