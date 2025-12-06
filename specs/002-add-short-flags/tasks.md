---
description: "Task list for Add Short Flags feature"
---

# Tasks: Add Short Flags

**Input**: Design documents from `/specs/002-add-short-flags/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, contracts/

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Verify existing project structure and dependencies in cmd/yt-url-fetcher/main.go

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T002 Verify `flag` package usage in cmd/yt-url-fetcher/main.go allows for multiple flags binding to same variable

**Checkpoint**: Foundation ready - user story implementation can now begin

---

## Phase 3: User Story 1 - ショートフラグによるチャンネル指定 (Priority: P1)

**Goal**: `-c` フラグでチャンネルIDを指定できるようにする

**Independent Test**: `go run ./cmd/yt-url-fetcher -c <CHANNEL_ID>` が動作すること

### Implementation for User Story 1

- [x] T003 [US1] Add `-c` flag definition binding to `channelID` variable in cmd/yt-url-fetcher/main.go
- [x] T004 [US1] Update usage string for `-c` flag to indicate it is an alias for `--channel` in cmd/yt-url-fetcher/main.go

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - ショートフラグによるプレイリスト指定 (Priority: P1)

**Goal**: `-p` フラグでプレイリストIDを指定できるようにする

**Independent Test**: `go run ./cmd/yt-url-fetcher -p <PLAYLIST_ID>` が動作すること

### Implementation for User Story 2

- [x] T005 [US2] Add `-p` flag definition binding to `playlistID` variable in cmd/yt-url-fetcher/main.go
- [x] T006 [US2] Update usage string for `-p` flag to indicate it is an alias for `--playlist` in cmd/yt-url-fetcher/main.go
- [x] T007 [US2] Verify existing exclusive flag logic handles aliased flags correctly (no code change expected, just verification)

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T008 Update README.md with short flag usage examples
- [x] T009 Manual integration test: Verify `-c` works
- [x] T010 Manual integration test: Verify `-p` works
- [x] T011 Manual integration test: Verify `-c` and `-p` combined triggers error
- [x] T012 Manual integration test: Verify `--channel` and `--playlist` still work (regression test)

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies
- **Foundational (Phase 2)**: Depends on Setup completion
- **User Story 1 (Phase 3)**: Depends on Foundational completion
- **User Story 2 (Phase 4)**: Depends on Foundational completion
- **Polish (Phase 5)**: Depends on completion of all stories

### Implementation Strategy

### MVP First (User Story 1)

1. Add `-c` flag
2. Validate: Can fetch using `-c`?

### Incremental Delivery (User Story 2)

1. Add `-p` flag
2. Validate: Can fetch using `-p`?
3. Verify: Exclusive logic works?

### Parallel Opportunities

- T003 and T005 can be implemented in parallel (same file, adjacent lines).
- T004 and T006 can be implemented in parallel.
