package oauth_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"mickamy.com/playground/internal/lib/oauth"
)

func TestValidateToken(t *testing.T) {
	t.Parallel()
	t.Skip("assign your token on testing this func")

	// arrange
	tokenStr := ""

	// act
	sut := oauth.New()
	payload, err := sut.ValidateToken(context.Background(), "google", tokenStr)

	t.Logf("payload: %#v", payload)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, payload.Provider)
	assert.NotEmpty(t, payload.UID)
	assert.NotEmpty(t, payload.Name)
	assert.NotEmpty(t, payload.Email)
	assert.NotEmpty(t, payload.Picture)
}
