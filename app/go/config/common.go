package config

import (
	"log/slog"

	"github.com/caarlos0/env/v11"
)

type CommonConfig struct {
	Env string `env:"ENV" envDefault:"development"`
}

var common = CommonConfig{}

func init() {
	if err := env.Parse(&common); err != nil {
		panic(err)
	}
}

func Common() CommonConfig {
	return common
}

func (config CommonConfig) LogLevel() slog.Level {
	switch config.Env {
	case "development", "test":
		return slog.LevelDebug
	case "staging", "production":
		return slog.LevelInfo
	}
	return slog.LevelError
}
