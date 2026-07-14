package tsvt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/uplang/tsvsheet.go/internal/tsvt"
)

func TestExpr_Number(t *testing.T) {
	t.Parallel()

	assert.Equal(t, tsvt.Number{Text: "42"}, formulaExpr(t, "42"))
	assert.Equal(t, tsvt.Number{Text: "3.14"}, formulaExpr(t, "3.14"))
}

func TestExpr_QuotedTokenIsNamedReference(t *testing.T) {
	t.Parallel()

	// A quoted operand parses as a named-column reference, not a string literal
	// (ADR 0003 rule 16); the string-literal role is recovered semantically.
	ref, ok := formulaExpr(t, `"hi"`).(tsvt.RefOperand)
	require.True(t, ok)
	assert.Equal(t, tsvt.ColNamed{Name: "hi"}, refCol(t, ref.Ref))
}

func TestExpr_RefOperand(t *testing.T) {
	t.Parallel()

	ref, ok := formulaExpr(t, "C").(tsvt.RefOperand)
	require.True(t, ok)
	assert.IsType(t, tsvt.RangeRef{}, ref.Ref)
}

func TestExpr_Paren(t *testing.T) {
	t.Parallel()

	// Grouping unwraps to the inner expression.
	assert.Equal(t, tsvt.Number{Text: "5"}, formulaExpr(t, "(5)"))
}

func TestExpr_Unary(t *testing.T) {
	t.Parallel()

	cases := map[string]tsvt.UnaryOp{
		"-5": tsvt.OpNeg,
		"+5": tsvt.OpPos,
	}
	for src, op := range cases {
		t.Run(src, func(t *testing.T) {
			t.Parallel()
			unary, ok := formulaExpr(t, src).(tsvt.Unary)
			require.True(t, ok)
			assert.Equal(t, op, unary.Op)
			assert.Equal(t, tsvt.Number{Text: "5"}, unary.X)
		})
	}
}

func TestExpr_BinaryOperators(t *testing.T) {
	t.Parallel()

	cases := map[string]tsvt.BinaryOp{
		"1 * 2":  tsvt.OpMul,
		"1 / 2":  tsvt.OpDiv,
		"1 % 2":  tsvt.OpMod,
		"1 + 2":  tsvt.OpAdd,
		"1 - 2":  tsvt.OpSub,
		"1 = 2":  tsvt.OpEq,
		"1 <> 2": tsvt.OpNe,
		"1 < 2":  tsvt.OpLt,
		"1 <= 2": tsvt.OpLe,
		"1 > 2":  tsvt.OpGt,
		"1 >= 2": tsvt.OpGe,
	}
	for src, op := range cases {
		t.Run(src, func(t *testing.T) {
			t.Parallel()
			binary, ok := formulaExpr(t, src).(tsvt.Binary)
			require.True(t, ok)
			assert.Equal(t, op, binary.Op)
		})
	}
}

func TestExpr_Precedence(t *testing.T) {
	t.Parallel()

	// min(A0:A1) + sum(B1:B2) * E groups as min(...) + (sum(...) * E) (§11.2).
	expr := formulaExpr(t, "min(A0:A1) + sum(B1:B2) * E")
	add, ok := expr.(tsvt.Binary)
	require.True(t, ok)
	assert.Equal(t, tsvt.OpAdd, add.Op)

	assert.IsType(t, tsvt.Call{}, add.Left)
	mul, ok := add.Right.(tsvt.Binary)
	require.True(t, ok)
	assert.Equal(t, tsvt.OpMul, mul.Op)
}

func TestExpr_CallLowercase(t *testing.T) {
	t.Parallel()

	call, ok := formulaExpr(t, "sum(A:H)").(tsvt.Call)
	require.True(t, ok)
	assert.Equal(t, "sum", call.Name)
	require.Len(t, call.Args, 1)
	assert.IsType(t, tsvt.RefOperand{}, call.Args[0])
}

func TestExpr_CallUppercaseKeyword(t *testing.T) {
	t.Parallel()

	// IF is lexed as a COL token but is a valid function name (§11.3).
	call, ok := formulaExpr(t, "IF(A,C$3,D$3)").(tsvt.Call)
	require.True(t, ok)
	assert.Equal(t, "IF", call.Name)
	assert.Len(t, call.Args, 3)
}

func TestExpr_CallNoArgs(t *testing.T) {
	t.Parallel()

	call, ok := formulaExpr(t, "now()").(tsvt.Call)
	require.True(t, ok)
	assert.Equal(t, "now", call.Name)
	assert.Empty(t, call.Args)
}

func TestExpr_CallMultiArg(t *testing.T) {
	t.Parallel()

	call, ok := formulaExpr(t, "min(1,2,3)").(tsvt.Call)
	require.True(t, ok)
	assert.Len(t, call.Args, 3)
}
