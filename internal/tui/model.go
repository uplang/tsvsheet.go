// Package tui is the terminal frontend: a bubbletea model over the shared
// session.Session, giving the same capabilities as the browser editor —
// navigate the computed grid, edit data cells and the template, recompute, and
// save — driven by the one engine. The model holds no language semantics; every
// mutation goes through the session.
package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/uplang/tsvsheet.go/internal/session"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// Saver persists the worksheet; injected so the model stays filesystem-free.
type Saver func() error

// key names shared across the input handlers.
const (
	keyEnter = "enter"
	keyEsc   = "esc"
)

// cursorPos is a grid cursor coordinate; editText is an edit-buffer value.
type (
	cursorPos int
	editText  string
)

// mode is the model's current input mode.
type mode int

const (
	modeNav      mode = iota // navigating the grid
	modeCell                 // editing a data cell
	modeTemplate             // editing the template text
)

// Model is the terminal spreadsheet, a tea.Model over a session.
type Model struct {
	session          *session.Session
	save             Saver
	buffer           string
	status           string
	state            session.State
	row              int
	col              int
	mode             mode
	isConfirmingQuit bool
	isQuitting       bool
}

// New builds a model over a session and its saver, taking an initial snapshot.
func New(s *session.Session, save Saver) Model {
	return Model{session: s, save: save, state: s.Snapshot(), status: helpNav}
}

// helpNav is the resting-mode hint.
const helpNav = "arrows/hjkl move · enter edit · t template · ctrl+s save · q quit"

// Init implements tea.Model; the initial state is already loaded.
func (Model) Init() tea.Cmd { return nil }

// Update implements tea.Model, dispatching key messages by mode.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}
	switch m.mode {
	case modeCell:
		return m.keyCell(key)
	case modeTemplate:
		return m.keyTemplate(key)
	default:
		return m.keyNav(key)
	}
}

// keyNav handles navigation-mode keys: cursor movement and mode/command entry.
func (m Model) keyNav(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	if moved, handled := m.move(key.String()); handled {
		return moved, nil
	}
	return m.command(key.String())
}

// move applies a cursor movement, reporting whether the key was a movement.
func (m Model) move(key string) (Model, bool) {
	switch key {
	case "up", "k":
		m.row, m.isConfirmingQuit = clampDown(cursorPos(m.row)), false
	case "down", "j":
		m.row, m.isConfirmingQuit = clampUp(cursorPos(m.row), cursorPos(len(m.state.Computed)-1)), false
	case "left", "h":
		m.col, m.isConfirmingQuit = clampDown(cursorPos(m.col)), false
	case "right", "l":
		m.col, m.isConfirmingQuit = clampUp(cursorPos(m.col), cursorPos(m.width()-1)), false
	default:
		return m, false
	}
	return m, true
}

// command handles the non-movement navigation keys.
func (m Model) command(key string) (Model, tea.Cmd) {
	switch key {
	case keyEnter, "i":
		return m.beginCellEdit(), nil
	case "t":
		m.mode, m.buffer, m.status, m.isConfirmingQuit = modeTemplate, m.state.Template, helpTemplate, false
		return m, nil
	case "ctrl+s":
		return m.doSave(), nil
	case "q", "ctrl+c", keyEsc:
		return m.quit()
	default:
		m.isConfirmingQuit = false
		return m, nil
	}
}

// beginCellEdit enters cell-edit mode on an editable cell, or reports that the
// current cell is computed.
func (m Model) beginCellEdit() Model {
	if !m.editable(m.row, m.col) {
		m.status, m.isConfirmingQuit = "That cell is computed — edit a data cell (unshaded).", false
		return m
	}
	m.mode, m.buffer, m.status, m.isConfirmingQuit = modeCell, m.dataValue(m.row, m.col), helpCell, false
	return m
}

// quit exits, warning once when there are unsaved changes.
func (m Model) quit() (Model, tea.Cmd) {
	if m.state.IsDirty && !m.isConfirmingQuit {
		m.isConfirmingQuit, m.status = true, "Unsaved changes. Press q again to quit, or ctrl+s to save."
		return m, nil
	}
	m.isQuitting = true
	return m, tea.Quit
}

// helpCell and helpTemplate are the edit-mode hints.
const (
	helpCell     = "type value · enter commit · esc cancel"
	helpTemplate = "edit template · enter newline · ctrl+d apply · esc cancel"
)

// keyCell handles cell-edit keys.
func (m Model) keyCell(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case keyEnter:
		return m.commitCell(), nil
	case keyEsc:
		return m.toNav(), nil
	default:
		m.buffer = editBuffer(editText(m.buffer), key)
		return m, nil
	}
}

// commitCell writes the buffer to the current data cell and returns to nav.
// beginCellEdit only enters cell mode on an editable, non-negative data cell,
// and the cursor is clamped to the grid, so SetDataCell cannot fail here — the
// error is provably nil and deliberately not branched on.
func (m Model) commitCell() Model {
	_ = m.session.SetDataCell(sheet.Address{Row: m.row, Col: m.col}, m.buffer)
	return m.refreshedNav()
}

// keyTemplate handles template-edit keys.
func (m Model) keyTemplate(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "ctrl+d":
		return m.applyTemplate(), nil
	case keyEsc:
		return m.toNav(), nil
	case keyEnter:
		m.buffer += "\n"
		return m, nil
	default:
		m.buffer = editBuffer(editText(m.buffer), key)
		return m, nil
	}
}

// applyTemplate replaces the template, staying in template mode on error so the
// buffer is not lost.
func (m Model) applyTemplate() Model {
	if err := m.session.SetTemplate([]byte(m.buffer)); err != nil {
		m.status = err.Error()
		return m
	}
	return m.refreshedNav()
}

// doSave persists the worksheet, reporting the outcome.
func (m Model) doSave() Model {
	if err := m.save(); err != nil {
		m.status = err.Error()
		return m
	}
	m.session.MarkSaved()
	m.state, m.status, m.isConfirmingQuit = m.session.Snapshot(), "Saved.", false
	return m
}

// toNav returns to navigation mode without applying the buffer.
func (m Model) toNav() Model {
	m.mode, m.status, m.isConfirmingQuit = modeNav, helpNav, false
	return m
}

// refreshedNav re-snapshots the session and returns to navigation mode.
func (m Model) refreshedNav() Model {
	m.state, m.mode, m.status, m.isConfirmingQuit = m.session.Snapshot(), modeNav, helpNav, false
	return m
}

// editBuffer applies a printable key or backspace to an edit buffer.
func editBuffer(buffer editText, key tea.KeyMsg) string {
	if key.Type == tea.KeyBackspace {
		return trimLastRune(buffer)
	}
	if key.Type == tea.KeyRunes || key.Type == tea.KeySpace {
		return string(buffer) + string(key.Runes)
	}
	return string(buffer)
}

// trimLastRune drops the last rune of s.
func trimLastRune(s editText) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return string(s)
	}
	return string(runes[:len(runes)-1])
}

// clampDown decrements toward zero.
func clampDown(v cursorPos) int {
	if v <= 0 {
		return 0
	}
	return int(v - 1)
}

// clampUp increments toward the maximum.
func clampUp(v, limit cursorPos) int {
	if v >= limit {
		return int(limit)
	}
	return int(v + 1)
}
