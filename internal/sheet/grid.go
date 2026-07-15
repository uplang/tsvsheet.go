// Package sheet is the tsvsheet processor: it loads a .tsv value grid, applies a
// parsed .tsvt template (internal/tsvt) per SPECIFICATION §9 with the semantics
// fixed in specs/decisions/0003-open-semantics.md, and emits the computed grid.
package sheet

import (
	"bufio"
	"io"
	"strings"

	"github.com/uplang/tsvsheet.go/internal/constants"
)

// tab is the single field separator; newline terminates a row.
const (
	tab     = "\t"
	newline = "\n"
)

// Grid is a rectangular value grid indexed [row][col], 0-based. Cells are raw
// strings; the .tsv side carries no formulas (§2).
type Grid [][]string

// ReadTSV reads a tab-separated value grid. Rows are newline-separated; a
// trailing newline does not add an empty row. A read failure surfaces as
// constants.ErrReadInput.
func ReadTSV(r io.Reader) (Grid, error) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, bufio.MaxScanTokenSize), maxLineBytes)

	grid := Grid{}
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), tab))
	}
	if err := scanner.Err(); err != nil {
		return nil, constants.ErrReadInput.With(err)
	}
	return grid, nil
}

// maxLineBytes bounds a single scanned row (1 MiB) so a pathological input
// cannot exhaust memory silently.
const maxLineBytes = 1 << 20

// WriteTSV writes the grid as tab-separated rows, each terminated by a newline.
// A write failure surfaces as constants.ErrWriteFile. Callers wanting buffering
// pass a bufio.Writer; WriteTSV writes each row directly so a write error is
// reported at its source.
func WriteTSV(w io.Writer, g Grid) error {
	for _, row := range g {
		if _, err := io.WriteString(w, strings.Join(row, tab)+newline); err != nil {
			return constants.ErrWriteFile.With(err)
		}
	}
	return nil
}
