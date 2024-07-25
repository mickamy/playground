package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"mickamy.com/playground/internal/api/controller"
	"mickamy.com/playground/internal/api/dto/input"
	"mickamy.com/playground/internal/api/dto/output"
	"mickamy.com/playground/internal/lib/jwt"
	"mickamy.com/playground/internal/model"
	"mickamy.com/playground/internal/usecase"
	mockUseCase "mickamy.com/playground/internal/usecase/mock"
	"mickamy.com/playground/test"
	"mickamy.com/playground/test/fixture"
)

func TestUser_SignUp_Unit(t *testing.T) {
	t.Parallel()

	// arrange
	var buf bytes.Buffer
	in := input.UserSignUp{
		Provider: model.UserAccountProviderGoogle.String(),
		IDToken:  "test-id-token",
	}
	assert.NoError(t, json.NewEncoder(&buf).Encode(in))
	tokens, err := jwt.New(uuid.NewString())
	assert.NoError(t, err)
	res := usecase.UserSignUpOutput{
		Account: fixture.UserAccount(nil),
		Profile: fixture.UserProfile(nil),
		Tokens:  tokens,
	}
	ctrl := gomock.NewController(t)
	uc := mockUseCase.NewMockUserSignUp(ctrl)
	uc.EXPECT().Do(gomock.Any(), gomock.Any()).Return(res, nil)

	// act
	sut := controller.NewUser(uc)
	c, recorder := test.NewRequest(t, http.MethodPost, "/user/sign_up", &buf)
	assert.NoError(t, sut.SignUp(c))

	// assert
	assert.Equal(t, http.StatusCreated, recorder.Code)
	cookies := recorder.Header().Get("Set-Cookie")
	assert.NotEmpty(t, cookies)
	var out output.UserAccountSignUp
	assert.NoError(t, json.NewDecoder(recorder.Body).Decode(&out))
	expected := output.UserAccountSignUp{
		UserID:   res.Account.UserID.String(),
		Email:    res.Account.Email,
		Provider: res.Account.Provider.String(),
		UID:      res.Account.UID,
	}
	assert.Equal(t, expected, out)
}
