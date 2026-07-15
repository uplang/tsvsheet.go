package cli

import (
	"github.com/urfave/cli/v3"

	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// cellView is the JSON projection of one non-empty cell: its A1 address, source
// text, and whether it is a formula.
type cellView struct {
	Cell      string `json:"cell"`
	Source    string `json:"source"`
	IsFormula bool   `json:"formula"`
}

// sheetView is the JSON projection of a parsed spreadsheet.
type sheetView struct {
	Cells []cellView `json:"cells"`
}

// runParse parses a spreadsheet and writes its non-empty cells as JSON — a
// stable, jq-friendly surface for scripting.
func runParse(streams Streams, source sourcePath) error {
	reader, release, err := source.open(streams.In)
	if err != nil {
		return err
	}
	defer func() { _ = release() }()

	parsed, err := parseSheet(reader)
	if err != nil {
		return err
	}
	return writeJSON(streams.Out, sheetView{Cells: cellViews(parsed)})
}

// cellViews projects every non-empty cell to its JSON view.
func cellViews(s sheet.Sheet) []cellView {
	cells := s.Cells()
	views := make([]cellView, len(cells))
	for i, c := range cells {
		views[i] = cellView{
			Cell:      c.Address.String(),
			Source:    c.Text,
			IsFormula: c.IsFormula,
		}
	}
	return views
}

// parseCommand builds the `parse` command.
func parseCommand() *cli.Command {
	return &cli.Command{
		Name:      cmdParse,
		Usage:     "Parse a spreadsheet and emit its cells as JSON.",
		ArgsUsage: argSheetOptional,
		Description: `Parse a .tsvt spreadsheet (positional; omitted or "-" reads stdin) and write
its non-empty cells — address, source, and whether each is a formula — as JSON
to stdout.

Examples:
  tsvsheet parse sheet.tsvt | jq '.cells[] | select(.formula)'
  cat sheet.tsvt | tsvsheet parse`,
		Action: streamAction(func(s Streams, args positional) error {
			return runParse(s, args.at(0))
		}),
	}
}
