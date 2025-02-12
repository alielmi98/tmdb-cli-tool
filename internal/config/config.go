package config

import (
	"fmt"
	"os"
)

type Config struct {
	APIKey  string
	BaseURL string
}

func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("TMDB_API_KEY environment variable is not set")
	}

	return &Config{
		APIKey:  apiKey,
		BaseURL: "https://api.themoviedb.org/3/movie/",
	}, nil
}
