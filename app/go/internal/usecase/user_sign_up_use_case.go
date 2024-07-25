package usecase

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"mickamy.com/playground/internal/lib/jwt"
	"mickamy.com/playground/internal/lib/oauth"
	"mickamy.com/playground/internal/model"
	"mickamy.com/playground/internal/repository"
)

type UserSignUpInput struct {
	Provider string
	IDToken  string
}
type UserSignUpOutput struct {
	Account model.UserAccount
	Tokens  jwt.AccessAndRefresh
}

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
type UserSignUp interface {
	Do(ctx context.Context, input UserSignUpInput) (UserSignUpOutput, error)
}

type userSignUp struct {
	db          *gorm.DB
	oauth       oauth.OAuth
	userRepo    repository.User
	accountRepo repository.UserAccount
}

func NewUserSignUp(
	db *gorm.DB,
	oauth oauth.OAuth,
	userRepo repository.User,
	accountRepo repository.UserAccount,
) UserSignUp {
	return userSignUp{
		db:          db,
		oauth:       oauth,
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
}

func (uc userSignUp) Do(ctx context.Context, input UserSignUpInput) (UserSignUpOutput, error) {
	output := UserSignUpOutput{}
	payload, err := uc.oauth.ValidateToken(ctx, input.Provider, input.IDToken)
	if err != nil {
		return output, err
	}

	err = uc.db.Transaction(func(tx *gorm.DB) error {
		user := model.User{}
		if err := uc.userRepo.WithTx(tx).Create(ctx, &user); err != nil {
			return errors.Wrap(err, "failed to create user")
		}
		output.Account = model.UserAccount{
			Email:    payload.Email,
			Provider: model.UserAccountProvider(payload.Provider),
			UID:      payload.UID,
			UserID:   user.ID,
		}
		err := uc.accountRepo.WithTx(tx).Create(ctx, &output.Account)
		if err != nil {
			return errors.Wrap(err, "failed to create account")
		}

		output.Tokens, err = jwt.New(user.ID.String())
		if err != nil {
			return errors.Wrap(err, "failed to create JWT")
		}
		return nil
	})
	return output, err
}
