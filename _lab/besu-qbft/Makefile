COMPOSE := docker compose

.PHONY: besu-up
## besu-up: Start the Besu private QBFT network.
besu-up:
	@$(COMPOSE) up besu-qbft-1 besu-qbft-2 besu-qbft-3 besu-qbft-4 --build
	@$(MAKE) besu-ip

.PHONY: besu-down
## besu-down: Stop the Besu private QBFT network.
besu-down:
	@$(COMPOSE) down besu-qbft-1 besu-qbft-2 besu-qbft-3 besu-qbft-4

.PHONY: besu-ip
## besu-ip: Get the IP address of the current machine and update 'CHAIN_URL'.
besu-ip:
	@./scripts/ip_mapper.sh

.PHONY: besu-purge
## besu-purge: Purge the existing network data (destructive action).
besu-purge:
	@echo "WARNING! This will remove:"
	@echo "  - all node/caches"
	@echo "  - all node/database\n"

	@read -p "Are you sure you want to continue? [y/N] " ans && ans=$${ans:-N} ; \
    if [ $${ans} = y ] || [ $${ans} = Y ]; then \
        sudo find ./*/data/* -type f ! \( -name 'key' -o -name 'key.pub' \) -delete ; \
		sudo find ./*/data/* -type d -empty -delete ; \
		printf "All data deleted.\n" ; \
    else \
        printf "Target aborted.\n" ; \
    fi

## help: Display all targets.
.PHONY: help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make [target]'
	@echo ''
	@echo 'Targets:'
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' |  sed -e 's/^/ /'

.DEFAULT_GOAL := help
