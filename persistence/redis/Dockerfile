ARG REDIS_VERSION=7.1-alpine
FROM redis:${REDIS_VERSION}

COPY ./docker-entrypoint-initdb.d /docker-entrypoint-initdb.d
