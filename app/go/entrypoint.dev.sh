#!/usr/bin/env bash

mysqldef -U "$DB_USER" -W "$DB_PASS" -h "$DB_HOST" -p "$DB_PORT" "$DB_NAME" < ./db/schema.sql
air
