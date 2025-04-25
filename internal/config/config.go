package config

import (
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	port := getEnv("PORT", "8080")

	cfg := &Config{
		Port: port,
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
