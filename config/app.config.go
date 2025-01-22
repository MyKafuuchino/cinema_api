package config

import "cinema_api/utils"

type AppConfig struct {
	AppPort   string
	SecretKey string
}

var GlobalAppConfig AppConfig

func SetupAppConfig() {
	GlobalAppConfig = AppConfig{
		AppPort:   utils.GetEnv("APP_PORT", ":8080"),
		SecretKey: utils.GetEnv("APP_SECRET_KEY", "DameDameDame"),
	}
}
