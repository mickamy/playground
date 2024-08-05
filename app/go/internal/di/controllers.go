package di

import (
	"github.com/google/wire"

	"mickamy.com/playground/internal/api/controller"
)

type Controllers struct {
	User controller.User
}

var controllers = wire.NewSet(
	controller.NewUser,
)
