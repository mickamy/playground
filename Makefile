.PHONY: \
	build \
	up \
	up-d \
	down \
	down-v \

build:
	docker compose build

up:
	docker compose up

up-d:
	docker compose up -d

down:
	docker compose down

down-v:
	docker compose down --volumes
