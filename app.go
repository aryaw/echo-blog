package application

import (
	"log"

	"blog/config"
	"blog/server"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
