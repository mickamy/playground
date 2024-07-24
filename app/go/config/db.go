package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type DBConfig struct {
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
	Host string `env:"DB_HOST"`
	Port int    `env:"DB_PORT"`
	Name string `env:"DB_NAME"`
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

func (config DBConfig) DatabaseURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.User, config.Pass, config.Host, config.Port, config.Name)
}
