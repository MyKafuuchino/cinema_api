package config

import "cinema_api/helper"

type AppConfig struct {
	AppPort   string
	SecretKey string
}

var GlobalAppConfig AppConfig

func SetupAppConfig() {
	GlobalAppConfig = AppConfig{
		AppPort:   helper.GetEnv("APP_PORT", ":8080"),
		SecretKey: helper.GetEnv("APP_SECRET_KEY", "DameDameDame"),
	}
}
