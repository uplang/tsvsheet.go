package get

import "github.com/uplang/tsvsheet.go/internal/config"

// Config holds the flags for the "config get" command.
type Config struct {
	DefaultValue config.Value
}
