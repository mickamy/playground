package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"mickamy.com/playground/internal/lib/jwt"
)

const (
	accessTokenName  = "access_token"
	refreshTokenName = "refresh_token"
)

func SetToken(c echo.Context, tokens jwt.AccessAndRefresh) {
	setCookie(c, accessTokenName, tokens.Access)
	setCookie(c, refreshTokenName, tokens.Refresh)
}

func setCookie(c echo.Context, name string, token jwt.Token) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token.Value
	cookie.Expires = token.ExpiresAt
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)
}

func RemoveToken(c echo.Context) {
	RemoveAccessToken(c)
	RemoveRefreshToken(c)
}

func RemoveAccessToken(c echo.Context) {
	removeCookie(c, accessTokenName)
}

func RemoveRefreshToken(c echo.Context) {
	removeCookie(c, refreshTokenName)
}

func removeCookie(c echo.Context, name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func AccessToken(c echo.Context) (string, error) {
	cookie, err := c.Cookie(accessTokenName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func RefreshToken(c echo.Context) (string, error) {
	cookie, err := c.Cookie(refreshTokenName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
