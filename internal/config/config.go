package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	Env         string
}

func LoadConfig() *Config {
	// In a real project, you would use a library like 'viper' or 'godotenv'
	// For this boilerplate, we'll use simple os.Getenv
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Env:         getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
