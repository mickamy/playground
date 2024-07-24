ARG MYSQL_VERSION=8.0.36-alpine
FROM mysql:${MYSQL_VERSION}

COPY ./docker-entrypoint-initdb.d /docker-entrypoint-initdb.d
