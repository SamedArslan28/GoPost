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
	// Attempt to load a .env file. This is useful for local development.
	// In production, environment variables are typically set directly.
	// We ignore the "file not found" error because it's not a critical issue.
	_ = godotenv.Load()

	// Create a new Config struct to hold the values.
	cfg := Config{
		DatabaseURL: os.Getenv("POSTGRES_URL"),
	}

	// Validate that required configuration values are present.
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("POSTGRES_URL environment variable must be set")
	}

	return cfg, nil
}
