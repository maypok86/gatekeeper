SHELL := /bin/bash

BIN := "./bin/gk"

.PHONY: setup
setup: ## Install all the build and lint dependencies
	bash scripts/setup.sh

.PHONY: build
build: ## Build a version
	bash scripts/build.sh $(BIN)

.PHONY: api
api: build ## Start api server
	echo -n > develop.log && $(BIN) api

.PHONY: up
up:
	docker-compose -f ./deployments/docker-compose.yml up -d

.PHONY: down
down:
	docker-compose -f ./deployments/docker-compose.yml down

.PHONY: version
version: build ## Build and view project version
	$(BIN) version

.PHONY: fmt
fmt: ## Run format tools on all go files
	bash scripts/fmt.sh

.PHONY: lint
lint: ## Run all the linters
	golangci-lint run ./...

.PHONY: unit-test
unit-test: ## Run all unit tests
	bash scripts/unit-tests.sh

.PHONY: integration-test
integration-test: ## Run all integration tests
	bash scripts/integration-tests.sh

.PHONY: cover
cover: unit-test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

.PHONY: ci
ci: lint unit-test ## Run all the tests and code checks

.PHONY: generate
generate: ## Generate code
	bash scripts/generate.sh

.PHONY: clean
clean: ## Remove temporary files
	bash scripts/clean.sh

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL:= help
