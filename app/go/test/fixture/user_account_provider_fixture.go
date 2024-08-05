package fixture

import (
	"github.com/brianvoe/gofakeit/v7"

	"mickamy.com/playground/internal/model"
)

func UserAccountProvider() model.UserAccountProvider {
	str := gofakeit.RandomString([]string{
		model.UserAccountProviderGoogle.String(),
	})
	return model.UserAccountProvider(str)
}
