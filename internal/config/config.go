package config

import (
	"log"
	"os"
)

type Config struct {
	ServiceMode  string
	ServicePort  string
	DBNAME       string
	DBPORT       string
	DBTECH       string
	DBURL        string
	DBUSER       string
	DBPASS       string
	LogLevel     string
	CacheTECH    string
	CacheURL     string
	CachePORT    string
	CacheUSER    string
	CachePASS    string
	Pacs002CBURL string
}

var cfg Config

func init() {
	log.Println("Initializing configuration...")
	cfg = Config{
		ServiceMode:  getEnv("SERVICE_MODE", "development"),
		ServicePort:  getEnv("SERVICE_PORT", "8080"),
		DBNAME:       getEnv("DB_NAME", "paymentsim"),
		DBPORT:       getEnv("DB_PORT", "27017"),
		DBTECH:       getEnv("DB_TECH", "mongodb"),
		DBURL:        getEnv("DB_URL", "localhost"),
		DBUSER:       getEnv("DB_USER", "paymentsimuser"),
		DBPASS:       getEnv("DB_PASS", "paymentsimuser"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		CacheTECH:    getEnv("CACHE_TECH", "redis"),
		CacheURL:     getEnv("CACHE_URL", "localhost"),
		CachePORT:    getEnv("CACHE_PORT", "6379"),
		CacheUSER:    getEnv("CACHE_USER", "admin"),
		CachePASS:    getEnv("CACHE_PASS", "admin"),
		Pacs002CBURL: getEnv("PACS002_CB_URL", "http://localhost:8081/pacs002"),
	}
}

func LoadConfig() Config {
	log.Println("Configuration loaded:", cfg)
	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && len(value) > 0 {
		log.Printf("Found env var %s=%s", key, value)
		return value
	}
	log.Printf("Env var %s not set, using default=%s", key, fallback)
	return fallback
}
