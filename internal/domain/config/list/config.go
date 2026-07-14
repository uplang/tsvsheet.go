package list

import "github.com/uplang/tsvsheet.go/internal/config"

// Config holds the flags for the "config list" command.
type Config struct {
	Prefix config.Prefix
}
