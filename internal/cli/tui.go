package cli

import (
	"github.com/urfave/cli/v3"
)

// tuiConfig binds the tui command's spreadsheet path.
type tuiConfig struct {
	source sourcePath
}

// tuiCommand builds the `tui` command.
func tuiCommand() *cli.Command {
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
		Action: streamAction(func(s Streams, args positional) error {
			cfg.source = args.at(0)
			return runTUI(s, cfg)
		}),
	}
}
