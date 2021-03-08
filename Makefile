.PHONY: help

UID = $(shell id -u)
GID = $(shell id -g)

default: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?##.*$$' $(MAKEFILE_LIST) | sort | awk '{split($$0, a, ":"); printf "\033[36m%-30s\033[0m %-30s %s\n", a[1], a[2], a[3]}'

#
# Executes a command in a running container, mainly useful to fix the terminal size on opening a shell session
#
# $(1) the options
#
define infra-shell
	docker-compose exec -e COLUMNS=`tput cols` -e LINES=`tput lines` $(1)
endef

######################################
#              APP                   #
######################################
.PHONY: app-build app-build-run app-run-binary

app-build: ## to build app into bin/server path
	@go build -o bin/server server.go

app-build-run: ## to build && to run binary application
	@make app-build app-run-binary

app-run-binary: ## to build app into bin/server path
	@./bin/server

########################################
#              INFRA                   #
########################################

.PHONY: infra-clean infra-shell-db infra-stop infra-up

infra-clean: ## to stop and remove containers, networks, images
	@docker-compose down --rmi all

infra-stop: ## to stop containers
	@docker-compose stop

infra-shell-db: ## to open a shell session in the db container
	@$(call infra-shell,db sh)

infra-up: ## to start all the containers
	@if [ ! -f .env -a -f .env.dist ]; then sed "s,#UID#,$(UID),g;s,#GID#,$(GID),g" .env.dist > .env; fi
	@docker-compose up --build -d