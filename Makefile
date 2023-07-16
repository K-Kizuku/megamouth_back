container_name := megamouth-api

DOCKER_COMPOSE_VERSION_CHECKER := $(shell docker compose > /dev/null 2>&1 ; echo $$?)
ifeq ($(DOCKER_COMPOSE_VERSION_CHECKER), 0)
	DOCKER_COMPOSE=docker compose
else
	DOCKER_COMPOSE=docker-compose
endif

init: up run

build:
	$(DOCKER_COMPOSE) build

up:
	$(DOCKER_COMPOSE) up -d

up-b:
	$(DOCKER_COMPOSE) up -d --build  

down:
	$(DOCKER_COMPOSE) down

exec:
	docker exec -it $(container_name) bash

run:
	docker exec -it $(container_name) bash -c "go run main.go"

swag:
	swag init -g main.go

d-build:
	docker build -t linux-go-api --platform linux/x86_64 . 


