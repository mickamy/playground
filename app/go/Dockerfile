ARG GO_VERSION
FROM golang:$GO_VERSION-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o api ./cmd/api/main.go

FROM debian:bullseye-slim AS deploy
RUN apt-get update
COPY --from=builder /app/api .

CMD ["./app"]

FROM golang:$GO_VERSION AS dev
WORKDIR /app

COPY go.mod go.sum cmd/api/main.go ./
RUN go mod download

RUN <<EOF
  go install github.com/air-verse/air@latest

  curl -LO https://github.com/sqldef/sqldef/releases/download/v0.17.11/psqldef_linux_amd64.tar.gz
  tar xf psqldef_linux_amd64.tar.gz
  mv psqldef /usr/local/bin/
  rm psqldef_linux_amd64.tar.gz

  curl -LO https://github.com/sqldef/sqldef/releases/download/v0.17.11/mysql_linux_amd64.tar.gz
  tar xf mysql_linux_amd64.tar.gz
  mv mysql /usr/local/bin/
  rm mysql_linux_amd64.tar.gz
EOF

ADD entrypoint.dev.sh /app/entrypoint.dev.sh
ENTRYPOINT ["/app/entrypoint.dev.sh"]
