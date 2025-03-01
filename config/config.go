package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration.
type Config struct {
	GeminiAPIKey string
	GeminiModel  string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	model := os.Getenv("GEMINI_MODEL")
	if apiKey == "" {
		return nil, fmt.Errorf("missing GEMINI_API_KEY environment variable")
	}
	if model == "" {
		model = "gemini-1.5-flash-latest" // Default model if not specified
	}
	return &Config{
		GeminiAPIKey: apiKey,
		GeminiModel:  model,
	}, nil
}
