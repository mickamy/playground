package test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func NewRequest(t *testing.T, method string, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	t.Helper()

	e := echo.New()
	req := httptest.NewRequest(method, path, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	c := e.NewContext(req, recorder)

	return c, recorder
}
