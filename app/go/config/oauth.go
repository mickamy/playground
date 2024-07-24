package config

import (
	"github.com/caarlos0/env/v11"
)

type OAuthConfig struct {
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
}

var oauth = OAuthConfig{}

func init() {
	if err := env.Parse(&oauth); err != nil {
		panic(err)
	}
}

func OAuth() OAuthConfig {
	return oauth
}
