# Feature Specification: Enhance Help Output

**Feature Branch**: `003-enhance-help-output`
**Created**: 2025-12-07
**Status**: Draft
**Input**: User description: "ユーザーにもっと CLI の使い方を分かりやすく伝えるために、--help の内容を充実させたい"

## User Scenarios & Testing

### User Story 1 - Comprehensive Help Display (Priority: P1)

As a CLI user, I want to see clear usage instructions, flag descriptions, and examples when I run the help command, so that I can use the tool correctly without external documentation.

**Why this priority**: The CLI is the primary interface, and users need to know how to operate it to get value. Clear help reduces friction and support needs.

**Independent Test**: Run the binary with `--help` and verify the output structure contains description, flags, and examples.

**Acceptance Scenarios**:

1. **Given** the CLI is installed, **When** I run `yt-url-fetcher --help` or `yt-url-fetcher -h`, **Then** I see a brief description of the tool.
2. **Given** the CLI is installed, **When** I run the help command, **Then** I see a list of all available flags (`--channel`/`-c`, `--playlist`/`-p`) with descriptions.
3. **Given** the CLI is installed, **When** I run the help command, **Then** I see at least two concrete usage examples (one for channel, one for playlist).

### Edge Cases

- What happens when the user provides invalid flags along with `--help`? (Should still show help)
- How does the system handle environment variable requirements in the help text? (Should mention `YOUTUBE_API_KEY` is required)

## Requirements

### Functional Requirements

- **FR-001**: The system MUST override the default help output to provide a custom usage message.
- **FR-002**: The help message MUST include a header describing the tool's purpose (e.g., "Fetches YouTube video URLs from channels or playlists").
- **FR-003**: The help message MUST list all defined flags (`-channel`, `-c`, `-playlist`, `-p`) with clear descriptions.
- **FR-004**: The help message MUST explicitly state that the `YOUTUBE_API_KEY` environment variable is required.
- **FR-005**: The help message MUST provide a section for "Examples" showing valid command invocations.
    - Example 1: Fetching by Channel ID.
    - Example 2: Fetching by Playlist ID.

### Key Entities

N/A

## Success Criteria

### Measurable Outcomes

- **SC-001**: Help output is displayed immediately (under 100ms) when requested.
- **SC-002**: 100% of available flags are documented in the help output.
- **SC-003**: Help output contains at least 2 copy-pasteable examples.
- **SC-004**: Help output mentions the required `YOUTUBE_API_KEY`.