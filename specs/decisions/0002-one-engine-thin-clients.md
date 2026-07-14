# ADR 0002: one Go engine; web and TUI are thin clients of the same session

## Status

Accepted (2026-07-14).

## Context

The product needs the same spreadsheet capabilities in three frontends: scriptable CLI, a browser spreadsheet served by the CLI, and a terminal TUI. The obvious web approach — an ANTLR-JavaScript parser plus a JS evaluator — would create a second implementation of the language semantics that must be kept bug-for-bug identical with the Go one.

## Decision

There is exactly one implementation of parsing and evaluation: the Go engine (`internal/tsvt` + `internal/sheet`), wrapped by one stateful editing model (`internal/session`). The web page served by `tsvsheet serve` is a thin client that round-trips every edit through the JSON API (the serving process is always local, so latency is negligible); the TUI drives the identical `Session` in-process. The ANTLR JavaScript target is therefore **not used** in v1.

## Consequences

- Feature parity between web and TUI is structural, not aspirational: both render `Session.Snapshot()` and issue the same mutations.
- The browser cannot compute offline; that is acceptable because the page only exists while the local server runs.
- If client-side compute is ever wanted, the sanctioned path is compiling this same engine to WASM — never a parallel JS engine. The lifted ANTLR-JS parser remains available in the grammar repo for editor tooling (syntax highlighting), which needs no semantics.
