container_name := megamouth-api

DOCKER_COMPOSE_VERSION_CHECKER := $(shell docker compose > /dev/null 2>&1 ; echo $$?)
ifeq ($(DOCKER_COMPOSE_VERSION_CHECKER), 0)
	DOCKER_COMPOSE=docker compose
else
	DOCKER_COMPOSE=docker-compose
endif

init: up-d run

build:
	$(DOCKER_COMPOSE) build

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

exec:
	docker exec -it $(container_name) bash

run:
	docker exec -it $(container_name) bash -c "go run main.go"
