package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	GinRunMode string
}

func Load() *Config {
	return &Config{
		ServerPort: getEnvOrDefault("SERVER_PORT", "8080"),
		GinRunMode: getEnvOrDefault("GIN_RUN_MODE", "debug"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func Get() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found !!!")
	}

	cfg := Load()

	return cfg
}
