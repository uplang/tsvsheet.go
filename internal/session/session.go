// Package session is the stateful editing model shared by every interactive
// frontend (serve, tui): one worksheet — template text plus a data grid —
// mutated by edits and recomputed through the engine (internal/tsvt,
// internal/sheet). It is the repo's one sanctioned pointer-receiver type: it
// wraps mutable state guarded for concurrent use, so a single Session backs the
// HTTP handlers and the TUI model alike.
package session

import (
	"strings"
	"sync"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/sheet"
	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

// State is the complete read model a frontend renders: the computed grid, the
// template source, the raw data grid, static diagnostics, and the dirty flag.
// It is a value snapshot; mutating it never affects the Session.
type State struct {
	Computed    [][]string         `json:"computed"`
	Template    string             `json:"template"`
	Data        [][]string         `json:"data"`
	Diagnostics []sheet.Diagnostic `json:"diagnostics"`
	Dirty       bool               `json:"dirty"`
}

// Session is a mutable worksheet. Its methods are safe for concurrent use.
type Session struct {
	mu          sync.Mutex
	templateSrc []byte
	template    tsvt.Template
	data        sheet.Grid
	computed    sheet.Grid
	diagnostics []sheet.Diagnostic
	dirty       bool
}

// New builds a session from template source and a data grid, parsing and
// computing eagerly. It fails on a syntax error or a template the processor
// rejects; the resulting session is clean (not dirty).
func New(templateSrc []byte, data sheet.Grid) (*Session, error) {
	s := &Session{data: cloneGrid(data)}
	if err := s.loadTemplate(templateSrc); err != nil {
		return nil, err
	}
	return s, nil
}

// loadTemplate parses and computes src against the current data, committing the
// new template state only on success. Because every field is assigned last,
// a failure leaves the session's prior state fully intact (atomic replace).
func (s *Session) loadTemplate(src []byte) error {
	tmpl, err := tsvt.Parse(tsvt.Source(src))
	if err != nil {
		return err
	}
	computed, err := sheet.Compute(tmpl, s.data)
	if err != nil {
		return err
	}
	s.templateSrc = append([]byte(nil), src...)
	s.template = tmpl
	s.computed = computed
	s.diagnostics = sheet.Check(tmpl)
	return nil
}

// SetTemplate replaces the template text, reparsing and recomputing. On a
// syntax error or rejected template the previous state is retained unchanged
// and the error is returned; on success the session is marked dirty.
func (s *Session) SetTemplate(src []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.loadTemplate(src); err != nil {
		return err
	}
	s.dirty = true
	return nil
}

// SetDataCell edits one raw data cell (growing the grid to reach an append
// position), then recomputes. A negative address is rejected; otherwise the
// template is unchanged and already valid, so the recompute does not newly fail.
func (s *Session) SetDataCell(a sheet.Address, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if a.Row < 0 || a.Col < 0 {
		return constants.ErrInvalidValue.With(nil, "cell", a.String())
	}
	s.data = setCell(s.data, a.Row, a.Col, value)
	s.dirty = true
	return s.loadTemplate(s.templateSrc)
}

// Snapshot returns a deep-copied read model safe for the caller to hold and
// mutate.
func (s *Session) Snapshot() State {
	s.mu.Lock()
	defer s.mu.Unlock()
	return State{
		Computed:    grid(s.computed),
		Template:    string(s.templateSrc),
		Data:        grid(s.data),
		Diagnostics: append([]sheet.Diagnostic(nil), s.diagnostics...),
		Dirty:       s.dirty,
	}
}

// Explain traces how the computed cell at addr was produced, over the current
// template and data (see sheet.Explain).
func (s *Session) Explain(addr sheet.Address) (sheet.Trace, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return sheet.Explain(s.template, s.data, addr)
}

// MarkSaved clears the dirty flag after a frontend persists the worksheet.
func (s *Session) MarkSaved() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dirty = false
}

// TemplateText returns a copy of the template source for saving.
func (s *Session) TemplateText() []byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]byte(nil), s.templateSrc...)
}

// DataTSV returns the raw data grid encoded as TSV for saving: tab-joined
// rows, each newline-terminated.
func (s *Session) DataTSV() []byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	var b strings.Builder
	for _, row := range s.data {
		b.WriteString(strings.Join(row, "\t"))
		b.WriteByte('\n')
	}
	return []byte(b.String())
}

// setCell writes value at (row, col), growing the grid with empty rows and
// cells so an append position (one past the end) creates new space.
func setCell(g sheet.Grid, row, col int, value string) sheet.Grid {
	out := cloneGrid(g)
	for len(out) <= row {
		out = append(out, []string{})
	}
	line := out[row]
	for len(line) <= col {
		line = append(line, "")
	}
	line[col] = value
	out[row] = line
	return out
}

// cloneGrid deep-copies a grid.
func cloneGrid(g sheet.Grid) sheet.Grid {
	out := make(sheet.Grid, len(g))
	for i, row := range g {
		out[i] = append([]string(nil), row...)
	}
	return out
}

// grid deep-copies a grid to a plain [][]string for a State snapshot.
func grid(g sheet.Grid) [][]string {
	out := make([][]string, len(g))
	for i, row := range g {
		out[i] = append([]string(nil), row...)
	}
	return out
}
