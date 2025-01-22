package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func InitEnvConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	SetupAppConfig()
	SetupDbConfig()
}
