COMPOSE := docker compose

## deploy: Deploy the Registry smart contract.
.PHONY: deploy
deploy:
	@$(COMPOSE) up deployer --build
	@sleep 1
	@$(COMPOSE) down deployer

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
