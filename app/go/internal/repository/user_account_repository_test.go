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

func TestUserAccount_Create(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(func(user *model.User) {})
	assert.NoError(t, repository.NewUser(db).Create(ctx, &user))
	account := fixture.UserAccount(func(m *model.UserAccount) { m.UserID = user.ID })

	// act
	sut := repository.NewUserAccount(db)
	err := sut.Create(ctx, &account)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, account.UserID)
	assert.NotEmpty(t, account.CreatedAt)
	assert.NotEmpty(t, account.UpdatedAt)
}

func TestUserAccount_Get(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(nil)
	assert.NoError(t, repository.NewUser(db).Create(ctx, &user))
	expected := fixture.UserAccount(func(m *model.UserAccount) { m.UserID = user.ID })
	assert.NoError(t, db.Save(&expected).Error)

	// act
	sut := repository.NewUserAccount(db)
	actual, err := sut.Get(ctx, expected.UserID.String())

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUserAccount_GetBySlug(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(func(user *model.User) {})
	assert.NoError(t, repository.NewUser(db).Create(ctx, &user))
	expected := fixture.UserAccount(func(m *model.UserAccount) {
		m.UserID = user.ID
	})
	assert.NoError(t, db.Save(&expected).Error)
	expected.User = user

	// act
	sut := repository.NewUserAccount(db)
	actual, err := sut.GetBySlug(ctx, user.Slug)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUserAccount_GetByIDToken(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(func(user *model.User) {})
	assert.NoError(t, repository.NewUser(db).Create(ctx, &user))
	expected := fixture.UserAccount(func(m *model.UserAccount) { m.UserID = user.ID })
	assert.NoError(t, db.Save(&expected).Error)

	// act
	sut := repository.NewUserAccount(db)
	actual, err := sut.GetByIDToken(ctx, expected.Provider, expected.UID)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUserAccount_UserAccountWithUser(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(func(user *model.User) {})
	assert.NoError(t, repository.NewUser(db).Create(ctx, &user))
	expected := fixture.UserAccount(func(m *model.UserAccount) {
		m.UserID = user.ID
	})
	assert.NoError(t, db.Save(&expected).Error)

	// act
	sut := repository.NewUserAccount(db)
	actual, err := sut.GetByIDToken(ctx, expected.Provider, expected.UID, repository.UserAccountWithUser)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, actual.User)
}
