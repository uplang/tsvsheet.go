package tsvt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/constants"
	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

// A fractional NUMBER in any integer position (row index, header count, column
// index) is a valid NUMBER token but not a valid integer, so buildTemplate
// rejects it with ErrSyntax. Each case drives that rejection through a distinct
// build path, covering every error-propagation branch.
func TestBuild_FractionalRejectedEverywhere(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"structural command ref": "=C1.5<",
		"placement ref":          "C1.5",
		"placement payload":      "E=C1.5",
		"unary operand":          "=-C1.5",
		"binary left operand":    "=C1.5 + D",
		"binary right operand":   "=D + C1.5",
		"ref operand":            "=C1.5",
		"call argument":          "=sum(C1.5)",
		"range from endpoint":    "C1.5:E",
		"range to endpoint":      "C:E1.5",
		"row wildcard":           "*1.5",
		"row after":              "C+1.5",
		"row last offset":        "C$1.5",
		"numeric column":         "[1.5]",
		"numeric row":            "[2,1.5]",
		"grouped range row":      "(C:E)1.5",
		"grouped numeric from":   "([1.5]:[5])",
		"grouped numeric to":     "([3]:[1.5])",
	}
	for name, src := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := tsvt.Parse(tsvt.Source(src))
			require.Error(t, err, "expected %q to be rejected", src)
			assert.ErrorIs(t, err, constants.ErrSyntax)
		})
	}
}
