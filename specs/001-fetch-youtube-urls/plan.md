# Implementation Plan: Fetch YouTube URLs

**Branch**: `001-fetch-youtube-urls` | **Date**: 2025-12-06 | **Spec**: [specs/001-fetch-youtube-urls/spec.md](../spec.md)
**Input**: Feature specification from `/specs/001-fetch-youtube-urls/spec.md`

## Summary

ユーザー指定のYouTubeチャンネルまたはプレイリストから、公開動画のURL一覧をテキスト形式で取得するCLIツールを実装する。YouTube Data API v3を使用し、コスト効率の良い「Uploadsプレイリスト」経由での取得を採用する。

## Technical Context

**Language/Version**: Go 1.21+ (推奨)
**Primary Dependencies**: `google.golang.org/api/youtube/v3` (Official Google Client)
**Storage**: N/A (Stateless CLI)
**Testing**: `testing` (Standard library), `github.com/stretchr/testify` (Optional, for assertions)
**Target Platform**: macOS, Linux, Windows (Cross-compiled CLI)
**Project Type**: CLI Tool
**Performance Goals**: 100件の動画取得を数秒以内（APIレイテンシ依存）
**Constraints**: API Quotaの最小化（Search APIの回避）
**Scale/Scope**: 小規模CLIツール

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **I. 単一責任の原則**: ✅ URL取得のみに特化する。
- **II. 標準入出力とCLIの重視**: ✅ 入力は引数、出力は標準出力、エラーは標準エラー出力。
- **III. 日本語優先**: ✅ ヘルプ、エラーメッセージ等は日本語。
- **IV. テスト駆動**: ✅ 単体テストと統合テストを計画に含める。
- **V. シンプルさ**: ✅ 複雑なフレームワークは使わず、標準ライブラリと公式クライアントを中心に構成。

## Project Structure

### Documentation (this feature)

```text
specs/001-fetch-youtube-urls/
├── plan.md              # This file
├── research.md          # API usage strategy
├── data-model.md        # Video entity definition
├── quickstart.md        # Installation & Usage
├── contracts/           
│   └── cli-interface.md # CLI arguments & output format
└── tasks.md             # Implementation tasks
```

### Source Code (repository root)

```text
# Go Standard Layout
.
├── go.mod
├── go.sum
├── cmd/
│   └── yt-url-fetcher/
│       └── main.go      # Entry point, flag parsing
├── internal/
│   ├── fetcher/         # Core logic package
│   │   ├── client.go    # YouTube API wrapper
│   │   └── service.go   # Business logic (Channel vs Playlist)
│   └── models/          # Data structures
└── tests/
    ├── integration/     # End-to-end CLI tests
    └── unit/            # Logic tests
```

**Structure Decision**: Goの標準的なディレクトリ構成（Standard Go Project Layout）を採用。CLIのエントリーポイントを`cmd/`に、ロジックを`internal/`に配置してカプセル化する。

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| N/A | | |