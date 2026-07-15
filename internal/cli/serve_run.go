package cli

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	httpserver "github.com/gomatic/go-httpserver"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/serve"
	"github.com/uplang/tsvsheet.go/internal/session"
)

// shutdownTimeout bounds graceful shutdown when serve is interrupted.
const shutdownTimeout = 5 * time.Second

// filePerm is the mode for a saved spreadsheet file.
const filePerm = 0o600

// runServe loads the spreadsheet into a session, hosts it over HTTP, and serves
// until ctx is cancelled (Ctrl-C). The sheet must be a file — serve saves edits
// back to it, so stdin is not a valid source.
func runServe(ctx context.Context, cfg serveConfig) error {
	server, err := loadServer(cfg.source)
	if err != nil {
		return err
	}
	http := httpserver.New(slog.Default(), httpserver.Host(cfg.host), httpserver.Port(cfg.port), server.Handler())
	slog.Info("serving spreadsheet", "url", "http://"+http.Addr())
	return http.Serve(ctx, shutdownTimeout)
}

// loadServer reads the spreadsheet file into a session and builds the HTTP
// server with a saver that writes edits back to that file.
func loadServer(source sourcePath) (serve.Server, error) {
	sess, persist, err := loadEditable(source)
	if err != nil {
		return serve.Server{}, err
	}
	return serve.NewServer(sess, persist), nil
}

// saver builds the persist function: it writes the session's current source
// back to the spreadsheet file. The plain func() error is assignable to both
// serve.Saver and tui.Saver.
func saver(sess *session.Session, source sourcePath) func() error {
	path := filepath.Clean(string(source))
	return func() error {
		if err := os.WriteFile(path, sess.Source(), filePerm); err != nil {
			return constants.ErrWriteFile.With(err, path)
		}
		return nil
	}
}
