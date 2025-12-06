# CLI Interface Contract (Updated)

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

\* チャンネルID (`--channel` or `-c`) または プレイリストID (`--playlist` or `-p`) のいずれか1つが必須。

## Notes

- `-c` と `-p` はそれぞれ `--channel` と `--playlist` のエイリアスとして機能する。
- 同じ種類のフラグを重複指定した場合（例: `-c ID1 --channel ID2`）、標準 `flag` パッケージの仕様により、最後に指定されたものが有効になる（エラーにはならない）。
- 異なる種類のフラグを同時指定した場合（例: `-c ID1 -p ID2`）、排他制御エラーとなる。
