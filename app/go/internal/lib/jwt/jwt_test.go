package jwt_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"mickamy.com/playground/internal/lib/jwt"
)

func Test_New(t *testing.T) {
	t.Parallel()

	// arrange
	userID, err := uuid.NewUUID()
	assert.NoError(t, err)

	// act
	token, err := jwt.New(userID.String())

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token.Access.Value)
	assert.NotEmpty(t, token.Access.ExpiresAt)
	assert.NotEmpty(t, token.Refresh.Value)
	assert.NotEmpty(t, token.Refresh.ExpiresAt)
}

func Test_Verify(t *testing.T) {
	t.Parallel()

	// arrange
	userID, err := uuid.NewUUID()
	assert.NoError(t, err)
	token, err := jwt.New(userID.String())
	assert.NoError(t, err)

	// act
	accessClaim, err := jwt.Verify(token.Access.Value)
	refreshClaim, err := jwt.Verify(token.Refresh.Value)

	// assert
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, accessClaim[jwt.UserIDKey], userID.String())
	assert.NotEmpty(t, accessClaim["exp"])
	assert.Equal(t, refreshClaim[jwt.UserIDKey], userID.String())
	assert.Equal(t, refreshClaim["jwt"], token.Access.Value)
	assert.NotEmpty(t, refreshClaim["exp"], token.Access.ExpiresAt)
}
