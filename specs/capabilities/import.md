# import — content-typed HTTP import of external values

## Goal

Let a `.tsvt` pull external data and the results of remote computation into the grid — the weather, a public IP or location, a value at a URL, the output of a program running an algorithm the formula language cannot express — through a **tsvsheet-native content-type handshake**, values only, off by default, without ever executing a local command or evaluating remote code. Design decisions (the handshake, the five functions, the values-only rule, the injected `Fetcher`, the `#IMPORT!` error, the refresh model, and the trust gate) are pinned in [ADR 0006](../decisions/0006-content-typed-import.md). This is the tsvsheet-native realization of the network/import capability [ADR 0004](../decisions/0004-excel-sheets-functions.md) deferred; the Google/Excel network functions (`IMPORTHTML`, `IMPORTXML`, `WEBSERVICE`, …) remain out of scope.

## Requirements

- **R1 — five granularity functions.** `IMPORTCELL`, `IMPORTROW`, `IMPORTCOLUMN`, `IMPORTRANGE`, `IMPORTSHEET`, each taking a single URL argument, each sending its fixed media type as the request `Accept` header, and each returning the engine shape in [ADR 0006 §2](../decisions/0006-content-typed-import.md) (scalar / 1×N cellset / N×1 cellset / spilling R×C array / embeddable nested grid). Ordinary function calls over the existing grammar — no grammar change (ADR 0006 §2).
- **R2 — two-sided handshake.** The request carries `Accept: <the granularity's media type>`; the response is ingested **only if** its `Content-Type` equals that type. A mismatch is `#IMPORT!`. Media types are the RFC 6838 vendor tree with the `+tsv` suffix: `application/vnd.tsvsheet+tsv` (sheet) and `application/vnd.tsvsheet.{cell,row,column,range}+tsv` (ADR 0006 §1).
- **R3 — values only.** The response body is a `.tsvt` fragment parsed by the engine's existing TSV reader into **values, never formulas**; a leading `=` enters as literal text. No remote code execution, no fetch recursion, no cross-URL cycle detection. Shape is strict: a mismatch (a multi-cell `cell`, a ragged `range`, …) is `#IMPORT!` (ADR 0006 §3).
- **R4 — injected `Fetcher`.** The engine takes a `Fetcher` collaborator through the compute pass (mirroring `Loader`); a `nil` `Fetcher` (the default) resolves every import to `#IMPORT!`. The engine imports no `net/http`. The `Fetcher` owns the cross-pass cache. Tests inject a mock (ADR 0006 §4).
- **R5 — error model.** A new `#IMPORT!` error value ([internal/sheet/value.go](../../internal/sheet/value.go), added to the `parseError` allow-set) is the single cell-level outcome of every import failure; `explain` carries the URL, HTTP status, content type received vs. expected, byte count, and duration. The `Fetcher`'s failure reasons are `errs.Const` sentinels in [internal/constants](../../internal/constants) (ADR 0006 §5).
- **R6 — refresh model.** Imports are **not** in `IsVolatile` and **never ride the isnow ticker**. A frontend-owned cache keyed by `(url, media type)` holds values across passes; ordinary and clock-ticker recomputes read the cache with no network I/O. A separate `HasImports` predicate reports import presence. Refresh is explicit: `render` one-shot, a TUI key, a `serve` refresh action — each decoupled from the clock (ADR 0006 §6).
- **R7 — trust gate.** Off by default behind a frontend `--allow-import` flag; a host allowlist supplied only by the operator (never from the `.tsvt`, never self-enabling); `https` only with TLS verification; redirects followed within a cap and re-checked against the allowlist per hop; per-import timeout and response-size cap (with grid imports also bounded by `ResultCells`) (ADR 0006 §7).
- **R8 — per-frontend gating.** `render`/`explain`/`tui` honor the flag and allowlist; `serve` additionally requires a **non-empty allowlist** and keeps the hard non-loopback refusal, bounding SSRF to allowlisted hosts and values-only responses (ADR 0006 §8).
- **R9 — coverage as contract.** Every function and every failure path (disabled, denied host, non-2xx, content-type mismatch, shape mismatch, oversize, timeout, redirect-escape) has a table-driven contract test against a mock `Fetcher`, asserting the specific `#IMPORT!` outcome and the `explain` detail. The 100% aggregate-coverage gate holds; `errs.Const` sentinels only; value receivers except `session.Session`.

## Acceptance Criteria

- Each of the five functions computes its representative result from a mock endpoint serving the matching media type, and returns `#IMPORT!` for each documented failure, asserted in tests.
- A response whose `Content-Type` does not equal the requested media type is refused (`#IMPORT!`), and a leading-`=` value is ingested as literal text, never evaluated — both asserted.
- With no `--allow-import` (nil `Fetcher`), every import is `#IMPORT!`; with the flag on but the host not allowlisted, the import is `#IMPORT!` and `explain` names the denied host — both asserted.
- A clock-ticker recompute of a sheet containing both `NOW()` and an import updates the clock value and performs **no** re-fetch (the import is served from cache); an explicit refresh re-fetches — both asserted (mock `Fetcher` call-count).
- `serve` refuses to start with `--allow-import` and an empty allowlist, and refuses a non-loopback bind, as asserted by the serve tests.
- `make check` is green (clean cache) after every phase: gofumpt, vet (grammar-excluded), staticcheck, golangci (gocognit ≤ 7), govulncheck, 100.0% aggregate coverage.

## Phases

Each phase lands behind the green gate; a function is absent from the "done" set until its contract tests pass.

- **Phase 0 — seam.** The `#IMPORT!` error value, the `Fetcher` interface, `ComputeOptions.Fetcher` wiring (nil ⇒ `#IMPORT!`), the `HasImports` predicate, and the per-pass memo — no frontend network yet, exercised entirely through a mock `Fetcher`.
- **Phase 1 — scalar & linear.** `IMPORTCELL`, `IMPORTROW`, `IMPORTCOLUMN`: the handshake, values-only parse, strict shape, and `explain` detail over the mock.
- **Phase 2 — grids.** `IMPORTRANGE` (spilling array, `ResultCells`-bounded) and `IMPORTSHEET` (embeddable nested grid), reusing the array/embed machinery.
- **Phase 3 — real `Fetcher`.** The frontend `http.Client` implementation: `--allow-import`, the operator host allowlist, `https`-only + TLS verify, redirect re-check, timeout, size cap, and the cross-pass cache with explicit invalidation.
- **Phase 4 — frontend refresh & gating.** The TUI refresh key and `serve` refresh action (decoupled from the clock ticker), `serve`'s non-empty-allowlist and non-loopback refusals, and the `render` one-shot path.

## Compatibility note

This capability addresses the deferral recorded in [functions.md](functions.md) "Out of scope (ADR 0004)". It does **not** implement the Google/Excel network functions named there; it introduces a tsvsheet-native, content-typed, values-only import — a distinct, narrower surface with a two-sided opt-in handshake.
