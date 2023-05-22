package application

import (
	"log"

	"blog/config"
	"blog/pkg"
	"blog/server"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	pkg.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
