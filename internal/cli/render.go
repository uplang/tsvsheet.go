package cli

import (
	"github.com/urfave/cli/v3"

	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// runRender parses the spreadsheet, computes it, and writes the resulting value
// grid as TSV. Errors go to the caller (and thence stderr); stdout carries only
// the computed grid, so render composes in unix pipelines.
func runRender(streams Streams, source sourcePath) error {
	reader, release, err := source.open(streams.In)
	if err != nil {
		return err
	}
	defer func() { _ = release() }()

	parsed, err := parseSheet(reader)
	if err != nil {
		return err
	}
	return sheet.WriteTSV(streams.Out, parsed.Compute())
}

// renderCommand builds the `render` command.
func renderCommand() *cli.Command {
	return &cli.Command{
		Name:      cmdRender,
		Usage:     "Compute a spreadsheet and write the values as TSV.",
		ArgsUsage: argSheetOptional,
		Description: `Compute a .tsvt spreadsheet — a grid of literal and =formula cells — and
write the computed value grid as TSV to stdout. The sheet is positional;
omitted or "-" reads stdin.

Examples:
  tsvsheet render sheet.tsvt
  cat sheet.tsvt | tsvsheet render`,
		Action: streamAction(func(s Streams, args positional) error {
			return runRender(s, args.at(0))
		}),
	}
}
