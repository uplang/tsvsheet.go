package sheet

import (
	"strings"

	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

// IsVolatile reports whether any formula calls a clock-dependent function
// (TODAY, NOW, ISNOW) — so a frontend knows its computed values change over
// time and can offer periodic recomputation.
func (s Sheet) IsVolatile() bool {
	volatile := false
	s.eachFormula(func(at Address) {
		walkCalls(s.cells[at.Row][at.Col].formula, func(call tsvt.Call) {
			if isVolatileCall(call) {
				volatile = true
			}
		})
	})
	return volatile
}

// isVolatileCall reports whether a call is to a clock-dependent function.
func isVolatileCall(call tsvt.Call) boolResult {
	switch strings.ToLower(call.Name) {
	case fnToday, fnNow, fnIsnow:
		return true
	default:
		return false
	}
}
