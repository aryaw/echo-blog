package pkg

import (
	"blog/pkg/hello"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(group *echo.Group) {
	group.GET("/hello", hello.SayHello)
}
