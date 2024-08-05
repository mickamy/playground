package fixture

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/mattn/go-gimei"

	"mickamy.com/playground/internal/model"
)

func UserProfile(setter func(m *model.UserProfile)) model.UserProfile {
	name := gimei.NewName()
	m := model.UserProfile{
		Name: name.Kanji(),
		Bio:  gofakeit.LoremIpsumParagraph(1, 1, 20, "."),
	}
	if setter != nil {
		setter(&m)
	}
	return m
}
