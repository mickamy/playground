.PHONY: \
	build \
	up \
	down \
	test \
	db-migrate \
	wire \

build:
	cd ../../ && make build
	docker compose build

up:
	cd ../../ && make down && make up-d
	docker compose up

down:
	cd ../../ && make down
	docker compose down

test: db-migrate
	go test ./...

db-migrate:
	mysqldef -u $$DB_USER -p$$DB_PASSWORD -h $$DB_HOST -P $$DB_PORT $$DB_NAME < ./db/schema.sql

wire:
	cd ./internal/di && wire
