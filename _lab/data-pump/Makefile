ifneq (,$(wildcard ./.env))
    include .env
    export
endif

COMPOSE := docker compose

## deploy: Deploy the Registry smart contract.
.PHONY: deploy
deploy:
	@$(COMPOSE) up deployer --build
	@sleep 1
	@$(COMPOSE) down deployer

## pump: Pump data to the Registry smart contract.
.PHONY: pump
pump:
	@$(COMPOSE) up pumper --build -d

## validate: Validate data from Registry smart contract.
.PHONY: validate
validate:
	@$(COMPOSE) up validator --scale validator=$(REPLICAS) --build
	@./cmd/shell/post-validate.sh

## api: Start API server for communication.
.PHONY: api
api:
	@$(COMPOSE) up api --build -d

## inject: Inject data to external API.
.PHONY: inject
inject:
	@$(COMPOSE) up injector --build -d

## enter: Enter the database.
.PHONY: enter
enter:
	@docker exec -it $(DB_CONTAINER) psql -d $(DB_NAME) -U $(DB_USER) -W

## down: Down all containers.
.PHONY: down
down:
	@$(COMPOSE) down

## help: Display all targets.
.PHONY: help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make [target]'
	@echo ''
	@echo 'Targets:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.DEFAULT_GOAL := help
