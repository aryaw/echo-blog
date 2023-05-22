package pkg

import (
	"blog/pkg/hello"
	"blog/pkg/post"
	"blog/pkg/user"
	s "blog/server"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func PublicRoute(group *echo.Group) {
	group.GET("/hello", hello.SayHello)
	group.GET("/first-post", post.FirstPost)
}

func RestrictedRoute(group *echo.Group) {
	group.GET("/hello", hello.SayHello)
}

func ConfigureRoutes(server *s.Server) {
	publicGroup := server.Echo.Group("/api")
	PublicRoute(publicGroup)

	server.Echo.Use(middleware.Logger())

	restricted := server.Echo.Group("/v1/api/")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(user.JwtCustomClaims)
		},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	restricted.Use(echojwt.WithConfig(config))
	PublicRoute(restricted)
}
