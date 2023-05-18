package app

import (
	"blog/app/hello"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(group *echo.Group) {
	group.GET("/hello", hello.SayHello)
}
