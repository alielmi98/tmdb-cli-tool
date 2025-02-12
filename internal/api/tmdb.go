package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alielmi98/tmdb-cli-tool/internal/config"
	"github.com/alielmi98/tmdb-cli-tool/internal/models"
)

// FetchMovies fetches movies from the TMDB API based on the specified type
func FetchMovies(movieType string) (*models.MovieRes, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s?api_key=%s", cfg.BaseURL, movieType, cfg.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch movies: %s", resp.Status)
	}
	resByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var moviesRes models.MovieRes

	err = json.Unmarshal(resByte, &moviesRes)
	if err != nil {
		return nil, err
	}

	return &moviesRes, nil
}
