package models

// Video represents a single YouTube video.
type Video struct {
	ID    string
	Title string
	URL   string
}

// FetchConfig holds the runtime configuration.
type FetchConfig struct {
	APIKey     string
	ChannelID  string
	PlaylistID string
}
