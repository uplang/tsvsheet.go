// Code generated from TsvsheetParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package tsvsheetgrammar // TsvsheetParser
import "github.com/antlr4-go/antlr/v4"

// BaseTsvsheetParserListener is a complete listener for a parse tree produced by TsvsheetParser.
type BaseTsvsheetParserListener struct{}

var _ TsvsheetParserListener = &BaseTsvsheetParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTsvsheetParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTsvsheetParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTsvsheetParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTsvsheetParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterWorksheet is called when production worksheet is entered.
func (s *BaseTsvsheetParserListener) EnterWorksheet(ctx *WorksheetContext) {}

// ExitWorksheet is called when production worksheet is exited.
func (s *BaseTsvsheetParserListener) ExitWorksheet(ctx *WorksheetContext) {}

// EnterLine is called when production line is entered.
func (s *BaseTsvsheetParserListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseTsvsheetParserListener) ExitLine(ctx *LineContext) {}

// EnterSectionCommand is called when production sectionCommand is entered.
func (s *BaseTsvsheetParserListener) EnterSectionCommand(ctx *SectionCommandContext) {}

// ExitSectionCommand is called when production sectionCommand is exited.
func (s *BaseTsvsheetParserListener) ExitSectionCommand(ctx *SectionCommandContext) {}

// EnterStructuralCommand is called when production structuralCommand is entered.
func (s *BaseTsvsheetParserListener) EnterStructuralCommand(ctx *StructuralCommandContext) {}

// ExitStructuralCommand is called when production structuralCommand is exited.
func (s *BaseTsvsheetParserListener) ExitStructuralCommand(ctx *StructuralCommandContext) {}

// EnterCells is called when production cells is entered.
func (s *BaseTsvsheetParserListener) EnterCells(ctx *CellsContext) {}

// ExitCells is called when production cells is exited.
func (s *BaseTsvsheetParserListener) ExitCells(ctx *CellsContext) {}

// EnterCell is called when production cell is entered.
func (s *BaseTsvsheetParserListener) EnterCell(ctx *CellContext) {}

// ExitCell is called when production cell is exited.
func (s *BaseTsvsheetParserListener) ExitCell(ctx *CellContext) {}

// EnterPayload is called when production payload is entered.
func (s *BaseTsvsheetParserListener) EnterPayload(ctx *PayloadContext) {}

// ExitPayload is called when production payload is exited.
func (s *BaseTsvsheetParserListener) ExitPayload(ctx *PayloadContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseTsvsheetParserListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseTsvsheetParserListener) ExitFormula(ctx *FormulaContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseTsvsheetParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseTsvsheetParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseTsvsheetParserListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseTsvsheetParserListener) ExitReference(ctx *ReferenceContext) {}

// EnterRangeRef is called when production rangeRef is entered.
func (s *BaseTsvsheetParserListener) EnterRangeRef(ctx *RangeRefContext) {}

// ExitRangeRef is called when production rangeRef is exited.
func (s *BaseTsvsheetParserListener) ExitRangeRef(ctx *RangeRefContext) {}

// EnterEndpoint is called when production endpoint is entered.
func (s *BaseTsvsheetParserListener) EnterEndpoint(ctx *EndpointContext) {}

// ExitEndpoint is called when production endpoint is exited.
func (s *BaseTsvsheetParserListener) ExitEndpoint(ctx *EndpointContext) {}

// EnterCellRef is called when production cellRef is entered.
func (s *BaseTsvsheetParserListener) EnterCellRef(ctx *CellRefContext) {}

// ExitCellRef is called when production cellRef is exited.
func (s *BaseTsvsheetParserListener) ExitCellRef(ctx *CellRefContext) {}

// EnterColRef is called when production colRef is entered.
func (s *BaseTsvsheetParserListener) EnterColRef(ctx *ColRefContext) {}

// ExitColRef is called when production colRef is exited.
func (s *BaseTsvsheetParserListener) ExitColRef(ctx *ColRefContext) {}

// EnterRowWildcard is called when production rowWildcard is entered.
func (s *BaseTsvsheetParserListener) EnterRowWildcard(ctx *RowWildcardContext) {}

// ExitRowWildcard is called when production rowWildcard is exited.
func (s *BaseTsvsheetParserListener) ExitRowWildcard(ctx *RowWildcardContext) {}

// EnterRowRef is called when production rowRef is entered.
func (s *BaseTsvsheetParserListener) EnterRowRef(ctx *RowRefContext) {}

// ExitRowRef is called when production rowRef is exited.
func (s *BaseTsvsheetParserListener) ExitRowRef(ctx *RowRefContext) {}

// EnterNumericRef is called when production numericRef is entered.
func (s *BaseTsvsheetParserListener) EnterNumericRef(ctx *NumericRefContext) {}

// ExitNumericRef is called when production numericRef is exited.
func (s *BaseTsvsheetParserListener) ExitNumericRef(ctx *NumericRefContext) {}

// EnterSignedInt is called when production signedInt is entered.
func (s *BaseTsvsheetParserListener) EnterSignedInt(ctx *SignedIntContext) {}

// ExitSignedInt is called when production signedInt is exited.
func (s *BaseTsvsheetParserListener) ExitSignedInt(ctx *SignedIntContext) {}

// EnterNumRow is called when production numRow is entered.
func (s *BaseTsvsheetParserListener) EnterNumRow(ctx *NumRowContext) {}

// ExitNumRow is called when production numRow is exited.
func (s *BaseTsvsheetParserListener) ExitNumRow(ctx *NumRowContext) {}

// EnterGroupedRange is called when production groupedRange is entered.
func (s *BaseTsvsheetParserListener) EnterGroupedRange(ctx *GroupedRangeContext) {}

// ExitGroupedRange is called when production groupedRange is exited.
func (s *BaseTsvsheetParserListener) ExitGroupedRange(ctx *GroupedRangeContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseTsvsheetParserListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseTsvsheetParserListener) ExitModifier(ctx *ModifierContext) {}

// EnterStringExpr is called when production stringExpr is entered.
func (s *BaseTsvsheetParserListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production stringExpr is exited.
func (s *BaseTsvsheetParserListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseTsvsheetParserListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseTsvsheetParserListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseTsvsheetParserListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseTsvsheetParserListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterRefExpr is called when production refExpr is entered.
func (s *BaseTsvsheetParserListener) EnterRefExpr(ctx *RefExprContext) {}

// ExitRefExpr is called when production refExpr is exited.
func (s *BaseTsvsheetParserListener) ExitRefExpr(ctx *RefExprContext) {}

// EnterNumberExpr is called when production numberExpr is entered.
func (s *BaseTsvsheetParserListener) EnterNumberExpr(ctx *NumberExprContext) {}

// ExitNumberExpr is called when production numberExpr is exited.
func (s *BaseTsvsheetParserListener) ExitNumberExpr(ctx *NumberExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseTsvsheetParserListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseTsvsheetParserListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseTsvsheetParserListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseTsvsheetParserListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseTsvsheetParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseTsvsheetParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterCompareExpr is called when production compareExpr is entered.
func (s *BaseTsvsheetParserListener) EnterCompareExpr(ctx *CompareExprContext) {}

// ExitCompareExpr is called when production compareExpr is exited.
func (s *BaseTsvsheetParserListener) ExitCompareExpr(ctx *CompareExprContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseTsvsheetParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseTsvsheetParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseTsvsheetParserListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseTsvsheetParserListener) ExitArgList(ctx *ArgListContext) {}
