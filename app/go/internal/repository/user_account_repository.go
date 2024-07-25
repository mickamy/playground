package repository

import (
	"context"

	"gorm.io/gorm"

	"mickamy.com/playground/internal/model"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
type UserAccount interface {
	Creator[model.UserAccount]
	Reader[model.UserAccount, string]
	GetBySlug(ctx context.Context, slug string, scopes ...Scope) (model.UserAccount, error)
	GetByIDToken(ctx context.Context, provider model.UserAccountProvider, uid string, scopes ...Scope) (model.UserAccount, error)
	WithTx(tx *gorm.DB) UserAccount
}

type userAccount struct {
	*gorm.DB
}

func NewUserAccount(db *gorm.DB) UserAccount {
	return userAccount{db}
}

func (repo userAccount) Create(ctx context.Context, m *model.UserAccount) error {
	return repo.WithContext(ctx).Create(m).Error
}

func (repo userAccount) Get(ctx context.Context, userID string, scopes ...Scope) (model.UserAccount, error) {
	var account model.UserAccount
	err := repo.WithContext(ctx).First(&account, "user_id = ?", model.ParseUUID(userID)).Error
	return account, err
}

func (repo userAccount) GetBySlug(ctx context.Context, slug string, scopes ...Scope) (model.UserAccount, error) {
	var account model.UserAccount
	scopes = append(scopes, UserAccountWithUser)
	err := repo.WithContext(ctx).Scopes(scopes...).Where("`User`.slug = ?", slug).First(&account).Error
	return account, err
}

func (repo userAccount) GetByIDToken(ctx context.Context, provider model.UserAccountProvider, uid string, scopes ...Scope) (model.UserAccount, error) {
	var account model.UserAccount
	err := repo.WithContext(ctx).Scopes(scopes...).Where(`provider = ? AND uid = ?`, provider, uid).First(&account).Error
	return account, err
}

func (repo userAccount) WithTx(tx *gorm.DB) UserAccount {
	return userAccount{tx}
}

func UserAccountWithUser(db *gorm.DB) *gorm.DB {
	return db.Joins("User")
}
