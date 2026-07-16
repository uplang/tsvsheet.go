# ADR 0006 — content-typed import: pulling external data into the grid over HTTP

## Status

Proposed (2026-07-16). Addresses the network/import capability that [ADR 0004](0004-excel-sheets-functions.md) deferred as "later or never" (its Goal lists `IMPORT*`/`WEBSERVICE` explicitly out of scope). This ADR does **not** revive the Google/Excel network functions named there (`IMPORTHTML`, `IMPORTXML`, `GOOGLEFINANCE`, `WEBSERVICE`, …), which parse arbitrary web content and remain out of scope; it introduces a **tsvsheet-native** import built on a content-type handshake. Governs a new capability spec, [import](../capabilities/import.md).

## Context

We want a `.tsvt` to pull real-world data and the results of remote computation into the grid: the weather for an area, the current location or public IP, a value fetched from a URL, or the output of a program running an algorithm the formula language cannot express. The unifying need is "reach information anywhere and bring it into a cell."

The engine is deliberately pure and deterministic — `Compute` is memoized, every type is a value receiver, and [internal/sheet/limits.go](../../internal/sheet/limits.go) states in plain terms that every `.tsvt` is **untrusted input**. Any capability that reaches outside the process must not compromise that posture. The entire security history of spreadsheets is one lesson repeated (DDE `=cmd|…`, CSV injection, VBA macros): the danger is not "reaching outside" but **ambient authority** — a document that *acts* when opened, mailed to someone who trusts documents.

Two designs were weighed. A local **shell function** (`SHELL(argv…)`) to run any allowlisted command was considered first and **rejected** (§9): it reintroduces code-execution-on-open, and its most compelling use case (running a program for a complex algorithm) is served just as well by a remote endpoint that speaks our content type. The chosen design is a **content-typed HTTP import**: the value crossing the wire is data, not code, and the source is one that was *deliberately built* to serve tsvsheet.

## Decisions

### 1. A two-sided content-type handshake

An import succeeds only when both ends opt in. The **operator** opts in by enabling the feature and allowlisting the host (§7); the **endpoint author** opts in by serving one of this project's media types. The formula's granularity determines the media type it sends as the request `Accept` header; the response's `Content-Type` **must equal** that type, or the import fails (§5). A response that carries HTML, JSON, or any type not ours is never ingested — so an import can only ever draw from a source someone intentionally built to speak tsvsheet. This is both the security property (no accidental ingestion of "whatever was at that URL") and the design property (every importable source is intentional).

### 2. Five functions, one per granularity

The function name *is* the content type; each maps onto an existing engine return shape (`cellset`/`scalar`/`[][]Value`, [internal/sheet/resolve.go](../../internal/sheet/resolve.go)).

| Function | `Accept` media type | Returns |
| --- | --- | --- |
| `IMPORTCELL(url)` | `application/vnd.tsvsheet.cell+tsv` | a scalar value |
| `IMPORTROW(url)` | `application/vnd.tsvsheet.row+tsv` | a 1×N cellset |
| `IMPORTCOLUMN(url)` | `application/vnd.tsvsheet.column+tsv` | an N×1 cellset |
| `IMPORTRANGE(url)` | `application/vnd.tsvsheet.range+tsv` | an R×C array that **spills** (ADR 0004 §4) |
| `IMPORTSHEET(url)` | `application/vnd.tsvsheet+tsv` | one embeddable value rendered as a nested grid |

The media types use the RFC 6838 vendor tree with a hierarchical subtype for granularity and the `+tsv` structured-syntax suffix. `IMPORTRANGE` and `IMPORTSHEET` carry the same payload (a grid of values); they differ only in placement — `IMPORTRANGE` spills a block into the grid, `IMPORTSHEET` renders the grid nested inside its one cell, the HTTP analogue of `SHEET()` ([internal/sheet/embed.go](../../internal/sheet/embed.go)) but values-only. The five names are ordinary function calls the existing grammar already admits — **no grammar change**.

### 3. Values only — the wire body is a `.tsvt` fragment

