package config

import (
	"github.com/caarlos0/env/v11"
)

type JWTConfig struct {
	SigningSecret string `env:"JWT_SIGNING_SECRET"`
}

var jwt = JWTConfig{}

func init() {
	if err := env.Parse(&jwt); err != nil {
		panic(err)
	}
}

func JWT() JWTConfig {
	return jwt
}
