#!/usr/bin/env bash

psqldef -U $DB_USER -W $DB_PASS -h $DB_HOST -p $DB_PORT $DB_NAME < ./db/schema.sql
air