Every import ingests **already-computed values, never formulas**. The response body is exactly a `.tsvt` fragment at the requested granularity (one cell, one TSV line, one value-per-line column, or a TSV grid), parsed by the engine's existing TSV reader — so "the format on the wire" equals "the format in a file." A leading `=` in an imported value enters the grid as the **literal text** `=…`, never re-parsed or evaluated. Three consequences follow directly:

- **No remote code execution.** We never evaluate a remote formula locally. This is the deliberate line against local `SHEET(file)`, which *does* execute: a local file is same-user, same-trust; a URL is not. Local embed runs trusted local code; remote import ingests untrusted remote data.
- **No fetch recursion, no cycle detection.** A value cannot trigger a fetch, so an imported grid cannot fan out into further fetches. The `visiting`/`#CIRC!` apparatus that foreign sheets require does not apply.
- **Shape is strict.** A `cell` response must be exactly one cell, a `row` exactly one line, a `column` one value per line, a `range`/`sheet` rectangular. A shape mismatch is `#IMPORT!` (§5), not a best-effort salvage.

### 4. The injected `Fetcher` seam

The engine stays filesystem- and network-free. A `Fetcher` collaborator is injected through the compute pass exactly as `Loader` is today ([internal/sheet/embed.go](../../internal/sheet/embed.go)): given a URL and the requested media type, it returns the response body and its `Content-Type`, or a frontend error. A **`nil` Fetcher** (the default) means every import resolves to `#IMPORT!` — the feature is simply absent, mirroring how a `nil` `Loader` makes `SHEET()` resolve to `#REF!`. The engine never imports `net/http`; the frontend constructs the `Fetcher` with the allowlist, `http.Client`, timeout, and size cap. The `Fetcher` also owns the **cross-pass cache** (§6): the engine calls it on every pass, but it serves cached results until the frontend invalidates it on an explicit refresh — so ordinary and clock-ticker recomputes trigger no network I/O. Tests inject a mock `Fetcher`, keeping the 100% aggregate-coverage gate reachable with no real network.

### 5. Error model — one cell error value, structured reasons via `explain`

A new spreadsheet error value **`#IMPORT!`** is added to [internal/sheet/value.go](../../internal/sheet/value.go) alongside the existing set (following the non-Excel `#CIRC!` precedent; it is also added to the `parseError` allow-set). Every import failure — feature disabled, host not allowed, non-2xx status, content-type mismatch, shape mismatch, oversize body, timeout — surfaces as `#IMPORT!` in the cell. The *reason* is carried by `explain` (the existing trace surface, [internal/sheet/explain.go](../../internal/sheet/explain.go)): `explain B2` reports the URL, the HTTP status, the content type received vs. expected, the byte count, and the duration. The frontend `Fetcher`'s failure reasons are `errs.Const` sentinels in [internal/constants](../../internal/constants) (joining the existing `ErrForbidden` CSRF sentinel), so each failure path is matchable with `errors.Is` and testable.

### 6. Refresh model — a cache across passes, never the clock ticker

Imports are **not** clock-volatile and **must never ride the isnow refresh ticker**. Clock volatility (`TODAY`/`NOW`/`ISNOW`) is cheap — reading an injected clock — so it is refreshed by a sub-second/second-scale ticker to keep the displayed time current. Import volatility is the opposite: each evaluation is a network request with external side effects. Binding it to that ticker would refetch every import URL on every tick, hammering endpoints, inviting rate-limiting, and stalling the recompute loop. So:

- **A frontend-owned import cache**, keyed by `(url, media type)`, holds fetched values **across compute passes**. A normal recompute — whether driven by a cell edit or by the clock ticker that keeps `NOW()` current — reads imported values from this cache and performs **no network I/O**. The engine calls the injected `Fetcher` on every pass (memoized once per pass, §4/below); the `Fetcher` serves from cache until the frontend invalidates it. Only an explicit refresh re-fetches.
- **Imports are absent from `IsVolatile`** ([internal/sheet/volatile.go](../../internal/sheet/volatile.go)) — that predicate stays the isnow-ticker signal for `TODAY`/`NOW`/`ISNOW` alone. A **separate `HasImports` predicate** reports that a sheet contains import calls, so a frontend can expose a refresh control without ever placing imports on the clock ticker.
- **Refresh is explicit and per-frontend, decoupled from the clock:** `render` is one-shot (each invocation fetches fresh; the cache only dedups within the single run); the TUI binds a distinct "refresh imports" key that clears the cache and recomputes; `serve` exposes an explicit refresh action (endpoint/button), separate from its `--refresh-interval` clock auto-refresh. An optional, **separate** import-refresh interval may be offered later — independent of the clock ticker, typically minutes-scale, and off by default.
- **Within a single pass**, `(url, media type)` is fetched at most once, so per-pass determinism holds regardless of cache state.

