// Package tui is the terminal frontend: a bubbletea model over the shared
// session.Session, editing the spreadsheet grid with the same capabilities as
// the browser editor — navigate cells, edit any cell (a value or an =formula),
// recompute, and save — driven by the one engine. The model holds no language
// semantics; every mutation goes through the session.
package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/uplang/tsvsheet.go/internal/session"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// Saver persists the spreadsheet; injected so the model stays filesystem-free.
type Saver func() error

// mode is the model's current input mode.
type mode int

const (
	modeNav  mode = iota // navigating the grid
	modeEdit             // editing the selected cell's source
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

// helpNav and helpEdit are the mode hints.
const (
	helpNav  = "arrows/hjkl move · enter edit · ctrl+s save · q quit"
	helpEdit = "type a value or =formula · enter commit · esc cancel"
)

// The key names the update loop dispatches on.
const (
	keyEnter = "enter"
	keyEsc   = "esc"
)

// editText is an in-progress cell edit buffer.
type editText string

// Init implements tea.Model; the initial state is already loaded.
func (Model) Init() tea.Cmd { return nil }

// Update implements tea.Model, dispatching key messages by mode.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}
	if m.mode == modeEdit {
		return m.keyEdit(key)
	}
	return m.keyNav(key)
}

// keyNav handles navigation-mode keys: cursor movement and commands.
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
		m.row, m.isConfirmingQuit = int(clampDown(cursorPos(m.row))), false
	case "down", "j":
		m.row, m.isConfirmingQuit = int(clampUp(cursorPos(m.row), cursorPos(m.height()-1))), false
	case "left", "h":
		m.col, m.isConfirmingQuit = int(clampDown(cursorPos(m.col))), false
	case "right", "l":
		m.col, m.isConfirmingQuit = int(clampUp(cursorPos(m.col), cursorPos(m.width()-1))), false
	default:
		return m, false
	}
	return m, true
}

// command handles the non-movement navigation keys.
func (m Model) command(key string) (Model, tea.Cmd) {
	switch key {
	case keyEnter, "i":
		m.mode, m.buffer, m.status, m.isConfirmingQuit = modeEdit, m.sourceAt(m.row, m.col), helpEdit, false
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

// quit exits, warning once when there are unsaved changes.
func (m Model) quit() (Model, tea.Cmd) {
	if m.state.IsDirty && !m.isConfirmingQuit {
		m.isConfirmingQuit, m.status = true, "Unsaved changes. Press q again to quit, or ctrl+s to save."
		return m, nil
	}
	m.isQuitting = true
	return m, tea.Quit
}

// keyEdit handles cell-edit keys.
func (m Model) keyEdit(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case keyEnter:
		return m.commit(), nil
	case keyEsc:
		return m.toNav(), nil
	default:
		m.buffer = string(editBuffer(editText(m.buffer), key))
		return m, nil
	}
}

// commit writes the buffer to the selected cell. A malformed formula keeps the
// model in edit mode (buffer preserved) and shows the error.
func (m Model) commit() Model {
	if err := m.session.SetCell(sheet.Address{Row: m.row, Col: m.col}, m.buffer); err != nil {
		m.status = err.Error()
		return m
	}
	return m.refreshedNav()
}

// doSave persists the spreadsheet, reporting the outcome.
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
func editBuffer(buffer editText, key tea.KeyMsg) editText {
	if key.Type == tea.KeyBackspace {
		return trimLastRune(buffer)
	}
	if key.Type == tea.KeyRunes || key.Type == tea.KeySpace {
		return buffer + editText(key.Runes)
	}
	return buffer
}

// trimLastRune drops the last rune of s.
func trimLastRune(s editText) editText {
	runes := []rune(s)
	if len(runes) == 0 {
		return s
	}
	return editText(runes[:len(runes)-1])
}

// clampDown decrements toward zero.
func clampDown(v cursorPos) cursorPos {
	if v <= 0 {
		return 0
	}
	return v - 1
}

// clampUp increments toward the maximum.
func clampUp(v, limit cursorPos) cursorPos {
	if v >= limit {
		return limit
	}
	return v + 1
}
