SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

BIN		:= pflagstruct
OUTDIR	:= bin
GO		:= go
LDFLAGS	:= -s -w

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.PHONY: build
build: ## Compiles the source code.
	$(GO) build -trimpath -ldflags '$(LDFLAGS)' -o $(OUTDIR)/$(BIN) .

.PHONY: test
test: ## Run all unit tests.
	$(GO) test ./...

.PHONY: install
install: ## Build and install the binary
	$(GO) install -trimpath -ldflags '$(LDFLAGS)' .
