package oauth

import (
	"context"
	"fmt"
)

const (
	ProviderGoogle = "google"
)

type TokenPayload struct {
	Provider string
	UID      string
	Name     string
	Email    string
	Picture  string
}

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
type OAuth interface {
	ValidateToken(ctx context.Context, provider string, tokenStr string) (TokenPayload, error)
}

type oauth struct{}

func (oauth oauth) ValidateToken(ctx context.Context, provider string, tokenStr string) (TokenPayload, error) {
	switch provider {
	case ProviderGoogle:
		return ValidateGoogleToken(ctx, tokenStr)
	default:
		return TokenPayload{}, fmt.Errorf("unknown provider %s", provider)
	}
}

func New() OAuth {
	return oauth{}
}
