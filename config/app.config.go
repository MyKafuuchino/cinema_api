package config

type AppConfig struct {
	AppPort   string
	SecretKey string
}

var GlobalAppConfig AppConfig

func SetupAppConfig() {
	GlobalAppConfig = AppConfig{
		AppPort:   GetEnv("APP_PORT", ":8080"),
		SecretKey: GetEnv("APP_SECRET_KEY", "DameDameDame"),
	}
}
