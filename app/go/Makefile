.PHONY: test db-migrate

test:
	PACKAGE_ROOT=$$(pwd) ENV=test go test -v ./...

db-migrate:
	mysqldef -u $$DB_USER -p$$DB_PASSWORD -h $$DB_HOST -P $$DB_PORT $$DB_NAME < ./db/schema.sql