### 7. The trust gate

- **Off by default.** A frontend `--allow-import` flag turns the capability on; unset, the `Fetcher` is `nil` and imports are `#IMPORT!`.
- **The allowlist lives only in the operator's environment** — a host allowlist supplied by the operator (repeated flags and/or a policy file), never read from the `.tsvt`. No cell, header, or directive inside a sheet can enable the feature or widen the allowlist. The file is data; the authority is the human who ran the tool. An empty allowlist with the flag on denies every host.
- **`https` only, TLS verified.** No `http`, no certificate-skip. Redirects are followed within a small cap, and **each hop is re-checked against the allowlist** so a redirect cannot escape it.
- **Bounded resources.** A per-import timeout (context deadline) and a response-size cap (reusing the `limits.go` discipline; a grid import is additionally bounded by `ResultCells`) mean a hung or spewing endpoint cannot wedge or OOM a compute pass.

### 8. Per-frontend gating

`render`/`explain`/`tui` honor `--allow-import` and the allowlist as above. **serve** takes the stricter posture, because an HTTP server that fetches outbound URLs is an SSRF surface: it requires a **non-empty allowlist** (no "allow-import with no hosts" default-open), keeps the existing hard non-loopback refusal, and its outbound reach is confined to allowlisted hosts and values-only responses — bounding SSRF to hosts the operator explicitly trusts. serve reuses the same `--allow-import`/allowlist rather than a second flag, because (unlike shell) the capability is outbound-data-only and the allowlist already scopes it.

### 9. Considered and rejected — local shell execution

A `SHELL(argv…)` function running any operator-allowlisted command (argv-as-list, no `sh -c`, stdout→value, stderr+exit→`explain`, off by default) was considered as the original framing of "the executable aspect of a tsvsheet." It is **rejected**:

- It reintroduces **code execution on open**. Even fully allowlisted, a command running as a side effect of computing an untrusted `.tsvt` is the exact ambient-authority hazard that has produced every historical spreadsheet RCE. The content-typed import moves the computation to a remote endpoint the operator explicitly trusts, and only *data* crosses into the grid.
- It is **not needed for the motivating use cases.** "Run a program for a complex algorithm" is served by a program that emits `application/vnd.tsvsheet.cell+tsv`; the four data cases (weather, IP, location, curl-a-value) are the same import shape. The import design delivers the executable aspect without a local subprocess.
- serve made it untenable: an HTTP endpoint that executes local commands is a remote shell, a categorically larger threat than outbound-data-only import.

Should local execution ever be revisited, it must extend this ADR with a new decision and its own capability spec — it is closed, not merely unbuilt.

## Consequences

- The engine gains one injected collaborator (`Fetcher`), one error value (`#IMPORT!`), and five value-typed, 100%-covered builtins, with no grammar change — the import functions ride entirely on existing syntax.
- `SHEET`/`OUTPUT`/`INPUT` (ADR 0005) were the first builtins with side inputs; the import functions are the first with an *external, non-deterministic* side input, threaded through the pass as an injected collaborator (never a global) and sampled once per pass, preserving determinism and testability.
- The `.tsvt` format gains network semantics with a values-only, handshake-gated boundary; `SPECIFICATION.md` (grammar repo) gains an "Imported values" section, with anything underspecified marked `[open]` rather than invented (ADR 0003 discipline).
- The security posture is preserved by construction: off by default, file-cannot-self-enable, values-only, two-sided opt-in, https-only, allowlisted, bounded — and the most dangerous alternative (local shell) is recorded as rejected so the decision is not silently relitigated.
