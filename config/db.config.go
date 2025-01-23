package config

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
		DbHost:     GetEnv("DB_HOST", "127.0.0.1"),
		DbPort:     GetEnv("DB_PORT", "3306"),
		DbUser:     GetEnv("DB_USER", "root"),
		DbPassword: GetEnv("DB_PASSWORD", ""),
		DbName:     GetEnv("DB_NAME", "inventory-management"),
	}
}
