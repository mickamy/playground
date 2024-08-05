package di

import (
	"github.com/google/wire"

	"mickamy.com/playground/config"
)

type Configs struct {
	API    config.APIConfig
	AWS    config.AWSConfig
	Common config.CommonConfig
	DB     config.DBConfig
	JWT    config.JWTConfig
	OAuth  config.OAuthConfig
}

func NewConfigs() Configs {
	return Configs{
		API:    config.API(),
		AWS:    config.AWS(),
		Common: config.Common(),
		DB:     config.DB(),
		JWT:    config.JWT(),
		OAuth:  config.OAuth(),
	}
}

var configs = wire.NewSet(
	config.API,
	config.AWS,
	config.Common,
	config.DB,
	config.JWT,
	config.OAuth,
)
