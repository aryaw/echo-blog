package main

import (
	"log"
	"os"

	"blog/app"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	port := os.Getenv("PORT")
	ech := echo.New()

	endPointGroup := ech.Group("/api")
	app.RegisterRoute(endPointGroup)

	log.Printf("\n\n PORT: %s \n ENV: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("API_VERSION"))
	ech.Logger.Fatal(ech.Start(":" + port))
}
