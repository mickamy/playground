.PHONY: \
	build \
	up \
	down \
	test \
	db-migrate \

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
