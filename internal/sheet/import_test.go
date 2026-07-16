package sheet_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// The content-typed import media wire strings (ADR 0006 §2). Hardcoded here as
// the black-box contract each IMPORT* function must send as its Accept header.
const (
	mediaSheetWire  sheet.MediaType = "application/vnd.tsvsheet+tsv"
	mediaCellWire   sheet.MediaType = "application/vnd.tsvsheet.cell+tsv"
	mediaRowWire    sheet.MediaType = "application/vnd.tsvsheet.row+tsv"
	mediaColumnWire sheet.MediaType = "application/vnd.tsvsheet.column+tsv"
	mediaRangeWire  sheet.MediaType = "application/vnd.tsvsheet.range+tsv"
)

// echoFetcher answers every fetch with body and a Content-Type that echoes the
// requested Accept, so the handshake always matches; it captures the Accept it
// was sent so a test can assert each function requests the correct media type.
type echoFetcher struct {
	accept *sheet.MediaType
	body   []byte
}

func (f echoFetcher) Fetch(_ sheet.ImportURL, accept sheet.MediaType) (sheet.FetchResult, error) {
	if f.accept != nil {
		*f.accept = accept
	}
	return sheet.FetchResult{Body: f.body, ContentType: accept}, nil
}

// fixedFetcher answers with a fixed result and error, for the failure paths
// (handshake mismatch, transport error).
type fixedFetcher struct {
	err    error
	result sheet.FetchResult
}

func (f fixedFetcher) Fetch(_ sheet.ImportURL, _ sheet.MediaType) (sheet.FetchResult, error) {
	return f.result, f.err
}

// importGrid parses src and computes it with the injected Fetcher and Limits.
func importGrid(t *testing.T, src string, f sheet.Fetcher, limits sheet.Limits) sheet.Grid {
	t.Helper()
	s, err := sheet.Parse([]byte(src))
	require.NoError(t, err)
	return s.ComputeWith(sheet.ComputeOptions{At: time.Now(), Fetcher: f, Limits: limits})
}

func TestHasImports(t *testing.T) {
	t.Parallel()

	assert.True(t, parse(t, "=importcell(\"https://x.example/v\")\n").HasImports())
	assert.True(t, parse(t, "=importsheet(\"https://x.example/v\")\n").HasImports())
	assert.False(t, parse(t, "=sum(A1:A2)\n").HasImports()) // a call, but not an import
	assert.False(t, parse(t, "plain\n").HasImports())       // no formula at all
}

func TestImportDisabledYieldsImportError(t *testing.T) {
	t.Parallel()

	for _, src := range []string{
		"=importcell(\"https://x.example/v\")\n",
		"=importrow(\"https://x.example/v\")\n",
		"=importcolumn(\"https://x.example/v\")\n",
		"=importrange(\"https://x.example/v\")\n",
		"=importsheet(\"https://x.example/v\")\n",
	} {
		assert.Equal(t, "#IMPORT!", cellAt(t, compute(t, src), 0, 0), "no Fetcher injected: %s must be #IMPORT!", src)
	}
}

func TestImportErrorLiteralPropagates(t *testing.T) {
	t.Parallel()

	// A cell literally holding #IMPORT! round-trips as an error value and
	// propagates through a reference (isErrorCode recognizes it).
	assert.Equal(t, "#IMPORT!", cellAt(t, compute(t, "#IMPORT!\t=A1\n"), 0, 1))
}

func TestImportCellScalar(t *testing.T) {
	t.Parallel()

	var accept sheet.MediaType
	grid := importGrid(
		t,
		"=importcell(\"u\")\n",
		echoFetcher{body: []byte("42\n"), accept: &accept},
		sheet.DefaultLimits(),
	)
	assert.Equal(t, "42", cellAt(t, grid, 0, 0))
	assert.Equal(t, mediaCellWire, accept, "IMPORTCELL must request the cell media type")
}

func TestImportRowSpillsHorizontally(t *testing.T) {
	t.Parallel()

	var accept sheet.MediaType
	grid := importGrid(
		t,
		"=importrow(\"u\")\n",
		echoFetcher{body: []byte("a\tb\tc\n"), accept: &accept},
		sheet.DefaultLimits(),
	)
	assert.Equal(t, "a", cellAt(t, grid, 0, 0))
	assert.Equal(t, "b", cellAt(t, grid, 0, 1))
	assert.Equal(t, "c", cellAt(t, grid, 0, 2))
	assert.Equal(t, mediaRowWire, accept, "IMPORTROW must request the row media type")
}

func TestImportColumnSpillsVertically(t *testing.T) {
	t.Parallel()

	var accept sheet.MediaType
	grid := importGrid(
		t,
		"=importcolumn(\"u\")\n",
		echoFetcher{body: []byte("x\ny\nz\n"), accept: &accept},
		sheet.DefaultLimits(),
	)
	assert.Equal(t, "x", cellAt(t, grid, 0, 0))
	assert.Equal(t, "y", cellAt(t, grid, 1, 0))
	assert.Equal(t, "z", cellAt(t, grid, 2, 0))
	assert.Equal(t, mediaColumnWire, accept, "IMPORTCOLUMN must request the column media type")
}

func TestImportRangeSpillsRectangle(t *testing.T) {
	t.Parallel()

	var accept sheet.MediaType
	grid := importGrid(
		t,
		"=importrange(\"u\")\n",
		echoFetcher{body: []byte("1\t2\n3\t4\n"), accept: &accept},
		sheet.DefaultLimits(),
	)
	assert.Equal(t, "1", cellAt(t, grid, 0, 0))
	assert.Equal(t, "2", cellAt(t, grid, 0, 1))
	assert.Equal(t, "3", cellAt(t, grid, 1, 0))
	assert.Equal(t, "4", cellAt(t, grid, 1, 1))
	assert.Equal(t, mediaRangeWire, accept, "IMPORTRANGE must request the range media type")
}

