# Data Model: Fetch YouTube URLs

## Entities

### Video
YouTube上の1つの動画を表す。

| Field | Type | Description | Validation |
|-------|------|-------------|------------|
| `ID` | `string` | YouTube動画ID | 11文字の英数字 |
| `Title` | `string` | 動画タイトル | APIから取得 |
| `URL` | `string` | 動画の完全なURL | `https://www.youtube.com/watch?v=` + ID |

### FetchConfig
実行時の設定を表す。

| Field | Type | Description |
|-------|------|-------------|
| `APIKey` | `string` | YouTube Data API Key |
| `ChannelID` | `string` | (Optional) 対象チャンネルID |
| `PlaylistID` | `string` | (Optional) 対象プレイリストID |

## Data Flow

1. **Input**: CLI引数 (`--channel`, `--playlist`) & 環境変数 (`YOUTUBE_API_KEY`)
2. **Process**: 
   - `FetchConfig` 生成
   - API Client 初期化
   - (Channel指定時) Channel ID -> Uploads Playlist ID 解決
   - Playlist Items 取得 (ページネーション処理)
   - `Video` リスト生成
3. **Output**: `Video.URL` を標準出力へ書き出し
