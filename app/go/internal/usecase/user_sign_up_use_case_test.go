package usecase_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"mickamy.com/playground/config"
	oauthMock "mickamy.com/playground/internal/lib/oauth/mock"
	"mickamy.com/playground/internal/model"
	repoMock "mickamy.com/playground/internal/repository/mock"
	"mickamy.com/playground/internal/usecase"
	"mickamy.com/playground/test"
	"mickamy.com/playground/test/fixture"
)

func TestUserSignUp_Do(t *testing.T) {
	t.Parallel()

	// arrange
	db := test.NewTestDB(t)
	controller := gomock.NewController(t)
	oauthInstance := oauthMock.NewMockOAuth(controller)
	userRepo := repoMock.NewMockUser(controller)
	accountRepo := repoMock.NewMockUserAccount(controller)
	profileRepo := repoMock.NewMockUserProfile(controller)
	avatarRepo := repoMock.NewMockUserAvatar(controller)
	uc := usecase.NewUserSignUp(
		db,
		oauthInstance,
		userRepo,
		accountRepo,
		profileRepo,
		avatarRepo,
	)

	provider := model.UserAccountProviderGoogle.String()
	idToken := "test-id-token"
	payload := fixture.TokenPayload(nil)
	userID := model.ParseUUID(uuid.NewString())
	user := model.User{
		ID:   userID,
		Slug: "",
	}
	account := model.UserAccount{
		UserID:   userID,
		Email:    payload.Email,
		Provider: model.UserAccountProvider(payload.Provider),
		UID:      payload.UID,
	}
	profile := model.UserProfile{
		UserID: userID,
		Name:   payload.Name,
	}
	avatar := model.UserAvatar{
		UserID: userID,
		Bucket: config.AWS().S3Bucket,
	}

	// mock
	oauthInstance.EXPECT().ValidateToken(gomock.Any(), gomock.Eq(provider), gomock.Eq(idToken)).Return(payload, nil)
	userRepo.EXPECT().WithTx(gomock.Any()).Return(userRepo)
	userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).SetArg(1, user)
	accountRepo.EXPECT().WithTx(gomock.Any()).Return(accountRepo)
	accountRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).SetArg(1, account)
	profileRepo.EXPECT().WithTx(gomock.Any()).Return(profileRepo)
	profileRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).SetArg(1, profile)
	avatarRepo.EXPECT().WithTx(gomock.Any()).Return(avatarRepo)
	avatarRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).SetArg(1, avatar)

	// act
	output, err := uc.Do(context.Background(), usecase.UserSignUpInput{
		Provider: provider,
		IDToken:  idToken,
	})

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, output)
	assert.NotEmpty(t, output.Account)
	assert.Equal(t, payload.Email, output.Account.Email)
	assert.Equal(t, payload.Provider, output.Account.Provider.String())
	assert.Equal(t, payload.UID, output.Account.UID)
	assert.NotEmpty(t, output.Tokens)
}
