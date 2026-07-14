# tsvsheet.go

> **A spreadsheet for plain text.** The Go implementation of [tsvsheet](https://github.com/uplang/tsvsheet): compute a `.tsvt` template (headers, formulas, sheet operations) over a `.tsv` value grid, and edit the two-file worksheet from the CLI, the browser, or the terminal — all through one engine.

This is the CLI/library implementation repo (module `github.com/uplang/tsvsheet.go`, binary `cmd/tsvsheet/`). The normative language is the grammar repo [uplang/tsvsheet](https://github.com/uplang/tsvsheet); this repo lifts its ANTLR-generated Go parser and layers the semantics. See [specs/](../specs/) — the decomposed SDD specs, the `engine-api` contract, and [ADR 0003](../specs/decisions/0003-open-semantics.md), which pins every choice made for a SPECIFICATION `[open]` item and records grammar consequences discovered while implementing.

## Architecture — one engine, three frontends

Every frontend consumes the same engine; none re-implements parsing or evaluation ([ADR 0002](../specs/decisions/0002-one-engine-thin-clients.md)).

- `src/grammar/tsvsheet` — the committed ANTLR-generated parser (package `tsvsheetgrammar`, `// Code generated … DO NOT EDIT`). Regenerate with `make go` in the grammar repo and copy `gen/go/*.go` here; excluded from the gates via [Makefile.local](../Makefile.local).
- `internal/tsvt` — the covered seam over the generated parser: source → immutable typed AST, or `constants.ErrSyntax`. No ANTLR type escapes this package.
- `internal/sheet` — the §9 processor: grid I/O, the §5 reference resolver, the §11 evaluator with error values, header/body/final phases, structural ops, `Check` diagnostics, `Explain` traces.
- `internal/session` — the one mutable editing model (the repo's sole pointer-receiver type), backing serve and tui.
- `internal/cli` — urfave/cli v3 commands: `render` `parse` `check` `explain` `serve` `tui`, unix stdin/stdout discipline.
- `internal/serve` — HTTP JSON API + embedded browser spreadsheet (`go:embed`, thin client of the session).
- `internal/tui` — bubbletea terminal editor over the same session.

## Non-negotiables

- **The engine is the single source of truth for semantics.** A new capability the web or TUI needs is added to `internal/session`/`internal/sheet` and consumed by both — never duplicated in a frontend. The ANTLR JavaScript target is deliberately unused (ADR 0002).
- **`[open]` SPECIFICATION items are decisions, not inventions.** Anything the 2011 source underspecified is pinned in [ADR 0003](../specs/decisions/0003-open-semantics.md) with a rationale; the §8 worked example is documented as internally inconsistent and is **not** a conformance target. New semantic choices extend that ADR.
- **The full gomatic Go gate applies:** `make check` green — gofumpt, vet, staticcheck, golangci (gocognit ≤ 7), govulncheck, **100% aggregate coverage** (every package, including `cmd/`; `src/grammar` excluded). Errors are `errs.Const` sentinels in `internal/constants`; no `fmt.Errorf`/`errors.New`. Value receivers except `session.Session`.
