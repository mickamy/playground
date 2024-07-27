package repository

import (
	"context"

	"gorm.io/gorm"

	"mickamy.com/playground/internal/model"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
type UserAvatar interface {
	Creator[model.UserAvatar]
	WithTx(tx *gorm.DB) UserAvatar
}

type userAvatar struct {
	*gorm.DB
}

func NewUserAvatar(db *gorm.DB) UserAvatar {
	return userAvatar{db}
}

func (repo userAvatar) Create(ctx context.Context, m *model.UserAvatar) error {
	return repo.WithContext(ctx).Create(&m).Error
}

func (repo userAvatar) WithTx(tx *gorm.DB) UserAvatar {
	return userAvatar{tx}
}
