# Research: Fetch YouTube URLs

## 決定事項

### 1. YouTube API クォータの最適化戦略
**Decision**: チャンネルの動画を取得する際、`Search: list`エンドポイントではなく、`Channels: list`で「アップロード済み動画（uploads）」のプレイリストIDを取得し、その後`PlaylistItems: list`を使用する。
**Rationale**: 
- `Search: list`は1リクエストあたり**100ユニット**のクォータを消費する。
- `Channels: list` (1ユニット) + `PlaylistItems: list` (1ユニット) の組み合わせなら、わずか**2ユニット**で済む。
- これにより、無料枠（1日10,000ユニット）内での実行回数を大幅に最大化できる。
**Alternatives Considered**: 
- `Search: list`: 実装は直感的だが、コストが高すぎるため却下。

### 2. CLIフラグの設計
**Decision**: Go標準の `flag` パッケージを使用する。
- `--channel <ID>`: チャンネルID指定
- `--playlist <ID>`: プレイリストID指定
- 排他的な指定を基本とするが、両方指定された場合はエラーとするか、順次処理するかを実装時に判断（今回はシンプルにエラーまたは順次処理とする）。
**Rationale**: 外部ライブラリ（Cobraなど）を導入するほどの複雑性がないため、Constitutionの「シンプルさ」に従う。

### 3. APIキーの管理
**Decision**: 環境変数 `YOUTUBE_API_KEY` から読み込む。
**Rationale**: CLI引数で渡すと履歴に残るリスクがあるため、環境変数がセキュリティ上好ましい。12-factor appの原則にも従う。
