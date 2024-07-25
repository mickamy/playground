package fixture

import (
	"github.com/brianvoe/gofakeit/v7"

	"mickamy.com/playground/internal/model"
)

func User(setter func(*model.User)) model.User {
	m := model.User{
		Slug: gofakeit.Name(),
	}
	if setter != nil {
		setter(&m)
	}
	return m
}
