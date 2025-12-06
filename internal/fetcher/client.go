package fetcher

import (
	"context"
	"fmt"

	"github.com/guromityan/yt-url-fetcher/internal/models"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Fetcher handles interactions with the YouTube Data API.
type Fetcher struct {
	service *youtube.Service
	config  models.FetchConfig
}

// NewFetcher creates a new Fetcher instance.
func NewFetcher(ctx context.Context, config models.FetchConfig) (*Fetcher, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(config.APIKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create YouTube service: %w", err)
	}

	return &Fetcher{
		service: service,
		config:  config,
	}, nil
}
