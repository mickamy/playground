package config

import (
	"fmt"
	"log/slog"

	"github.com/caarlos0/env/v11"
)

type CommonConfig struct {
	Env         string `env:"ENV" envDefault:"development"`
	PackageRoot string `env:"PACKAGE_ROOT"`
}

var common = CommonConfig{}

func init() {
	if err := env.Parse(&common); err != nil {
		panic(err)
	}
	if common.Env == "" || common.PackageRoot == "" {
		panic(fmt.Errorf("some of required environment variables are missing: %#v", db))
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
