# docker compose
TARGET ?=
ENV_FILE ?= .env
COMPOSE_CMD = docker compose -f compose.yaml --env-file $(ENV_FILE)

####
##@ General
####

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

####
##@ Development
####

.PHONY: test
test: ## Run Go tests.
	go test ./...

.PHONY: gen-oapi-client
gen-oapi-client: oapi-codegen ## Generate OpenAPI client codes.
	$(OAPI_CODEGEN) --config=oapi-codegen-client.yaml ./api/openapi.yaml

.PHONY: gen-oapi-server
gen-oapi-server: oapi-codegen ## Generate OpenAPI server codes.
	$(OAPI_CODEGEN) --config=oapi-codegen-server.yaml ./api/openapi.yaml

####
##@ Docker Compose
####

.PHONY: compose-up
compose-up: ## Run components.
	$(COMPOSE_CMD) up -d $(TARGET)

.PHONY: compose-down
compose-down: ## Shutdown components.
	$(COMPOSE_CMD) down $(TARGET)

.PHONY: compose-ps
compose-ps: ## Print running components.
	$(COMPOSE_CMD) ps $(TARGET)

.PHONY: compose-logs
compose-logs: ## Tail logs of components.
	$(COMPOSE_CMD) logs -f $(TARGET)

####
##@ Tools
####

# Location to install dependencies to.
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

# Modified path environment variable including dependencies.
LOCALPATH ?= $(LOCALBIN):$(PATH)

# Tool Binaries
OAPI_CODEGEN ?= $(LOCALBIN)/oapi-codegen

.PHONY: oapi-codegen
oapi-codegen: $(OAPI_CODEGEN) ## Install protoc-gen-go locally if necessary.
$(OAPI_CODEGEN): $(LOCALBIN)
	@test -s $(OAPI_CODEGEN) || \
	GOBIN=$(LOCALBIN) go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
