# ADR 0001: `tsvsheet.go` is the Go implementation repo; the parser is lifted, generated, committed

## Status

Accepted (2026-07-14).

## Context

[uplang/tsvsheet](https://github.com/uplang/tsvsheet) is grammar-only: `TsvsheetLexer.g4` + `TsvsheetParser.g4` are the normative language definition, its Makefile generates every target language into an ignored `gen/`, and its own docs say to "lift a generated tree into a new `tsvsheet.<lang>` implementation repo" — the same model as `up.grammar` → `up.go`.

## Decision

- The Go implementation lives in `uplang/tsvsheet.go` (module `github.com/uplang/tsvsheet.go` — the toolchain accepts the dotted final element, and `gomatic/template.cli` is the same-shape precedent), seeded from `gomatic/template.cli`, class CLI (`go.mod` + `cmd/tsvsheet/`).
- The generated parser is produced in the grammar repo (`make go`, Docker-isolated ANTLR) and committed here under `src/grammar/tsvsheet/` as package `tsvsheetgrammar`, so plain `go build`/`go test`/CI stay Docker-free. Regeneration is: `make go` in `uplang/tsvsheet`, copy `gen/go/*.go` over `src/grammar/tsvsheet/`.
- `src/grammar` is excluded from the coverage gate via `Makefile.local` (`COVER_PKGS`); its files carry the `Code generated … DO NOT EDIT` marker that excludes them from the file-list gates.

## Alternatives considered

- Implementing inside `uplang/tsvsheet`: rejected — that repo's charter forbids committed implementations and keeps the grammar language-neutral.
- Generating the parser at build time in this repo: rejected — it would make every build depend on Docker+Java, contrary to the fleet standard of committing generated code.
