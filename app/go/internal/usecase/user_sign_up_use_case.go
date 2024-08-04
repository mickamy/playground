package usecase

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"mickamy.com/playground/config"
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
	profileRepo repository.UserProfile
	avatarRepo  repository.UserAvatar
}

func NewUserSignUp(
	db *gorm.DB,
	oauth oauth.OAuth,
	userRepo repository.User,
	accountRepo repository.UserAccount,
	profileRepo repository.UserProfile,
	avatarRepo repository.UserAvatar,
) UserSignUp {
	return userSignUp{
		db:          db,
		oauth:       oauth,
		userRepo:    userRepo,
		accountRepo: accountRepo,
		profileRepo: profileRepo,
		avatarRepo:  avatarRepo,
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
			return fmt.Errorf("failed to create user: %w", err)
		}
		output.Account = model.UserAccount{
			UserID:   user.ID,
			Email:    payload.Email,
			Provider: model.UserAccountProvider(payload.Provider),
			UID:      payload.UID,
		}
		if err := uc.accountRepo.WithTx(tx).Create(ctx, &output.Account); err != nil {
			return fmt.Errorf("failed to create account: %w", err)
		}
		profile := model.UserProfile{
			UserID: user.ID,
			Name:   payload.Name,
		}
		if err := uc.profileRepo.WithTx(tx).Create(ctx, &profile); err != nil {
			return fmt.Errorf("failed to create profile: %w", err)
		}
		if payload.Picture != "" {
			avatar := model.UserAvatar{
				UserID: user.ID,
				Bucket: config.AWS().S3Bucket,
			}
			if err := uc.avatarRepo.WithTx(tx).Create(ctx, &avatar); err != nil {
				return fmt.Errorf("failed to create avatar: %w", err)
			}
		}

		output.Tokens, err = jwt.New(user.ID.String())
		if err != nil {
			return fmt.Errorf("failed to create JWT: %w", err)
		}
		return nil
	})
	return output, err
}
