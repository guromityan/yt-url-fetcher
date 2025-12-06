package fetcher

import (
	"fmt"

	"github.com/guromityan/yt-url-fetcher/internal/models"
)

// GetUploadsPlaylistID retrieves the uploads playlist ID for a given channel.
func (f *Fetcher) GetUploadsPlaylistID(channelID string) (string, error) {
	call := f.service.Channels.List([]string{"contentDetails"}).Id(channelID)
	resp, err := call.Do()
	if err != nil {
		return "", fmt.Errorf("failed to fetch channel details: %w", err)
	}

	if len(resp.Items) == 0 {
		return "", fmt.Errorf("channel not found: %s", channelID)
	}

	uploadsPlaylistID := resp.Items[0].ContentDetails.RelatedPlaylists.Uploads
	if uploadsPlaylistID == "" {
		return "", fmt.Errorf("no uploads playlist found for channel: %s", channelID)
	}

	return uploadsPlaylistID, nil
}

// GetPlaylistItems retrieves all videos from a playlist.
func (f *Fetcher) GetPlaylistItems(playlistID string) ([]models.Video, error) {
	var videos []models.Video
	pageToken := ""

	for {
		call := f.service.PlaylistItems.List([]string{"snippet"}).
			PlaylistId(playlistID).
			MaxResults(50).
			PageToken(pageToken)

		resp, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("failed to fetch playlist items: %w", err)
		}

		for _, item := range resp.Items {
			// Skip private/deleted videos if necessary, but API usually handles visibility.
			// We want public videos.
			resourceID := item.Snippet.ResourceId
			if resourceID.Kind == "youtube#video" {
				videos = append(videos, models.Video{
					ID:    resourceID.VideoId,
					Title: item.Snippet.Title,
					URL:   "https://www.youtube.com/watch?v=" + resourceID.VideoId,
				})
			}
		}

		pageToken = resp.NextPageToken
		if pageToken == "" {
			break
		}
	}

	return videos, nil
}

// FetchByChannel retrieves all video URLs for a channel.
func (f *Fetcher) FetchByChannel(channelID string) ([]models.Video, error) {
	playlistID, err := f.GetUploadsPlaylistID(channelID)
	if err != nil {
		return nil, err
	}
	return f.GetPlaylistItems(playlistID)
}

// FetchByPlaylist retrieves all video URLs for a playlist.
func (f *Fetcher) FetchByPlaylist(playlistID string) ([]models.Video, error) {
	return f.GetPlaylistItems(playlistID)
}
