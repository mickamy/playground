ARG GO_VERSION
FROM golang:$GO_VERSION-bullseye AS builder

ENV SQLDEF_VERSION=v0.17.14

RUN curl -LO https://github.com/sqldef/sqldef/releases/download/$SQLDEF_VERSION/psqldef_linux_amd64.tar.gz
RUN tar xf psqldef_linux_amd64.tar.gz
RUN mv psqldef /usr/local/bin/
RUN rm psqldef_linux_amd64.tar.gz

RUN curl -LO https://github.com/sqldef/sqldef/releases/download/$SQLDEF_VERSION/mysqldef_linux_amd64.tar.gz
RUN tar xf mysqldef_linux_amd64.tar.gz
RUN mv mysqldef /usr/local/bin/
RUN  rm mysqldef_linux_amd64.tar.gz

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o api ./cmd/api/main.go

FROM debian:bullseye-slim AS deploy
RUN apt-get update
COPY --from=builder /app/api .
COPY --from=builder /usr/local/bin/psqldef /usr/local/bin/
COPY --from=builder /usr/local/bin/mysqldef /usr/local/bin/

CMD ["./app"]

FROM golang:$GO_VERSION AS dev

RUN apt-get update
RUN apt-get install -y docker.io
RUN rm -rf /var/lib/apt/lists/*

RUN go install github.com/air-verse/air@latest
RUN go install github.com/google/wire/cmd/wire@latest

WORKDIR /app

COPY go.mod go.sum cmd/api/main.go ./
RUN go mod download

COPY --from=builder /usr/local/bin/psqldef /usr/local/bin/
COPY --from=builder /usr/local/bin/mysqldef /usr/local/bin/

ADD entrypoint.dev.sh /app/entrypoint.dev.sh
ENTRYPOINT ["/app/entrypoint.dev.sh"]
