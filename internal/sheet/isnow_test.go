package sheet_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// isnowAt computes a single ISNOW formula against a fixed instant.
func isnowAt(t *testing.T, formula string, at time.Time) string {
	t.Helper()
	s, err := sheet.Parse([]byte("=" + formula + "\n"))
	require.NoError(t, err)
	return s.ComputeAt(at)[0][0]
}

func TestIsnow_MatchesTheClock(t *testing.T) {
	t.Parallel()

	noon := time.Date(2026, 1, 5, 12, 0, 0, 0, time.UTC) // a Monday, 12:00
	assert.Equal(t, "TRUE", isnowAt(t, `isnow("noon")`, noon))
	assert.Equal(t, "FALSE", isnowAt(t, `isnow("midnight")`, noon))
}

func TestIsnow_Errors(t *testing.T) {
	t.Parallel()

	at := time.Date(2026, 1, 5, 12, 0, 0, 0, time.UTC)
	assert.Equal(t, "#VALUE!", isnowAt(t, `isnow()`, at))      // arity
	assert.Equal(t, "#VALUE!", isnowAt(t, `isnow("$$$")`, at)) // malformed pattern
	assert.Equal(t, "#DIV/0!", isnowAt(t, `isnow(1/0)`, at))   // argument error propagates
}

func TestIsnow_IsKnownToCheck(t *testing.T) {
	t.Parallel()

	// ISNOW is a known function — Check must not flag it.
	s, err := sheet.Parse([]byte(`=isnow("M-F noon")` + "\n"))
	require.NoError(t, err)
	assert.Empty(t, sheet.Check(s))
}
