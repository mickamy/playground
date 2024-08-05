package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"mickamy.com/playground/internal/model"
	"mickamy.com/playground/internal/repository"
	"mickamy.com/playground/test"
	"mickamy.com/playground/test/fixture"
)

func TestUserAvatar_Create(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(func(user *model.User) {})
	assert.NoError(t, repository.NewUser(db).Create(ctx, &user))
	avatar := fixture.UserAvatar(func(m *model.UserAvatar) { m.UserID = user.ID })

	// act
	sut := repository.NewUserAvatar(db)
	err := sut.Create(ctx, &avatar)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, avatar.UserID)
	assert.NotEmpty(t, avatar.CreatedAt)
	assert.NotEmpty(t, avatar.UpdatedAt)
}
