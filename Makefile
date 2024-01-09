default: help

help: ## Help.
	@echo "Use \`make <target>\` where <target> one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
	    awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

up: ## Compose up server and client.
	docker-compose -f ./deploy/docker-compose.yaml up --detach --renew-anon-volumes --remove-orphans

down: ## Compose down server and client.
	docker-compose -f ./deploy/docker-compose.yaml down --remove-orphans --rmi local
