package config

import (
	"github.com/caarlos0/env/v11"
)

type DBConfig struct {
	URL string `env:"DATABASE_URL"`
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
