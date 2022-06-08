CURRENT_DIR := $(PWD)
VERSION := $(shell git describe --exact-match --tags 2> /dev/null || git rev-parse --short HEAD)

postgres:
	@docker run --name postgres-14 -p 5000:5000 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=01000 -d postgres:14-alpine

createDB:
	@docker exec -it exec postgres-14 createdb --username=root --owner=root bank

dropDB:
	@docker exec -it postgres-14 dropdb bank

.PHONY: lint-go
lint-go:
	@golangci-lint run --timeout=5m

.PHONY: test
test: ## Run tests with data race detector
	@go test ./... -race -count=1 -covermode=atomic -coverprofile=coverage.out # -count=1 is the idiomatic way to disable test caching explicitly.

.PHONY: test-start-postgres
test-start-postgres: ## Run tests with data race detector
	@docker-compose run postgresql

.PHONY: migrate-new
migrate-new:
	@migrate create -ext sql -dir migrations $(NAME)

.PHONY: generate
generate:
	@go generate ./...

.PHONY: format
format:
	@npm run format:prettier
