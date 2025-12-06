package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/guromityan/yt-url-fetcher/internal/fetcher"
	"github.com/guromityan/yt-url-fetcher/internal/models"
)

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "YouTubeのチャンネルIDまたはプレイリストIDから動画のURLリストを取得します。")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "使用法:")
	fmt.Fprintln(w, "  yt-url-fetcher [フラグ]")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "フラグ:")
	flag.VisitAll(func(f *flag.Flag) {
		// Only print flags that we have defined.
		// Since we have short/long aliases sharing the same variable, flag pkg registers both.
		// We can print them as is.
		var s string
		if len(f.Name) > 1 {
			s = fmt.Sprintf("  -%s string", f.Name)
		} else {
			s = fmt.Sprintf("  -%s string", f.Name)
		}
		// Align descriptions
		if len(s) <= 4 { // space for short flags
			s += "\t"
		} else {
			s += "\n    \t"
		}
		s += strings.ReplaceAll(f.Usage, "\n", "\n    \t")
		fmt.Fprintln(w, s)
	})
	fmt.Fprintln(w)
	fmt.Fprintln(w, "環境変数:")
	fmt.Fprintln(w, "  YOUTUBE_API_KEY  (必須) YouTube Data API v3 のAPIキー")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "例:")
	fmt.Fprintln(w, "  # チャンネルIDで取得")
	fmt.Fprintln(w, "  yt-url-fetcher -c UCxxxxxxxxxxxx")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "  # プレイリストIDで取得")
	fmt.Fprintln(w, "  yt-url-fetcher -p PLxxxxxxxxxxxx")
}

func main() {
	var channelID string
	var playlistID string

	flag.StringVar(&channelID, "channel", "", "YouTube Channel ID")
	flag.StringVar(&channelID, "c", "", "YouTube Channel ID (alias for --channel)")
	flag.StringVar(&playlistID, "playlist", "", "YouTube Playlist ID")
	flag.StringVar(&playlistID, "p", "", "YouTube Playlist ID (alias for --playlist)")

	flag.Usage = func() {
		printUsage(flag.CommandLine.Output())
	}
	flag.Parse()

	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "エラー: APIキーが設定されていません。環境変数 YOUTUBE_API_KEY を設定してください。")
		os.Exit(1)
	}

	if channelID == "" && playlistID == "" {
		fmt.Fprintln(os.Stderr, "エラー: --channel (-c) または --playlist (-p) のいずれかを指定してください。")
		flag.Usage()
		os.Exit(1)
	}

	if channelID != "" && playlistID != "" {
		fmt.Fprintln(os.Stderr, "エラー: --channel (-c) と --playlist (-p) は同時に指定できません。")
		os.Exit(1)
	}

	config := models.FetchConfig{
		APIKey:     apiKey,
		ChannelID:  channelID,
		PlaylistID: playlistID,
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
