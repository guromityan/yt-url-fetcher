# Quickstart: yt-url-fetcher

YouTubeチャンネルやプレイリストから動画URLを一括取得するCLIツール。

## 前提条件

- **Go 1.21+** がインストールされていること
- **YouTube Data API v3** のAPIキーを取得済みであること

## インストール

ソースコードからビルドしてインストールします。

```bash
git clone https://github.com/guromityan/yt-url-fetcher.git
cd yt-url-fetcher
go install ./cmd/yt-url-fetcher
```

## セットアップ

APIキーを環境変数に設定します。

```bash
export YOUTUBE_API_KEY="your_api_key_here"
```

## 使い方

### チャンネルの動画一覧を取得

チャンネルIDを指定して実行します。

```bash
yt-url-fetcher --channel UC_x5XG1OV2P6uZZ5FSM9Ttw
```

### プレイリストの動画一覧を取得

プレイリストIDを指定して実行します。

```bash
yt-url-fetcher --playlist PL59FEE129ADFF2B12
```

### 結果をファイルに保存

標準出力を使用するため、リダイレクトで保存可能です。

```bash
yt-url-fetcher --channel UC_x5XG1OV2P6uZZ5FSM9Ttw > urls.txt
```

### エラー処理

エラーが発生した場合（無効なID、API制限など）、標準エラー出力に詳細が表示され、終了ステータス `1` で終了します。
