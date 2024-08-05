package controller

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"mickamy.com/playground/internal/api/dto/input"
	"mickamy.com/playground/internal/api/dto/output"
	"mickamy.com/playground/internal/lib/cookie"
	"mickamy.com/playground/internal/lib/validator"
	"mickamy.com/playground/internal/usecase"
)

type User struct {
	signUpUseCase usecase.UserSignUp
}

func NewUser(signUpUseCase usecase.UserSignUp) User {
	return User{
		signUpUseCase: signUpUseCase,
	}
}

func (ctrl User) SignUp(c echo.Context) error {
	var in input.UserSignUp
	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, output.NewErrorMessage(err))
	}
	ctx := c.Request().Context()
	if msgs := validator.Struct(ctx, in); msgs != nil {
		return c.JSON(http.StatusBadRequest, output.NewErrorMessages(msgs))
	}
	res, err := ctrl.signUpUseCase.Do(ctx, usecase.UserSignUpInput{Provider: in.Provider, IDToken: in.IDToken})
	if err != nil {
		slog.Error("failed to sign up", "err", err)
		return c.JSON(http.StatusInternalServerError, output.ErrorDefault)
	}
	cookie.SetToken(c, res.Tokens)
	out := output.UserAccountSignUp{
		UserID:   res.Account.UserID.String(),
		Email:    res.Account.Email,
		Provider: res.Account.Provider.String(),
		UID:      res.Account.UID,
	}
	return c.JSON(http.StatusCreated, out)
}
