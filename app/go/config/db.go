package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type DBConfig struct {
	User     string `env:"DB_USER" envDefault:"playground"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     int    `env:"DB_PORT" envDefault:"3306"`
	Name     string `env:"DB_NAME" envDefault:"playground"`
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
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Name)
}
