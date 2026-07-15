package sheet_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uplang/tsvsheet.go/internal/sheet"
)

// num formats a float as tsvsheet renders it.
func num(x float64) string { return strconv.FormatFloat(x, 'f', -1, 64) }

func TestMath_HappyPaths(t *testing.T) {
	t.Parallel()

	// A1=2, C1=3, D1=4 (B1 holds the formula).
	cases := map[string]string{
		"pi()":                num(math.Pi),
		"sign(-D1)":           "-1",
		"sign(D1)":            "1",
		"sign(D1 - D1)":       "0",
		"int(3.9)":            "3",
		"trunc(-3.9)":         "-3",
		"sqrt(144)":           "12",
		"sqrtpi(1)":           num(math.Sqrt(math.Pi)),
		"power(A1, 10)":       "1024",
		"exp(0)":              "1",
		"ln(1)":               "0",
		"log10(1000)":         "3",
		"log(8, A1)":          "3", // base 2
		"log(100)":            "2", // default base 10
		"quotient(17, 5)":     "3",
		"product(A1, C1, D1)": "24",
		"sumsq(C1, D1)":       "25", // 9 + 16
		"cos(0)":              "1",
		"sin(0)":              "0",
		"tan(0)":              "0",
		"asin(0)":             "0",
		"acos(1)":             "0",
		"atan(0)":             "0",
		"atan2(1, 1)":         num(math.Atan2(1, 1)),
		"sinh(0)":             "0",
		"cosh(0)":             "1",
		"tanh(0)":             "0",
		"degrees(pi())":       "180",
		"radians(180)":        num(math.Pi),
	}
	for expr, want := range cases {
		t.Run(expr, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, want, formula1(t, expr))
		})
	}
}

func TestMath_DomainErrorsAreNum(t *testing.T) {
	t.Parallel()

	// A NaN or infinite result is #NUM! (Excel domain error / overflow).
	for _, expr := range []string{"sqrt(-1)", "ln(-1)", "asin(2)", "exp(1000)"} {
		t.Run(expr, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, string(sheet.ErrNum), formula1(t, expr))
		})
	}
}

func TestMath_ZeroAndEmptyEdges(t *testing.T) {
	t.Parallel()

	assert.Equal(t, string(sheet.ErrDiv), formula1(t, "quotient(1, 0)"))  // divide by zero
	assert.Equal(t, "1", cellAt(t, compute(t, "\t=product(A1)\n"), 0, 1)) // empty product is 1
}

func TestMath_NonNumericIsValueError(t *testing.T) {
	t.Parallel()

	// A1 is text; each numeric math builtin reports #VALUE! via its operands.
	cases := []string{
		"=sqrt(A1)", "=power(A1, 2)", "=power(2, A1)", "=log(A1)", "=log(8, A1)",
		"=quotient(A1, 1)", "=quotient(1, A1)", "=product(A1)", "=sumsq(A1)",
		"=atan2(A1, 1)", "=atan2(1, A1)", "=sign(A1)",
	}
	for _, expr := range cases {
		t.Run(expr, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, "#VALUE!", cellAt(t, compute(t, "hi\t"+expr+"\n"), 0, 1))
		})
	}
}
