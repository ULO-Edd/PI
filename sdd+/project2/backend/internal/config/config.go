package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Port  string
	DBUrl string
}

// Load loads configuration from .env file or environment variables
func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Port:  getEnv("PORT", "8080"),
		DBUrl: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/events?sslmode=disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
