.PHONY: \
	build \
	up \
	down \
	test \
	db-create \
	db-drop \
	db-migrate \
	db-migrate-dry \
	db-reset \
	wire \

build:
	cd ../../ && make build
	docker compose build

up: down
	cd ../../ && make down && make up-d
	docker compose up

down:
	cd ../../ && make down
	docker compose down

test:
	go test ./...

db-create:
	go run ./cmd/db/create/main.go

db-drop:
	go run ./cmd/db/drop/main.go

db-migrate:
	go run ./cmd/db/migrate/main.go

db-reset:
	go run ./cmd/db/reset/main.go

wire:
	cd ./internal/di && wire
