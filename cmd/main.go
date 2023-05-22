package main

import (
	application "blog"
	"blog/config"
)

func main() {
	cfg := config.NewConfig()

	application.Start(cfg)
}
