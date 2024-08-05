package router

import (
	"github.com/labstack/echo/v4"

	"mickamy.com/playground/internal/api/controller"
)

func User(e *echo.Echo, ctrl controller.User) {
	e.POST("/user/sign_up", ctrl.SignUp)
}
