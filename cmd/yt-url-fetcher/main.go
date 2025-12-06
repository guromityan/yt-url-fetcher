package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/guromityan/yt-url-fetcher/internal/fetcher"
	"github.com/guromityan/yt-url-fetcher/internal/models"
)

func main() {
	channelID := flag.String("channel", "", "YouTube Channel ID")
	playlistID := flag.String("playlist", "", "YouTube Playlist ID")
	flag.Parse()

	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "エラー: APIキーが設定されていません。環境変数 YOUTUBE_API_KEY を設定してください。")
		os.Exit(1)
	}

	if *channelID == "" && *playlistID == "" {
		fmt.Fprintln(os.Stderr, "エラー: --channel または --playlist のいずれかを指定してください。")
		flag.Usage()
		os.Exit(1)
	}

	if *channelID != "" && *playlistID != "" {
		// For now, simplistic error handling as per plan (or we could prioritize one)
		// The tasks say "Update main to handle exclusive flags error" in Phase 4 (T015).
		// So for now I can leave it or implement basic check. I'll implement basic check.
		fmt.Fprintln(os.Stderr, "エラー: --channel と --playlist は同時に指定できません。")
		os.Exit(1)
	}

	config := models.FetchConfig{
		APIKey:     apiKey,
		ChannelID:  *channelID,
		PlaylistID: *playlistID,
	}

	ctx := context.Background()
	f, err := fetcher.NewFetcher(ctx, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: Fetcherの初期化に失敗しました: %v\n", err)
		os.Exit(1)
	}

	var videos []models.Video

	if config.ChannelID != "" {
		videos, err = f.FetchByChannel(config.ChannelID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: チャンネルからの動画取得に失敗しました: %v\n", err)
			os.Exit(1)
		}
	} else if config.PlaylistID != "" {
		videos, err = f.FetchByPlaylist(config.PlaylistID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: プレイリストからの動画取得に失敗しました: %v\n", err)
			os.Exit(1)
		}
	}

	for _, v := range videos {
		fmt.Println(v.URL)
	}
}
