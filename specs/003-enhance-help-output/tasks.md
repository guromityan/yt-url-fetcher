# Tasks: Enhance Help Output

**Feature**: Enhance Help Output
**Status**: In Progress
**Spec**: [spec.md](./spec.md)

## Phase 1: Setup
*Goal: Verify environment for TDD loop.*

- [x] T001 Verify Go environment and project dependencies in `cmd/yt-url-fetcher/main.go`

## Phase 2: Foundation
*Goal: N/A (Feature is self-contained)*

## Phase 3: User Story 1 (Comprehensive Help Display)
*Goal: Users see clear Japanese help with examples when running with -h/--help.*
*Independent Test: `go run cmd/yt-url-fetcher/main.go -h` matches contract.*

### Testing (TDD)
- [x] T002 [US1] Create test file `cmd/yt-url-fetcher/main_test.go` with a test case `TestPrintUsage` that captures output and asserts Japanese content/format (Red state) in `cmd/yt-url-fetcher/main_test.go`

### Implementation
- [x] T003 [US1] Refactor `cmd/yt-url-fetcher/main.go` to extract help printing logic into a new function `printUsage(w io.Writer)`
- [x] T004 [US1] Implement `printUsage` in `cmd/yt-url-fetcher/main.go` with Japanese description, flag list, env vars, and examples to pass `TestPrintUsage`
- [x] T005 [US1] Update `main()` in `cmd/yt-url-fetcher/main.go` to override `flag.Usage` using the new `printUsage` function

## Final Phase: Polish
*Goal: Manual verification and code quality check.*

- [x] T006 [P] Run `go vet ./...` and `go fmt ./...` to ensure code quality
- [x] T007 Manual verification: Run `go run cmd/yt-url-fetcher/main.go --help` and verify output visually against `specs/003-enhance-help-output/contracts/cli-interface.md`

## Dependencies

1. **Setup** (T001)
2. **User Story 1** (T002 -> T003 -> T004 -> T005)
3. **Polish** (T006, T007)

## Parallel Execution Examples

- T006 (Linting) can be run in parallel with manual verification T007.

## Implementation Strategy

1. **TDD Approach**: Write the test for the expected output first (T002).
2. **Refactor**: Modify the main package structure to support testing (T003).
3. **Implement**: Write the actual help text (T004).
4. **Integrate**: Hook it up to the real CLI entry point (T005).