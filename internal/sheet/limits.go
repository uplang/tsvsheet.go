package sheet

// Resource limits guard against out-of-memory from untrusted formula or edit
// input: no single cell may drive an unbounded array, string, or grid
// allocation. They are process-wide and set once at program start — the CLI
// keeps DefaultLimits (or honors --max-cells), the WASM build applies the
// smaller BrowserLimits — then only read, so the concurrent reads on the server
// path are race-free. A package-level resource limit mirrors the standard
// library (bufio.MaxScanTokenSize, http.DefaultMaxHeaderBytes); threading it
// through every builtin, the resolver, and the immutable grid would be far more
// invasive for no added safety.

// Limits bounds the sizes an untrusted sheet may drive an allocation to.
type Limits struct {
	ResultCells int // cells in one array formula result (e.g. SEQUENCE)
	GridDim     int // the highest row or column index the grid may grow to (Set)
	ResultBytes int // bytes in one string formula result (e.g. REPT)
}

// DefaultLimits are generous for real spreadsheets while still bounding OOM.
var DefaultLimits = Limits{ResultCells: 5_000_000, GridDim: 1_000_000, ResultBytes: 1 << 20}

// BrowserLimits are the tighter ceilings the WASM build applies, sized for a
// browser tab rather than a workstation.
var BrowserLimits = Limits{ResultCells: 100_000, GridDim: 20_000, ResultBytes: 64 << 10}

var active = DefaultLimits

// SetLimits replaces the process-wide limits. Call once at program start,
// before any sheet is computed, edited, or served.
func SetLimits(l Limits) { active = l }

// tooManyCells reports whether an rows×cols array result exceeds the cell
// budget (computed in int64 so the product cannot overflow).
func tooManyCells(rows, cols int) bool {
	return int64(rows)*int64(cols) > int64(active.ResultCells)
}
