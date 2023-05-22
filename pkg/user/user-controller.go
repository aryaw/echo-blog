package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, /v1/api/ World!")
}
