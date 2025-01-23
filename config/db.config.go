package config

import "cinema_api/helper"

type DbConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

var GlobalDbConfig DbConfig

func SetupDbConfig() {
	GlobalDbConfig = DbConfig{
		DbHost:     helper.GetEnv("DB_HOST", "127.0.0.1"),
		DbPort:     helper.GetEnv("DB_PORT", "3306"),
		DbUser:     helper.GetEnv("DB_USER", "root"),
		DbPassword: helper.GetEnv("DB_PASSWORD", ""),
		DbName:     helper.GetEnv("DB_NAME", "inventory-management"),
	}
}
