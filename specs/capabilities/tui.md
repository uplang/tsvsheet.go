# tui — terminal spreadsheet over the same session

## Goal

`tsvsheet tui` opens the worksheet in a bubbletea terminal UI with the same capabilities the web UI has, proving the session seam is frontend-agnostic.

## Requirements

- R1: Grid view of the computed sheet (viewport-scrollable), cursor navigation (arrows/hjkl), column letters + row numbers, error-valued cells visibly marked.
- R2: Edit mode on a data cell writes through `Session.SetDataCell` and repaints from `Snapshot()`; a template editor pane (toggle) edits the whole `.tsvt` text through `SetTemplate`, showing syntax errors without losing the buffer.
- R3: Save (`ctrl+s`) writes both files via the same injected writers the CLI uses; quit (`q`/`ctrl+c`) warns on unsaved changes.
- R4: The model is a pure `tea.Model` over `Session`; no engine logic in the TUI layer.

## Acceptance Criteria

- Model Update/View unit tests: navigation bounds, edit commit/cancel, template error display, save invocation, quit-with-dirty warning — 100% coverage of the model package (bubbletea models are plain functions and fully testable).
- Manual smoke on the worked example.
