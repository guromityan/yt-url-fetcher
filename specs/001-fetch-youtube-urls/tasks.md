---
description: "Task list for Fetch YouTube URLs feature"
---

# Tasks: Fetch YouTube URLs

**Input**: Design documents from `/specs/001-fetch-youtube-urls/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Initialize Go project and go.mod
- [x] T002 Create project directory structure (cmd/yt-url-fetcher, internal/fetcher, internal/models)
- [x] T003 Install dependencies (`google.golang.org/api/youtube/v3`, `github.com/stretchr/testify`)

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T004 Define `Video` and `FetchConfig` structs in internal/models/video.go
- [x] T005 Implement initial `Fetcher` struct and constructor in internal/fetcher/client.go
- [x] T006 Implement API key validation logic in internal/fetcher/client.go
- [x] T007 Setup main CLI entry point with flag parsing in cmd/yt-url-fetcher/main.go (skeleton only)

**Checkpoint**: Foundation ready - user story implementation can now begin

---

## Phase 3: User Story 1 - チャンネルからの動画一覧取得 (Priority: P1)

**Goal**: チャンネルIDを指定して、そのチャンネルの全公開動画URLを取得する

**Independent Test**: `go run ./cmd/yt-url-fetcher -channel <CHANNEL_ID>` で動画URLリストが出力されること

### Implementation for User Story 1

- [x] T008 [US1] Implement `GetUploadsPlaylistID` method in internal/fetcher/service.go (API call to Channels.List)
- [x] T009 [US1] Implement `GetPlaylistItems` method in internal/fetcher/service.go (API call to PlaylistItems.List with pagination)
- [x] T010 [US1] Implement `FetchByChannel` method in internal/fetcher/service.go (orchestrates T008 and T009)
- [x] T011 [US1] Integrate `FetchByChannel` logic into cmd/yt-url-fetcher/main.go for `--channel` flag
- [x] T012 [US1] Add basic unit test for `FetchByChannel` logic (mocking API responses if possible) in internal/fetcher/service_test.go

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - プレイリストからの動画一覧取得 (Priority: P2)

**Goal**: プレイリストIDを指定して、そのプレイリスト内の動画URLを取得する

**Independent Test**: `go run ./cmd/yt-url-fetcher -playlist <PLAYLIST_ID>` で動画URLリストが出力されること

### Implementation for User Story 2

- [x] T013 [US2] Implement `FetchByPlaylist` method in internal/fetcher/service.go (reuses T009 `GetPlaylistItems`)
- [x] T014 [US2] Integrate `FetchByPlaylist` logic into cmd/yt-url-fetcher/main.go for `--playlist` flag
- [x] T015 [US2] Update `main` to handle exclusive flags error (if both channel and playlist provided)
- [x] T016 [US2] Add unit test for `FetchByPlaylist` in internal/fetcher/service_test.go

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T017 Improve error messages for invalid API keys or quotas in cmd/yt-url-fetcher/main.go
- [x] T018 Verify cross-compilation capability (build for linux/amd64, windows/amd64)
- [x] T019 Update README.md (or verify quickstart.md is sufficient)
- [x] T020 Final integration test (manual or scripted) verifying output format against `contracts/cli-interface.md`

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies
- **Foundational (Phase 2)**: Depends on Setup completion
- **User Story 1 (Phase 3)**: Depends on Foundational completion
- **User Story 2 (Phase 4)**: Depends on Foundational completion (and conceptually reuses logic from US1 components like `GetPlaylistItems`, though task T013 is distinct)
- **Polish (Phase 5)**: Depends on completion of all stories

### Implementation Strategy

### MVP First (User Story 1)

1. Complete Setup & Foundation
2. Implement Channel Fetching logic (`GetUploadsPlaylistID` -> `GetPlaylistItems`)
3. Validate: Can fetch from a channel?

### Incremental Delivery (User Story 2)

1. Add Playlist Fetching logic (reuse `GetPlaylistItems`)
2. Validate: Can fetch from a playlist?
3. Verify: Channel fetching still works?

### Parallel Opportunities

- T004, T005, T006, T007 within Phase 2 are largely sequential due to code dependencies, but T007 (CLI skeleton) could start early.
- T012 (Tests) can be written in parallel with T010/T011 once interfaces are defined.
