# yt-url-fetcher

YouTubeチャンネルやプレイリストから動画URLを一括取得するCLIツール。
NotebookLMなどのツールにインポートするためのソースリスト作成に最適です。

## 特徴

- **単一責任**: YouTubeのURL取得に特化
- **APIクォータ節約**: `Search` APIを使わず、コストの低い `PlaylistItems` を使用
- **シンプル**: Go言語製のシングルバイナリ

## 前提条件

- **Go 1.21+**
- **YouTube Data API v3 Key**

## インストール

```bash
go install github.com/guromityan/yt-url-fetcher/cmd/yt-url-fetcher@latest
```

またはソースからビルド:

```bash
git clone https://github.com/guromityan/yt-url-fetcher.git
cd yt-url-fetcher
go build ./cmd/yt-url-fetcher
```

## 使い方

APIキーを環境変数に設定します。

```bash
export YOUTUBE_API_KEY="your_api_key_here"
```

### チャンネルの動画一覧を取得

```bash
./yt-url-fetcher -c <CHANNEL_ID>
# または
./yt-url-fetcher --channel <CHANNEL_ID>
```

### プレイリストの動画一覧を取得

```bash
./yt-url-fetcher -p <PLAYLIST_ID>
# または
./yt-url-fetcher --playlist <PLAYLIST_ID>
```

### 出力例

```text
https://www.youtube.com/watch?v=dQw4w9WgXcQ
https://www.youtube.com/watch?v=xyz123abcde
...
```