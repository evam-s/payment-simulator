package config

import (
	"os"
)

type Config struct {
	Port  string
	DBUrl string
}

func LoadConfig() *Config {
	return &Config{
		Port:  getEnv("PORT", "8080"),
		DBUrl: getEnv("DB_URL", "mongodb://localhost:27017"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
