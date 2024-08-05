package di

import (
	"github.com/google/wire"

	"mickamy.com/playground/internal/usecase"
)

type UseCases struct {
	UserSignUp usecase.UserSignUp
}

var useCases = wire.NewSet(
	usecase.NewUserSignUp,
)
