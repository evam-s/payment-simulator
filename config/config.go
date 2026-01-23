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
		DBUrl: getEnv("DB_URL", "postgres://user:pass@localhost:5432/db"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}