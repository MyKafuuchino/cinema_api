package main

import (
	"cinema_api/config"
	"cinema_api/database"
	"cinema_api/middleware"
	"cinema_api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config.InitEnvConfig()
	database.InitDb()

	appPort := config.GlobalAppConfig.AppPort
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler(),
		JSONDecoder:  middleware.JSONDecoder(),
	})

	route.InitRouter(app)

	err := app.Listen(appPort)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
