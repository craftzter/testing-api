package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// APP_PORT=8080
// APP_HOST=localhost
// DB_USER=root
// DB_PASSWORD=secret
// DB_NAME=testing

type Config struct {
	AppPort    string
	AppHost    string
	DBUser     string
	DBPassword string
	DBName     string
}

// function to make sure server can run if the env not set enought
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// function to load config from env
func LoadingConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading file: %v", err)
	}
	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		AppHost:    getEnv("APP_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "sevret"),
		DBName:     getEnv("DB_NAME", "testing"),
	}, nil
}
