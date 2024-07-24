ARG POSTGRES_VERSION=14.6.8-alpine
FROM postgres:${POSTGRES_VERSION}

COPY ./docker-entrypoint-initdb.d /docker-entrypoint-initdb.d
