package application

import (
	"log"

	"blog/config"
	"blog/migrations"
	"blog/pkg"
	"blog/server"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	pkg.ConfigureRoutes(app)
	migrations.RunMigrations(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
