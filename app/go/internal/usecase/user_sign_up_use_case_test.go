package usecase_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

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
	uc := usecase.NewUserSignUp(
		db,
		oauthInstance,
		userRepo,
		accountRepo,
	)

	provider := model.UserAccountProviderGoogle.String()
	idToken := "test-id-token"
	payload := fixture.TokenPayload(nil)

	// mock
	oauthInstance.EXPECT().ValidateToken(gomock.Any(), gomock.Eq(provider), gomock.Eq(idToken)).Return(payload, nil)
	userID := model.ParseUUID(uuid.NewString())
	userRepo.EXPECT().WithTx(gomock.Any()).Return(userRepo)
	user := model.User{
		ID:   userID,
		Slug: "",
	}
	userRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).SetArg(1, user)
	accountRepo.EXPECT().WithTx(gomock.Any()).Return(accountRepo)
	account := model.UserAccount{
		UserID:   userID,
		Email:    payload.Email,
		Provider: model.UserAccountProvider(payload.Provider),
		UID:      payload.UID,
	}
	accountRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).SetArg(1, account)

	// act
	output, err := uc.Do(context.Background(), usecase.UserSignUpInput{
		Provider: provider,
		IDToken:  idToken,
	})
	assert.NoError(t, err)

	// assert
	assert.NotEmpty(t, output)
	assert.NotEmpty(t, output.Account)
	assert.NotEmpty(t, output.Tokens)
}
