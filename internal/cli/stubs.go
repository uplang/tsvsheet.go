package cli

import "github.com/uplang/tsvsheet.go/internal/constants"

// runTUI is implemented in tui_run.go (task 8); placeholder until then.
func runTUI(_ Streams, _ tuiConfig) error {
	return constants.ErrUnsupported.With(nil, "message", "tui not yet implemented")
}
