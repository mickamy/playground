package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"mickamy.com/playground/config"
)

const (
	day  = time.Hour * 24
	week = day * 7

	accessTokenExpiresIn  = time.Hour
	refreshTokenExpiresIn = week * 2

	UserIDKey = "user_id"

	expiredTokenErrorMessage = "Token is expired"
)

var (
	ExpiredTokenError = errors.New("token is expired")
)

var (
	signingMethod = jwt.SigningMethodHS256
	signingSecret = []byte(config.JWT().SigningSecret)
)

type Token struct {
	Value     string
	ExpiresAt time.Time
}

type AccessAndRefresh struct {
	Access  Token
	Refresh Token
}

func New(userID string) (AccessAndRefresh, error) {
	access, err := generateAccessToken(userID)
	if err != nil {
		return AccessAndRefresh{}, err
	}
	refresh, err := generateRefreshToken(userID, access)
	if err != nil {
		return AccessAndRefresh{}, err
	}

	return AccessAndRefresh{Access: access, Refresh: refresh}, nil
}

func generateAccessToken(userID string) (Token, error) {
	claims := jwt.MapClaims{}
	claims[UserIDKey] = userID
	exp := time.Now().Add(accessTokenExpiresIn)
	claims["exp"] = exp.Unix()

	jwtToken := jwt.NewWithClaims(signingMethod, claims)
	accessTokenValue, err := jwtToken.SignedString(signingSecret)
	if err != nil {
		return Token{}, fmt.Errorf("failed to sign access token jwt: %w", err)
	}
	return Token{Value: accessTokenValue, ExpiresAt: exp}, nil
}

func generateRefreshToken(userID string, accessToken Token) (Token, error) {
	claims := jwt.MapClaims{}
	claims[UserIDKey] = userID
	claims["jwt"] = accessToken.Value
	exp := time.Now().Add(refreshTokenExpiresIn)
	claims["exp"] = exp.Unix()

	jwtToken := jwt.NewWithClaims(signingMethod, claims)
	refreshTokenValue, err := jwtToken.SignedString(signingSecret)
	if err != nil {
		return Token{}, fmt.Errorf("failed to sign refresh token jwt: %w", err)
	}
	return Token{Value: refreshTokenValue, ExpiresAt: exp}, nil
}

func Verify(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != signingMethod {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingSecret, nil
	})

	if err != nil {
		if err.Error() == expiredTokenErrorMessage {
			return nil, ExpiredTokenError
		}
		return nil, fmt.Errorf("failed to parse jwt: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid jwt")
}
