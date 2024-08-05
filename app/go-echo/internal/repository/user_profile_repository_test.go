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

func TestUserProfile_Create(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(nil)
	assert.NoError(t, db.WithContext(ctx).Create(&user).Error)
	profile := fixture.UserProfile(func(m *model.UserProfile) {
		m.UserID = user.ID
	})

	// act
	sut := repository.NewUserProfile(db)
	err := sut.Create(ctx, &profile)

	// assert
	assert.NoError(t, err)
	var created = model.UserProfile{}
	assert.NoError(t, db.WithContext(ctx).First(&created, "user_id = ?", user.ID).Error)
	assert.Equal(t, profile.Name, created.Name)
	assert.Equal(t, profile.Bio, created.Bio)
}

func TestUserProfile_Get(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(nil)
	assert.NoError(t, db.WithContext(ctx).Create(&user).Error)
	profile := fixture.UserProfile(func(m *model.UserProfile) {
		m.UserID = user.ID
	})
	db.WithContext(ctx).Create(&profile)

	// act
	sut := repository.NewUserProfile(db)
	actual, err := sut.Get(ctx, user.ID.String())

	// assert
	assert.NoError(t, err)
	assert.Equal(t, profile.Name, actual.Name)
	assert.Equal(t, profile.Bio, actual.Bio)
}

func TestUserProfile_Get_WithUser(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(nil)
	assert.NoError(t, db.WithContext(ctx).Create(&user).Error)
	profile := fixture.UserProfile(func(m *model.UserProfile) {
		m.UserID = user.ID
	})
	assert.NoError(t, db.WithContext(ctx).Create(&profile).Error)

	// act
	sut := repository.NewUserProfile(db)
	actual, err := sut.Get(ctx, user.ID.String(), repository.UserAccountWithUser)

	// assert
	assert.NoError(t, err)
	profile.User = user
	assert.Equal(t, profile, actual)
}

func TestUserProfile_GetBySlug(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	ctx := context.Background()
	user := fixture.User(nil)
	assert.NoError(t, db.WithContext(ctx).Create(&user).Error)
	profile := fixture.UserProfile(func(m *model.UserProfile) {
		m.UserID = user.ID
	})
	db.WithContext(ctx).Create(&profile)

	// act
	sut := repository.NewUserProfile(db)
	actual, err := sut.GetBySlug(ctx, user.Slug)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, profile.Name, actual.Name)
	assert.Equal(t, profile.Bio, actual.Bio)
}
