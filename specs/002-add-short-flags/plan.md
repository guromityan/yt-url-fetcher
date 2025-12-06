# Implementation Plan: Add Short Flags

**Branch**: `002-add-short-flags` | **Date**: 2025-12-06 | **Spec**: [specs/002-add-short-flags/spec.md](../spec.md)
**Input**: Feature specification from `/specs/002-add-short-flags/spec.md`

## Summary

CLI引数のパースロジックを更新し、`--channel` の代わりに `-c`、`--playlist` の代わりに `-p` を使用可能にする。既存の機能と排他制御ロジックはそのまま維持する。

## Technical Context

**Language/Version**: Go 1.21+
**Primary Dependencies**: `google.golang.org/api/youtube/v3`, `flag` (Go Standard Library)
**Storage**: N/A
**Testing**: `testing`, `github.com/stretchr/testify` (Optional)
**Target Platform**: macOS, Linux, Windows (Cross-compiled CLI)
**Project Type**: CLI Tool
**Performance Goals**: N/A
**Constraints**: 既存のフラグ解析ロジックへの影響を最小限にする。
**Scale/Scope**: 小規模な修正

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **I. 単一責任の原則**: ✅ 影響なし。
- **II. 標準入出力とCLIの重視**: ✅ CLIの利便性向上に貢献。
- **III. 日本語優先**: ✅ ヘルプメッセージの更新が必要（日本語維持）。
- **IV. テスト駆動**: ✅ E2Eテストでショートフラグの動作を確認する。
- **V. シンプルさ**: ✅ 標準の `flag` パッケージの機能をそのまま利用する（変数の共有）。

## Project Structure

### Documentation (this feature)

```text
specs/002-add-short-flags/
├── plan.md              # This file
├── research.md          # Flag aliasing strategy
├── data-model.md        # N/A (No data model changes)
├── quickstart.md        # Updated usage examples
├── contracts/           
│   └── cli-interface.md # Updated CLI contract
└── tasks.md             # Implementation tasks
```

### Source Code (repository root)

```text
.
├── cmd/
│   └── yt-url-fetcher/
│       └── main.go      # Update flag definitions
```

**Structure Decision**: 既存の `main.go` のフラグ定義部分のみを変更するため、ディレクトリ構造への影響はない。

## Complexity Tracking

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| N/A | | |