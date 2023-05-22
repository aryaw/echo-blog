package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FirstPost(c echo.Context) error {
	return c.String(http.StatusOK, "Labore elit quis adipisicing aute laborum est exercitation laborum nulla aute aliqua dolore. Enim proident pariatur occaecat et ad irure esse sit anim. Amet ullamco culpa aliqua elit officia in nulla qui nulla sit.")
}
