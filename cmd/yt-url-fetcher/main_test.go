package main

import (
	"bytes"
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintUsage(t *testing.T) {
	// Save original CommandLine and restore it after test
	origCommandLine := flag.CommandLine
	defer func() { flag.CommandLine = origCommandLine }()

	// Create a new FlagSet for testing
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)

	// Define flags expected by printUsage
	var channelID, playlistID string
	flag.StringVar(&channelID, "channel", "", "YouTube Channel ID")
	flag.StringVar(&channelID, "c", "", "YouTube Channel ID (alias for --channel)")
	flag.StringVar(&playlistID, "playlist", "", "YouTube Playlist ID")
	flag.StringVar(&playlistID, "p", "", "YouTube Playlist ID (alias for --playlist)")

	// Capture the output
	var buf bytes.Buffer
	printUsage(&buf)

	output := buf.String()

	// Check for Japanese description
	assert.Contains(t, output, "YouTubeのチャンネルIDまたはプレイリストIDから動画のURLリストを取得します。")

	// Check for Usage section
	assert.Contains(t, output, "使用法:")

	// Check for Flags section
	assert.Contains(t, output, "フラグ:")
	assert.Contains(t, output, "-channel")
	assert.Contains(t, output, "-playlist")

	// Check for Environment Variable section
	assert.Contains(t, output, "環境変数:")
	assert.Contains(t, output, "YOUTUBE_API_KEY")

	// Check for Examples section
	assert.Contains(t, output, "例:")
	assert.Contains(t, output, "yt-url-fetcher -c UCxxxxxxxxxxxx")
}
