package list

import (
	app "github.com/gomatic/go-app"
	"github.com/urfave/cli/v3"

	domain "github.com/uplang/tsvsheet.go/internal/domain/config/list"
)

const (
	name        = `list`
	usage       = `List all configuration values`
	description = `List all configuration keys and values.

Examples:
  tsvsheet config list
  tsvsheet config list --prefix app.
  tsvsheet config list --prefix database.

Output is a JSON object of key/value pairs, matching the structured output
used across every command and easy to parse with standard tools (e.g. jq).

This command demonstrates:
  - No positional arguments
  - Optional filter flags
  - Structured JSON output`
)

const (
	prefixFlag = "prefix"
)

var (
	cfg       domain.Config
	runAction = domain.Run
)

// Command returns the CLI command definition.
func Command() *cli.Command {
	return &cli.Command{
		Name:        name,
		Usage:       usage,
		Description: description,
		Action:      app.Default(&cfg, runAction),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        prefixFlag,
				Aliases:     []string{"p"},
				Sources:     cli.EnvVars("CONFIG_PREFIX"),
				Usage:       "Only list keys with this prefix",
				Destination: (*string)(&cfg.Prefix),
			},
		},
	}
}
