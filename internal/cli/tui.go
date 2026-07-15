package cli

import (
	"github.com/urfave/cli/v3"
)

// tuiConfig binds the tui command's spreadsheet path and path-access mode.
type tuiConfig struct {
	source       sourcePath
	isUnconfined pathAccess
}

// tuiCommand builds the `tui` command.
func tuiCommand() *cli.Command {
	isUnconfined := false
	cfg := tuiConfig{}
	return &cli.Command{
		Name:      cmdTUI,
		Usage:     "Edit a spreadsheet in a terminal UI.",
		ArgsUsage: "<sheet>",
		Description: `Open the spreadsheet in a terminal grid: navigate cells, edit any cell (a
value or an =formula), recompute, and save — the same capabilities as the
browser editor, driven by the same engine. The sheet is a required positional
file path.

Examples:
  tsvsheet tui sheet.tsvt`,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: flagAllowAnyPaths, Usage: usageAllowAnyPaths, Destination: &isUnconfined},
		},
		Action: streamAction(func(s Streams, args positional) error {
			cfg.source = args.at(0)
			cfg.isUnconfined = pathAccess(isUnconfined)
			return runTUI(s, cfg)
		}),
	}
}
