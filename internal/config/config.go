package config

import (
	"log"
	"os"
)

type Config struct {
	ServiceMode string
	ServicePort string
	DBNAME      string
	DBPORT      string
	DBTECH      string
	DBURL       string
	LogLevel    string
}

func LoadConfig() *Config {
	log.Println("Loading configuration...")

	cfg := &Config{
		ServicePort: getEnv("SERVICE_PORT", "8080"),
		DBNAME:      getEnv("DB_NAME", "paymentsim"),
		DBPORT:      getEnv("DB_PORT", "27017"),
		DBTECH:      getEnv("DB_TECH", "mongodb"),
		DBURL:       getEnv("DB_URL", "localhost"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		ServiceMode: getEnv("SERVICE_MODE", "development"),
	}

	log.Println("Configuration loaded:", cfg)

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		log.Printf("Found env var %s=%s", key, value)
		return value
	}
	log.Printf("Env var %s not set, using default=%s", key, fallback)
	return fallback
}
