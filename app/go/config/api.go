package config

import (
	"github.com/caarlos0/env/v11"
)

type APIConfig struct {
	Port         int    `env:"PORT" envDefault:"3000"`
	BackendHost  string `env:"BACKEND_HOST"`
	FrontendHost string `env:"FRONTEND_HOST"`
}

func (config APIConfig) BackendURL() string {
	return "https://" + config.BackendHost
}

func (config APIConfig) FrontendURL() string {
	return "https://" + config.FrontendHost
}

var api = APIConfig{}

func init() {
	if err := env.Parse(&api); err != nil {
		panic(err)
	}
}

func API() APIConfig {
	return api
}
