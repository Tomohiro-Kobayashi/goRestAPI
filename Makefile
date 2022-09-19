.PHONY: help build build-local up down logs as test
.DEFAULT_GOAL := help


DOCKER_TAG := latest
build:
	docker build -t tkobayashi200/gotodo:${DOCKER_TAG} --target deploy ./

build-local: ## build docker image to deply
	docker compose build --no-cache

up: ## do docker compose up
	docker compose up -d

down: ## do docker compose down
	docker compose down

logs: ## tail docker compose logs
	docker compose logs

ps: ## check container status
	docker compose ps

test: ## do tests
	go test -race -shuffle=on ./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'