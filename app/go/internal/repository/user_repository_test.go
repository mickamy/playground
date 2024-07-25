package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"mickamy.com/playground/internal/model"
	"mickamy.com/playground/internal/repository"
	"mickamy.com/playground/test"
	"mickamy.com/playground/test/fixture"
)

func TestUser_Create(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(nil)

	// act
	sut := repository.NewUser(db)
	err := sut.Create(ctx, &user)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}

func TestUser_Get(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	slug := "test"
	expected := fixture.User(func(user *model.User) {
		user.Slug = slug
	})
	assert.NoError(t, db.WithContext(ctx).Create(&expected).Error)

	// act
	sut := repository.NewUser(db)
	actual, err := sut.Get(ctx, expected.ID.String())

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, slug, actual.Slug)
}

func TestUser_Update(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(nil)
	assert.NoError(t, db.WithContext(ctx).Create(&user).Error)

	// act
	sut := repository.NewUser(db)
	slug := "updated"
	user.Slug = slug
	err := sut.Update(ctx, user)

	// assert
	assert.NoError(t, err)
	var actual model.User
	err = db.WithContext(ctx).First(&actual, "id = ?", user.ID).Error
	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, slug, actual.Slug)
}

func TestUser_Delete(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(nil)
	assert.NoError(t, db.WithContext(ctx).Create(&user).Error)

	// act
	sut := repository.NewUser(db)
	err := sut.Delete(ctx, user.ID.String())
	assert.NoError(t, err)

	// assert
	err = db.WithContext(ctx).First(&user, "id = ?", user.ID).Error
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestUser_GetBySlug(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	expected := fixture.User(nil)
	assert.NoError(t, db.WithContext(ctx).Create(&expected).Error)

	// act
	sut := repository.NewUser(db)
	actual, err := sut.GetBySlug(ctx, expected.Slug)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Slug, actual.Slug)
}
