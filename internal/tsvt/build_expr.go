package tsvt

import (
	"github.com/antlr4-go/antlr/v4"

	grammar "github.com/uplang/tsvsheet.go/src/grammar/tsvsheet"
)

// buildExpr converts a §11 expression by its labeled alternative. The grammar's
// stringExpr alternative is unreachable — a STRING operand always matches
// refExpr first as a named-column reference (§11.4 / ADR 0003 rule 16) — so
// numberExpr is the only remaining reachable alternative and forms the default.
func buildExpr(ctx grammar.IExpressionContext) (Expr, error) {
	switch c := ctx.(type) {
	case *grammar.ParenExprContext:
		return buildExpr(c.Expression())
	case *grammar.UnaryExprContext:
		return buildUnary(c)
	case *grammar.MulExprContext:
		return buildBinary(c.GetOp(), c.AllExpression())
	case *grammar.AddExprContext:
		return buildBinary(c.GetOp(), c.AllExpression())
	case *grammar.CompareExprContext:
		return buildBinary(c.GetOp(), c.AllExpression())
	case *grammar.CallExprContext:
		return buildCall(c.FunctionCall())
	case *grammar.RefExprContext:
		return buildRefOperand(c)
	default: // NumberExprContext
		return Number{Text: ctx.GetText()}, nil
	}
}

// buildUnary converts a unary sign expression.
func buildUnary(ctx *grammar.UnaryExprContext) (Expr, error) {
	operand, err := buildExpr(ctx.Expression())
	if err != nil {
		return nil, err
	}
	return Unary{Op: UnaryOp(ctx.GetOp().GetText()), X: operand}, nil
}

// buildBinary converts a two-operand expression; the operator token text is
// exactly the BinaryOp spelling.
func buildBinary(op antlr.Token, operands []grammar.IExpressionContext) (Expr, error) {
	left, err := buildExpr(operands[0])
	if err != nil {
		return nil, err
	}
	right, err := buildExpr(operands[1])
	if err != nil {
		return nil, err
	}
	return Binary{Op: BinaryOp(op.GetText()), Left: left, Right: right}, nil
}

// buildRefOperand converts a reference used as an operand.
func buildRefOperand(ctx *grammar.RefExprContext) (Expr, error) {
	ref, err := buildReference(ctx.Reference())
	if err != nil {
		return nil, err
	}
	return RefOperand{Ref: ref}, nil
}

// buildCall converts a function call; the name token is NAME or an uppercase
// COL (`IF`).
func buildCall(ctx grammar.IFunctionCallContext) (Expr, error) {
	args, err := buildArgs(ctx.ArgList())
	if err != nil {
		return nil, err
	}
	return Call{Name: callName(ctx), Args: args}, nil
}

// callName extracts the case-preserved function name.
func callName(ctx grammar.IFunctionCallContext) string {
	if name := ctx.NAME(); name != nil {
		return name.GetText()
	}
	return ctx.COL().GetText()
}

// buildArgs converts an optional argument list.
func buildArgs(ctx grammar.IArgListContext) ([]Expr, error) {
	if ctx == nil {
		return nil, nil
	}
	contexts := ctx.AllExpression()
	args := make([]Expr, 0, len(contexts))
	for _, c := range contexts {
		arg, err := buildExpr(c)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	return args, nil
}
