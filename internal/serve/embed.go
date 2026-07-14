package serve

import (
	_ "embed"
	"net/http"
)

// indexHTML is the embedded single-page UI (all CSS and JS inline, no external
// assets), served for the site root.
//
//go:embed web/index.html
var indexHTML []byte

// uiHandler serves the embedded UI.
func uiHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write(indexHTML)
	})
}
