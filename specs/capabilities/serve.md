# serve — web spreadsheet over the session API

## Goal

`tsvsheet serve --template sheet.tsvt --data sheet.tsv` hosts a local web spreadsheet: an editable grid view of the computed sheet that manages the `.tsvt`/`.tsv` files through the one shared engine.

## Requirements

- R1: HTTP JSON API exactly as fixed in [contracts/engine-api.md](../contracts/engine-api.md) (`/api/state`, `/api/template`, `/api/data/cell`, `/api/save`, `/api/explain`), backed by one `session.Session` guarded for concurrent access.
- R2: The embedded UI (`go:embed`, self-contained — no CDN) renders: computed grid with column letters and row numbers, per-cell editing (data cells edit the `.tsv`; a formula bar shows/edits the template), a template text panel, syntax errors surfaced inline (422 body), save button, and `#REF!`-class error styling.
- R3: An edit round-trips: PUT → recompute → fresh `State` → repaint; latency budget is one local HTTP hop.
- R4: `--host`/`--port`/`--shutdown-timeout` flags per the template's serve idiom (go-httpserver, graceful shutdown); startup logs the URL.
- R5: Nothing is written to disk except on explicit `/api/save`.

## Acceptance Criteria

- httptest-driven tests cover every endpoint: state shape, template edit success + 422 syntax error, cell edit, save (to a temp worksheet), explain; 100% coverage on handler and domain code.
- Manual smoke: `tsvsheet serve` on the worked example, edit a cell in the browser, save, observe the `.tsv` change.
