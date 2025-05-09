package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	GinRunMode      string
	LogSaveName     string
	LogFileExt      string
	RuntimeRootPath string
	TimeFormat      string
	LogSavePath     string
}

func Load() *Config {
	return &Config{
		ServerPort:      getEnvOrDefault("SERVER_PORT", "8080"),
		GinRunMode:      getEnvOrDefault("GIN_RUN_MODE", "debug"),
		LogSaveName:     getEnvOrDefault("LOG_SAVE_NAME", "log"),
		LogFileExt:      getEnvOrDefault("LOG_FILE_EXT", "log"),
		RuntimeRootPath: getEnvOrDefault("RUNTIME_ROOT_PATH", "runtime/"),
		TimeFormat:      getEnvOrDefault("TIME_FORMAT", "20010101"),
		LogSavePath:     getEnvOrDefault("LOG_SAVE_PATH", "logs/"),
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
