package oauth

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/idtoken"

	"mickamy.com/playground/config"
)

func GoogleConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.OAuth().GoogleClientID,
		ClientSecret: config.OAuth().GoogleClientSecret,
		RedirectURL:  config.API().BackendURL() + "/oauth/callback/google",
		Scopes: []string{
			"openid",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func ValidateGoogleToken(ctx context.Context, tokenStr string) (TokenPayload, error) {
	cnf := config.OAuth()
	payload, err := idtoken.Validate(ctx, tokenStr, cnf.GoogleClientID)
	if err != nil {
		return TokenPayload{}, fmt.Errorf("google token validation failed: %w", err)
	}

	claims := payload.Claims
	sub, ok1 := claims["sub"].(string)
	name, ok2 := claims["name"].(string)
	email, ok3 := claims["email"].(string)
	picture, ok4 := claims["picture"].(string)

	if !ok1 || sub == "" {
		return TokenPayload{}, fmt.Errorf("google token missing 'sub' claim")
	}
	if !ok2 || name == "" {
		return TokenPayload{}, fmt.Errorf("google token missing 'name' claim")
	}
	if !ok3 || email == "" {
		return TokenPayload{}, fmt.Errorf("google token missing 'email' claim")
	}
	if !ok4 || picture == "" {
		picture = ""
	}

	return TokenPayload{
		Provider: ProviderGoogle,
		UID:      sub,
		Name:     name,
		Email:    email,
		Picture:  picture,
	}, nil
}
