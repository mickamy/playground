package config

import (
	"fmt"
	"net"
	"strings"

	"github.com/caarlos0/env/v11"
)

type DBConfig struct {
	URL string `env:"DATABASE_URL"`
}

func (cfg DBConfig) ParsedDBConfig() (ParsedDBConfig, error) {
	parts := strings.Split(cfg.URL, "@")
	if len(parts) != 2 {
		return ParsedDBConfig{}, fmt.Errorf("invalid DB URL format")
	}

	credentials := strings.Split(parts[0], ":")
	if len(credentials) != 2 {
		return ParsedDBConfig{}, fmt.Errorf("invalid credentials format")
	}
	user := credentials[0]
	pass := credentials[1]

	hostParts := strings.Split(parts[1], "/")
	if len(hostParts) < 1 {
		return ParsedDBConfig{}, fmt.Errorf("invalid host format")
	}

	hostPort := strings.Trim(hostParts[0], "tcp()")
	host, port, err := net.SplitHostPort(hostPort)
	if err != nil {
		return ParsedDBConfig{}, fmt.Errorf("failed to split host and port: %w", err)
	}

	name := strings.Split(hostParts[1], "?")[0]

	return ParsedDBConfig{
		User: user,
		Pass: pass,
		Host: host,
		Port: port,
		Name: name,
	}, nil
}

type ParsedDBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

var db = DBConfig{}

func init() {
	if err := env.Parse(&db); err != nil {
		panic(err)
	}
}

func DB() DBConfig {
	return db
}
