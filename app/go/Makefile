.PHONY: \
	build \
	up \
	down \
	test \
	db-migrate \
	wire \

build:
	cd ../../ && make build && cd -
	docker compose build

up:
	cd ../../ && make down && make up-d && cd -
	docker compose up

down:
	cd ../../ && make down && cd -
	docker compose down

test: db-migrate
	go test ./...

db-migrate:
	mysqldef -u $$DB_USER -p$$DB_PASSWORD -h $$DB_HOST -P $$DB_PORT $$DB_NAME < ./db/schema.sql

proto:
	protoc \
		--proto_path=../../proto \
		--go_out=./internal/api/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=./internal/api/pb \
		--go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		$$(find ../../proto -name "*.proto")

wire:
	cd ./internal/di && wire && cd -