func TestImportSheetSpillsLikeRange(t *testing.T) {
	t.Parallel()

	// For this engine chunk IMPORTSHEET spills like IMPORTRANGE; only the
	// requested Accept media type differs (the nested-grid rendering is deferred).
	var accept sheet.MediaType
	grid := importGrid(
		t,
		"=importsheet(\"u\")\n",
		echoFetcher{body: []byte("1\t2\n3\t4\n"), accept: &accept},
		sheet.DefaultLimits(),
	)
	assert.Equal(t, "1", cellAt(t, grid, 0, 0))
	assert.Equal(t, "4", cellAt(t, grid, 1, 1))
	assert.Equal(t, mediaSheetWire, accept, "IMPORTSHEET must request the sheet media type")
}

func TestImportLeadingEqualsStaysLiteral(t *testing.T) {
	t.Parallel()

	// A values-only import never compiles a cell: a leading `=` is literal text.
	grid := importGrid(t, "=importcell(\"u\")\n", echoFetcher{body: []byte("=A1\n")}, sheet.DefaultLimits())
	assert.Equal(t, "=A1", cellAt(t, grid, 0, 0))
}

func TestImportHandshakeMismatch(t *testing.T) {
	t.Parallel()

	// The server declares the cell type for an IMPORTROW request: a mismatch.
	f := fixedFetcher{result: sheet.FetchResult{Body: []byte("a\tb\n"), ContentType: mediaCellWire}}
	grid := importGrid(t, "=importrow(\"u\")\n", f, sheet.DefaultLimits())
	assert.Equal(t, "#IMPORT!", cellAt(t, grid, 0, 0))
}

func TestImportFetchError(t *testing.T) {
	t.Parallel()

	f := fixedFetcher{err: constants.ErrReadInput}
	grid := importGrid(t, "=importcell(\"u\")\n", f, sheet.DefaultLimits())
	assert.Equal(t, "#IMPORT!", cellAt(t, grid, 0, 0))
}

func TestImportTSVParseError(t *testing.T) {
	t.Parallel()

	// A single line longer than the scanner's 1 MiB token cap makes ReadTSV fail.
	body := bytes.Repeat([]byte("a"), (1<<20)+1)
	f := echoFetcher{body: body}
	grid := importGrid(t, "=importcell(\"u\")\n", f, sheet.DefaultLimits())
	assert.Equal(t, "#IMPORT!", cellAt(t, grid, 0, 0))
}

func TestImportEmptyBody(t *testing.T) {
	t.Parallel()

	f := echoFetcher{body: []byte("")}
	grid := importGrid(t, "=importcell(\"u\")\n", f, sheet.DefaultLimits())
	assert.Equal(t, "#IMPORT!", cellAt(t, grid, 0, 0))
}

func TestImportShapeMismatches(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		src  string
		body string
	}{
		"cell with two cells":   {"=importcell(\"u\")\n", "1\t2\n"},
		"cell with two rows":    {"=importcell(\"u\")\n", "1\n2\n"},
		"row with two rows":     {"=importrow(\"u\")\n", "a\nb\n"},
		"column with wide row":  {"=importcolumn(\"u\")\n", "1\n2\t3\n"},
		"range with ragged row": {"=importrange(\"u\")\n", "1\t2\n3\n"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			grid := importGrid(t, tc.src, echoFetcher{body: []byte(tc.body)}, sheet.DefaultLimits())
			assert.Equal(t, "#IMPORT!", cellAt(t, grid, 0, 0))
		})
	}
}

func TestImportOversizeRejected(t *testing.T) {
	t.Parallel()

	// A tiny cell budget rejects each spilling shape as #IMPORT! (oversize).
	tight := sheet.Limits{ResultCells: 2, GridDim: 20_000, ResultBytes: 64 << 10}
	cases := map[string]struct {
		src  string
		body string
	}{
		"row":    {"=importrow(\"u\")\n", "a\tb\tc\n"},
		"column": {"=importcolumn(\"u\")\n", "a\nb\nc\n"},
		"range":  {"=importrange(\"u\")\n", "1\t2\n3\t4\n"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			grid := importGrid(t, tc.src, echoFetcher{body: []byte(tc.body)}, tight)
			assert.Equal(t, "#IMPORT!", cellAt(t, grid, 0, 0))
		})
	}
}

func TestImportWrongArity(t *testing.T) {
	t.Parallel()

	for _, src := range []string{
		"=importcell()\n",
		"=importcell(\"a\", \"b\")\n",
	} {
		grid := importGrid(t, src, echoFetcher{body: []byte("1\n")}, sheet.DefaultLimits())
		assert.Equal(t, "#VALUE!", cellAt(t, grid, 0, 0), "wrong arity is #VALUE!: %s", src)
	}
}

func TestImportErrorValuedURLPropagates(t *testing.T) {
	t.Parallel()

	// The URL argument evaluates to #DIV/0!, which propagates unchanged — the
	// fetch never happens, so the result is the argument's error, not #IMPORT!.
	grid := importGrid(t, "=importcell(1/0)\n", echoFetcher{body: []byte("1\n")}, sheet.DefaultLimits())
	assert.Equal(t, "#DIV/0!", cellAt(t, grid, 0, 0))
}
