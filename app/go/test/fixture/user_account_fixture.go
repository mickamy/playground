package fixture

import (
	"github.com/brianvoe/gofakeit/v7"

	"mickamy.com/playground/internal/model"
)

func UserAccount(setter func(m *model.UserAccount)) model.UserAccount {
	m := model.UserAccount{
		Email:    gofakeit.Email(),
		Provider: UserAccountProvider(),
		UID:      gofakeit.UUID(),
	}
	if setter != nil {
		setter(&m)
	}
	return m
}

func UserAccountGoogle(setter func(account *model.UserAccount)) model.UserAccount {
	m := model.UserAccount{
		Provider: model.UserAccountProviderGoogle,
	}
	if setter != nil {
		setter(&m)
	}
	return m
}
