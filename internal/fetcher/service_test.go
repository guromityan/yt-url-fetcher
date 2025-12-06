package fetcher

import (
	"context"
	"testing"

	"github.com/guromityan/yt-url-fetcher/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewFetcher(t *testing.T) {
	ctx := context.Background()

	// Test with empty API key
	config := models.FetchConfig{APIKey: ""}
	f, err := NewFetcher(ctx, config)
	assert.Error(t, err)
	assert.Nil(t, f)
	assert.Equal(t, "API key is required", err.Error())

	// Test with valid API key (struct creation only, no API call)
	config = models.FetchConfig{APIKey: "dummy_key"}
	f, err = NewFetcher(ctx, config)
	assert.NoError(t, err)
	assert.NotNil(t, f)
}
