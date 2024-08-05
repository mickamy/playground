package fixture

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"

	"mickamy.com/playground/internal/lib/oauth"
)

func TokenPayload(setter func(payload *oauth.TokenPayload)) oauth.TokenPayload {
	payload := oauth.TokenPayload{
		Provider: UserAccountProvider().String(),
		UID:      uuid.NewString(),
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Picture:  gofakeit.URL(),
	}
	if setter != nil {
		setter(&payload)
	}
	return payload
}
