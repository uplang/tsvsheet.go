// Package serve exposes a session.Session over an HTTP JSON API and hosts the
// embedded browser spreadsheet that consumes it (contracts/engine-api.md). It
// is a thin projection: every endpoint round-trips through the one shared
// engine, so the web UI holds no language semantics.
package serve

import (
	"encoding/json"
	"net/http"

	"github.com/uplang/tsvsheet.go/internal/session"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// Saver persists the worksheet's current template and data. It is injected so
// the server stays filesystem-free and testable.
type Saver func() error

// Server serves the JSON API and embedded UI for one worksheet session.
type Server struct {
	session *session.Session
	save    Saver
}

// NewServer builds a server over a session and a save function.
func NewServer(s *session.Session, save Saver) *Server {
	return &Server{session: s, save: save}
}

// Handler returns the HTTP handler: the JSON API under /api/ and the embedded
// single-page UI at the root.
func (srv *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/state", srv.handleState)
	mux.HandleFunc("PUT /api/template", srv.handleTemplate)
	mux.HandleFunc("PUT /api/data/cell", srv.handleDataCell)
	mux.HandleFunc("POST /api/save", srv.handleSave)
	mux.HandleFunc("GET /api/explain", srv.handleExplain)
	mux.Handle("GET /", uiHandler())
	return mux
}

// handleState returns the current worksheet read model.
func (srv *Server) handleState(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, srv.session.Snapshot())
}

// templateRequest is the PUT /api/template body.
type templateRequest struct {
	Text string `json:"text"`
}

// handleTemplate replaces the template text, returning the new state or 422
// with the syntax/validation error.
func (srv *Server) handleTemplate(w http.ResponseWriter, r *http.Request) {
	var req templateRequest
	if !decode(w, r, &req) {
		return
	}
	if err := srv.session.SetTemplate([]byte(req.Text)); err != nil {
		writeError(w, http.StatusUnprocessableEntity, err)
		return
	}
	writeJSON(w, http.StatusOK, srv.session.Snapshot())
}

// cellRequest is the PUT /api/data/cell body; row and col are 0-based raw-grid
// indices (contracts/engine-api.md).
type cellRequest struct {
	Row   int    `json:"row"`
	Col   int    `json:"col"`
	Value string `json:"value"`
}

// handleDataCell edits one raw data cell and returns the new state.
func (srv *Server) handleDataCell(w http.ResponseWriter, r *http.Request) {
	var req cellRequest
	if !decode(w, r, &req) {
		return
	}
	if err := srv.session.SetDataCell(sheet.Address{Row: req.Row, Col: req.Col}, req.Value); err != nil {
		writeError(w, http.StatusUnprocessableEntity, err)
		return
	}
	writeJSON(w, http.StatusOK, srv.session.Snapshot())
}

// saveResponse is the POST /api/save body.
type saveResponse struct {
	Saved bool `json:"saved"`
}

// handleSave persists the worksheet and clears the dirty flag.
func (srv *Server) handleSave(w http.ResponseWriter, _ *http.Request) {
	if err := srv.save(); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	srv.session.MarkSaved()
	writeJSON(w, http.StatusOK, saveResponse{Saved: true})
}

// handleExplain traces one computed cell named by the `cell` query parameter.
func (srv *Server) handleExplain(w http.ResponseWriter, r *http.Request) {
	at, err := sheet.ParseAddress(r.URL.Query().Get("cell"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	trace, err := srv.session.Explain(at)
	if err != nil {
		writeError(w, http.StatusNotFound, err)
		return
	}
	writeJSON(w, http.StatusOK, trace)
}

// decode reads a JSON request body into v, writing a 400 and returning false on
// a malformed body.
func decode(w http.ResponseWriter, r *http.Request, v any) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return false
	}
	return true
}

// errorResponse is the JSON error envelope.
type errorResponse struct {
	Error string `json:"error"`
}

// writeError writes a JSON error envelope with the given status.
func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, errorResponse{Error: err.Error()})
}

// writeJSON encodes v as JSON with the given status.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
