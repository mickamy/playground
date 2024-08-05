//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

func InitializeConfigs() Configs {
	wire.Build(
		NewConfigs,
	)
	return Configs{}
}

func InitializeLibs() Libs {
	wire.Build(
		libs,
		wire.Struct(new(Libs), "*"),
	)
	return Libs{}
}

func InitializeInfra() Infras {
	wire.Build(
		infras,
		configs,
		wire.Struct(new(Infras), "*"),
	)
	return Infras{}
}

func InitializeRepositories() Repositories {
	wire.Build(
		repositories,
		configs,
		infras,
		wire.Struct(new(Repositories), "*"),
	)
	return Repositories{}
}

func InitializeUseCases() UseCases {
	wire.Build(
		useCases,
		configs,
		libs,
		infras,
		repositories,
		wire.Struct(new(UseCases), "*"),
	)
	return UseCases{}
}

func InitializeControllers() Controllers {
	wire.Build(
		controllers,
		configs,
		libs,
		infras,
		repositories,
		useCases,
		wire.Struct(new(Controllers), "*"),
	)
	return Controllers{}
}
