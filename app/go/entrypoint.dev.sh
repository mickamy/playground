#!/usr/bin/env bash

mysqldef -u "$DB_USER" -p "$DB_PASSWORD" -h "$DB_HOST" -P "$DB_PORT" "$DB_NAME" < ./db/schema.sql

exec "$@"
