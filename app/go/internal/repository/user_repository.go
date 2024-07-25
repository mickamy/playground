package repository

import (
	"context"

	"gorm.io/gorm"

	"mickamy.com/playground/internal/model"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
type User interface {
	Creator[model.User]
	Reader[model.User, model.BinaryUUID]
	Updater[model.User]
	Deleter[model.BinaryUUID]
	GetBySlug(ctx context.Context, slug string, scopes ...Scope) (model.User, error)
	WithTx(tx *gorm.DB) User
}

type user struct {
	*gorm.DB
}

func NewUser(db *gorm.DB) User {
	return user{db}
}

func (repo user) Create(ctx context.Context, m *model.User) error {
	return repo.WithContext(ctx).Create(m).Error
}

func (repo user) Get(ctx context.Context, id model.BinaryUUID, scopes ...Scope) (model.User, error) {
	var user model.User
	err := repo.WithContext(ctx).Scopes(scopes...).First(&user, "id = ?", id).Error
	return user, err
}

func (repo user) Update(ctx context.Context, m model.User) error {
	err := repo.WithContext(ctx).Updates(m).Error
	return err
}

func (repo user) Delete(ctx context.Context, id model.BinaryUUID) error {
	return repo.WithContext(ctx).Delete(&model.User{ID: id}).Error
}

func (repo user) GetBySlug(ctx context.Context, slug string, scopes ...Scope) (model.User, error) {
	var user model.User
	err := repo.WithContext(ctx).Scopes(scopes...).Where("slug = ?", slug).First(&user).Error
	return user, err
}

func (repo user) WithTx(tx *gorm.DB) User {
	return user{tx}
}