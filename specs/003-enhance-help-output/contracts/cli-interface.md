# CLI Interface Contract: Help Output

## Command
`yt-url-fetcher --help`
`yt-url-fetcher -h`

## Output Stream
`STDERR`

## Output Format (Japanese)

```text
YouTubeのチャンネルIDまたはプレイリストIDから動画のURLリストを取得します。

使用法:
  yt-url-fetcher [フラグ]

フラグ:
  -channel string, -c string
    	YouTube チャンネルID
  -playlist string, -p string
    	YouTube プレイリストID

環境変数:
  YOUTUBE_API_KEY  (必須) YouTube Data API v3 のAPIキー

例:
  # チャンネルIDで取得
  yt-url-fetcher -c UCxxxxxxxxxxxx

  # プレイリストIDで取得
  yt-url-fetcher -p PLxxxxxxxxxxxx
```

## Behavior
- The process MUST exit with status code 0 after displaying help (handled by `flag` package default behavior, or manual if customized extensively, but standard `flag` behavior is preferred for exit code).
- Note: `flag.Usage` usually doesn't exit itself, `flag.Parse` does.
