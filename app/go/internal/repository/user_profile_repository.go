package repository

import (
	"context"

	"gorm.io/gorm"

	"mickamy.com/playground/internal/model"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
type UserProfile interface {
	Creator[model.UserProfile]
	Reader[model.UserProfile, string]
	GetBySlug(ctx context.Context, slug string, scopes ...Scope) (model.UserProfile, error)
	WithTx(tx *gorm.DB) UserProfile
}

type userProfile struct {
	*gorm.DB
}

func NewUserProfile(db *gorm.DB) UserProfile {
	return userProfile{db}
}

func (repo userProfile) Create(ctx context.Context, m *model.UserProfile) error {
	return repo.WithContext(ctx).Create(m).Error
}

func (repo userProfile) Get(ctx context.Context, userID string, scopes ...Scope) (model.UserProfile, error) {
	var m model.UserProfile
	err := repo.WithContext(ctx).Scopes(scopes...).First(&m, "user_id = ?", model.ParseUUID(userID)).Error
	return m, err
}

func (repo userProfile) GetBySlug(ctx context.Context, slug string, scopes ...Scope) (model.UserProfile, error) {
	var m model.UserProfile
	scopes = append(scopes, UserProfileWithUser)
	err := repo.WithContext(ctx).Scopes(scopes...).First(&m, "`User`.slug = ?", slug).Error
	return m, err
}

func (repo userProfile) WithTx(tx *gorm.DB) UserProfile {
	return userProfile{tx}
}

func UserProfileWithUser(db *gorm.DB) *gorm.DB {
	return db.Joins("User")
}
