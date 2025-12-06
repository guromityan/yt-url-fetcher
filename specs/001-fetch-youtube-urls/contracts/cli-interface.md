# CLI Interface Contract

## Usage

```bash
yt-url-fetcher [flags]
```

## Flags

| Flag | Short | Description | Required | Example |
|------|-------|-------------|----------|---------|
| `--channel` | `-c` | YouTubeチャンネルID | Conditional* | `UC_x5XG1OV2P6uZZ5FSM9Ttw` |
| `--playlist` | `-p` | YouTubeプレイリストID | Conditional* | `PL59FEE129ADFF2B12` |
| `--help` | `-h` | ヘルプメッセージを表示 | No | |

\* `--channel` または `--playlist` のいずれか1つが必須。両方指定された場合はエラーとする。

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `YOUTUBE_API_KEY` | Google Cloud Consoleで発行したYouTube Data API v3のキー | Yes |

## Output

### Standard Output (stdout)
成功時、動画URLのリストを改行区切りで出力する。余計なヘッダーやフッターは含めない。

```text
https://www.youtube.com/watch?v=dQw4w9WgXcQ
https://www.youtube.com/watch?v=xyz123abcde
...
```

### Standard Error (stderr)
エラー発生時、日本語のエラーメッセージを出力する。

```text
エラー: APIキーが設定されていません。環境変数 YOUTUBE_API_KEY を設定してください。
```

or

```text
エラー: 指定されたチャンネル (ID: xxx) は見つかりませんでした。
```

## Exit Codes

| Code | Description |
|------|-------------|
| `0` | 成功 |
| `1` | 一般的なエラー（入力エラー、APIエラー等） |
