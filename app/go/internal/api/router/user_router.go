package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"mickamy.com/playground/internal/api/controller"
	"mickamy.com/playground/internal/lib/oauth"
	"mickamy.com/playground/internal/repository"
	"mickamy.com/playground/internal/usecase"
)

func User(e *echo.Echo, db *gorm.DB) {
	auth := oauth.New()
	userRepo := repository.NewUser(db)
	accountRepo := repository.NewUserAccount(db)
	profileRepo := repository.NewUserProfile(db)
	avatarRepo := repository.NewUserAvatar(db)
	signUpUseCase := usecase.NewUserSignUp(db, auth, userRepo, accountRepo, profileRepo, avatarRepo)
	ctrl := controller.NewUser(
		signUpUseCase,
	)
	e.POST("/user/sign_up", ctrl.SignUp)
}
