# Research: Enhance Help Output

**Feature**: Enhance Help Output
**Date**: 2025-12-07

## Unknowns & Clarifications

### 1. How to test `flag.Usage` without `os.Exit`?

**Problem**: `flag.Parse()` handles `-h` by printing usage and calling `os.Exit(0)`. This makes unit testing the full flow difficult.
**Decision**: Refactor the help printing logic into a standalone function `printUsage(w io.Writer)` or similar.
**Rationale**:
- Allows passing a `bytes.Buffer` in tests to capture and verify output.
- `main()` can assign `flag.Usage = func() { printUsage(flag.CommandLine.Output()) }`.
- Keeps `main` clean.

### 2. Output Destination

**Problem**: Should help go to `stdout` or `stderr`?
**Decision**: `stderr` (Go default for `flag.Usage`).
**Rationale**: Standard practice. Help text is not the "data" result of the tool.

### 3. Japanese Localization

**Problem**: Ensure compliance with Constitution Principle III.
**Decision**: Hardcode Japanese strings in `printUsage`.
**Rationale**: No requirement for multi-language support yet. YAGNI.

## Technology Decisions

| Choice | Rationale | Alternatives Considered |
|--------|-----------|-------------------------|
| `flag` (stdlib) | Project already uses it. Sufficient for current needs. | `spf13/cobra`, `urfave/cli` (Too heavy/complex for this simple change) |
| `text/template` | Not needed. Simple `fmt.Fprintf` is sufficient and faster/simpler for static text. | `fmt.Println` (Too many calls), `text/template` (Overkill) |

## Best Practices

- **Separation of Concerns**: Isolate the help text generation from the exit mechanism.
- **Consistency**: Match the style of existing error messages (though help is informational).
