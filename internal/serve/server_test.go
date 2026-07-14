package serve_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/serve"
	"github.com/uplang/tsvsheet.go/internal/session"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

var (
	sampleData     = sheet.Grid{{"1", "2", "3", "4"}, {"2", "3", "4", "5"}, {"3", "4", "5", "6"}}
	sampleTemplate = []byte("=body\nE=C + D\n")
)

// testServer builds a server over a fresh session and records whether save ran.
func testServer(t *testing.T) (serve.Server, *bool) {
	t.Helper()
	sess, err := session.New(sampleTemplate, sampleData)
	require.NoError(t, err)
	saved := false
	return serve.NewServer(sess, func() error { saved = true; return nil }), &saved
}

// do issues a request against a server's handler and returns the recorder.
func do(t *testing.T, srv serve.Server, method, target, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	rec := httptest.NewRecorder()
	srv.Handler().ServeHTTP(rec, req)
	return rec
}

func TestState(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodGet, "/api/state", "")
	require.Equal(t, http.StatusOK, rec.Code)

	var state session.State
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &state))
	assert.Equal(t, "7", state.Computed[0][4])
	assert.False(t, state.IsDirty)
}

func TestSetTemplate_OK(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodPut, "/api/template", `{"text":"=body\nE=C * D\n"}`)
	require.Equal(t, http.StatusOK, rec.Code)

	var state session.State
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &state))
	assert.Equal(t, "12", state.Computed[0][4])
	assert.True(t, state.IsDirty)
}

func TestSetTemplate_SyntaxError(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodPut, "/api/template", `{"text":"=sum("}`)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")
}

func TestSetTemplate_BadBody(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodPut, "/api/template", `not json`)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSetDataCell_OK(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodPut, "/api/data/cell", `{"row":0,"col":2,"value":"10"}`)
	require.Equal(t, http.StatusOK, rec.Code)

	var state session.State
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &state))
	assert.Equal(t, "14", state.Computed[0][4]) // C+D = 10+4
}

func TestSetDataCell_InvalidAddress(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodPut, "/api/data/cell", `{"row":-1,"col":0,"value":"x"}`)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
}

func TestSetDataCell_BadBody(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodPut, "/api/data/cell", `{`)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSave_OK(t *testing.T) {
	t.Parallel()

	srv, saved := testServer(t)
	rec := do(t, srv, http.MethodPost, "/api/save", "")
	require.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, *saved)
	assert.Contains(t, rec.Body.String(), `"saved":true`)
}

func TestSave_Error(t *testing.T) {
	t.Parallel()

	sess, err := session.New(sampleTemplate, sampleData)
	require.NoError(t, err)
	srv := serve.NewServer(sess, func() error { return errors.New("disk full") })

	rec := do(t, srv, http.MethodPost, "/api/save", "")
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "disk full")
}

func TestExplain_OK(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodGet, "/api/explain?cell=E1", "")
	require.Equal(t, http.StatusOK, rec.Code)

	var trace sheet.Trace
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &trace))
	assert.Equal(t, "7", trace.Value)
	assert.Equal(t, "C + D", trace.Formula)
}

func TestExplain_BadCell(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodGet, "/api/explain?cell=bogus", "")
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestExplain_OutOfGrid(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodGet, "/api/explain?cell=Z99", "")
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUI_ServesHTML(t *testing.T) {
	t.Parallel()

	srv, _ := testServer(t)
	rec := do(t, srv, http.MethodGet, "/", "")
	require.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "<!doctype html>")
	assert.Contains(t, rec.Header().Get("Content-Type"), "text/html")
}
