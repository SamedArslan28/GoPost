package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application.
// It's a central place to manage all environment-specific values.
type Config struct {
	DatabaseURL string
}

// LoadConfig reads configuration from a .env file and the environment.
func LoadConfig() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		DatabaseURL: os.Getenv("POSTGRES_URL"),
	}

	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("POSTGRES_URL environment variable must be set")
	}

	return cfg, nil
}
