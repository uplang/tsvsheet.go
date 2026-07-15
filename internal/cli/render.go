package cli

import (
	"path/filepath"
	"time"

	"github.com/urfave/cli/v3"

	"github.com/uplang/tsvsheet.go/internal/loader"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// runRender parses the spreadsheet, computes it (resolving SHEET(...) references
// when the source is a file), and writes the resulting value grid as TSV.
// Errors go to the caller (and thence stderr); stdout carries only the computed
// grid, so render composes in unix pipelines.
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
	return sheet.WriteTSV(streams.Out, parsed.ComputeWith(computeOptions(source)))
}

// computeOptions builds the compute options for a source: a filesystem sheet
// loader rooted at the file's directory (so SHEET embeds sibling sheets), or no
// loader for stdin (SHEET resolves to #REF!).
func computeOptions(source sourcePath) sheet.ComputeOptions {
	if source.isStdin() {
		return sheet.ComputeOptions{At: time.Now()}
	}
	path := filepath.Clean(string(source))
	return sheet.ComputeOptions{
		At:     time.Now(),
		Loader: loader.FS(loader.Dir(filepath.Dir(path))),
		Base:   sheet.Path(filepath.Base(path)),
	}
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
