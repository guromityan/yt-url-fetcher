# Feature Specification: Add Short Flags

**Feature Branch**: `002-add-short-flags`  
**Created**: 2025-12-06  
**Status**: Draft  
**Input**: User description: "--channel は -c でも指定可能に、--playlist は -p でも指定可能にして"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - ショートフラグによるチャンネル指定 (Priority: P1)

ユーザーは、長い `--channel` フラグの代わりに、短い `-c` フラグを使用してチャンネルIDを指定できる。これにより、タイプ量を減らし、より素早くコマンドを実行できる。

**Why this priority**: CLIツールのユーザビリティ向上のため。

**Independent Test**: `yt-url-fetcher -c <CHANNEL_ID>` を実行し、`--channel` 指定時と同じ結果が得られることを確認する。

**Acceptance Scenarios**:

1. **Given** 有効なチャンネルID, **When** `-c` フラグで実行, **Then** 動画URL一覧が出力される
2. **Given** 無効なチャンネルID, **When** `-c` フラグで実行, **Then** 適切なエラーメッセージが表示される

---

### User Story 2 - ショートフラグによるプレイリスト指定 (Priority: P1)

ユーザーは、長い `--playlist` フラグの代わりに、短い `-p` フラグを使用してプレイリストIDを指定できる。

**Why this priority**: チャンネル同様、頻繁に使用されるオプションの入力を簡略化するため。

**Independent Test**: `yt-url-fetcher -p <PLAYLIST_ID>` を実行し、`--playlist` 指定時と同じ結果が得られることを確認する。

**Acceptance Scenarios**:

1. **Given** 有効なプレイリストID, **When** `-p` フラグで実行, **Then** 動画URL一覧が出力される
2. **Given** `-c` と `-p` の同時指定, **When** 実行, **Then** 排他制御エラーが表示される

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: システムは、引数 `-c` を `--channel` のエイリアスとして受け入れなければならない (MUST)。
- **FR-002**: システムは、引数 `-p` を `--playlist` のエイリアスとして受け入れなければならない (MUST)。
- **FR-003**: ショートフラグ使用時の挙動は、ロングフラグ使用時と完全に同一でなければならない (MUST)。これには排他制御エラーのチェックも含まれる。
- **FR-004**: ヘルプメッセージ (`-h` / `--help`) にショートフラグの記述を含める (SHOULD)。

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 既存の `--channel`, `--playlist` を使用しているスクリプト等が、変更後も問題なく動作すること（後方互換性の維持）。
- **SC-002**: `-c`, `-p` を使用して、正常系および異常系の操作が期待通りに行えること。