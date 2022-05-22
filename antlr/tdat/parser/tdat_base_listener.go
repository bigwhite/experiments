// Code generated from Tdat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tdat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTdatListener is a complete listener for a parse tree produced by TdatParser.
type BaseTdatListener struct{}

var _ TdatListener = &BaseTdatListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTdatListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTdatListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTdatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTdatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseTdatListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseTdatListener) ExitProg(ctx *ProgContext) {}

// EnterRuleLine is called when production ruleLine is entered.
func (s *BaseTdatListener) EnterRuleLine(ctx *RuleLineContext) {}

// ExitRuleLine is called when production ruleLine is exited.
func (s *BaseTdatListener) ExitRuleLine(ctx *RuleLineContext) {}

// EnterRuleID is called when production ruleID is entered.
func (s *BaseTdatListener) EnterRuleID(ctx *RuleIDContext) {}

// ExitRuleID is called when production ruleID is exited.
func (s *BaseTdatListener) ExitRuleID(ctx *RuleIDContext) {}

// EnterEnumerableFunc is called when production enumerableFunc is entered.
func (s *BaseTdatListener) EnterEnumerableFunc(ctx *EnumerableFuncContext) {}

// ExitEnumerableFunc is called when production enumerableFunc is exited.
func (s *BaseTdatListener) ExitEnumerableFunc(ctx *EnumerableFuncContext) {}

// EnterWindowsWithZeroIndex is called when production WindowsWithZeroIndex is entered.
func (s *BaseTdatListener) EnterWindowsWithZeroIndex(ctx *WindowsWithZeroIndexContext) {}

// ExitWindowsWithZeroIndex is called when production WindowsWithZeroIndex is exited.
func (s *BaseTdatListener) ExitWindowsWithZeroIndex(ctx *WindowsWithZeroIndexContext) {}

// EnterWindowsWithLowAndHighIndex is called when production WindowsWithLowAndHighIndex is entered.
func (s *BaseTdatListener) EnterWindowsWithLowAndHighIndex(ctx *WindowsWithLowAndHighIndexContext) {}

// ExitWindowsWithLowAndHighIndex is called when production WindowsWithLowAndHighIndex is exited.
func (s *BaseTdatListener) ExitWindowsWithLowAndHighIndex(ctx *WindowsWithLowAndHighIndexContext) {}

// EnterConditionExpr is called when production conditionExpr is entered.
func (s *BaseTdatListener) EnterConditionExpr(ctx *ConditionExprContext) {}

// ExitConditionExpr is called when production conditionExpr is exited.
func (s *BaseTdatListener) ExitConditionExpr(ctx *ConditionExprContext) {}

// EnterBracketExprInPrimaryExpr is called when production BracketExprInPrimaryExpr is entered.
func (s *BaseTdatListener) EnterBracketExprInPrimaryExpr(ctx *BracketExprInPrimaryExprContext) {}

// ExitBracketExprInPrimaryExpr is called when production BracketExprInPrimaryExpr is exited.
func (s *BaseTdatListener) ExitBracketExprInPrimaryExpr(ctx *BracketExprInPrimaryExprContext) {}

// EnterRightLiteralInPrimaryExpr is called when production RightLiteralInPrimaryExpr is entered.
func (s *BaseTdatListener) EnterRightLiteralInPrimaryExpr(ctx *RightLiteralInPrimaryExprContext) {}

// ExitRightLiteralInPrimaryExpr is called when production RightLiteralInPrimaryExpr is exited.
func (s *BaseTdatListener) ExitRightLiteralInPrimaryExpr(ctx *RightLiteralInPrimaryExprContext) {}

// EnterArithmeticExprInPrimaryExpr is called when production ArithmeticExprInPrimaryExpr is entered.
func (s *BaseTdatListener) EnterArithmeticExprInPrimaryExpr(ctx *ArithmeticExprInPrimaryExprContext) {
}

// ExitArithmeticExprInPrimaryExpr is called when production ArithmeticExprInPrimaryExpr is exited.
func (s *BaseTdatListener) ExitArithmeticExprInPrimaryExpr(ctx *ArithmeticExprInPrimaryExprContext) {}

// EnterBuildinExprInPrimaryExpr is called when production BuildinExprInPrimaryExpr is entered.
func (s *BaseTdatListener) EnterBuildinExprInPrimaryExpr(ctx *BuildinExprInPrimaryExprContext) {}

// ExitBuildinExprInPrimaryExpr is called when production BuildinExprInPrimaryExpr is exited.
func (s *BaseTdatListener) ExitBuildinExprInPrimaryExpr(ctx *BuildinExprInPrimaryExprContext) {}

// EnterMetricInPrimaryExpr is called when production MetricInPrimaryExpr is entered.
func (s *BaseTdatListener) EnterMetricInPrimaryExpr(ctx *MetricInPrimaryExprContext) {}

// ExitMetricInPrimaryExpr is called when production MetricInPrimaryExpr is exited.
func (s *BaseTdatListener) ExitMetricInPrimaryExpr(ctx *MetricInPrimaryExprContext) {}

// EnterArithmeticOp is called when production arithmeticOp is entered.
func (s *BaseTdatListener) EnterArithmeticOp(ctx *ArithmeticOpContext) {}

// ExitArithmeticOp is called when production arithmeticOp is exited.
func (s *BaseTdatListener) ExitArithmeticOp(ctx *ArithmeticOpContext) {}

// EnterBuiltin is called when production builtin is entered.
func (s *BaseTdatListener) EnterBuiltin(ctx *BuiltinContext) {}

// ExitBuiltin is called when production builtin is exited.
func (s *BaseTdatListener) ExitBuiltin(ctx *BuiltinContext) {}

// EnterLogicalOp is called when production logicalOp is entered.
func (s *BaseTdatListener) EnterLogicalOp(ctx *LogicalOpContext) {}

// ExitLogicalOp is called when production logicalOp is exited.
func (s *BaseTdatListener) ExitLogicalOp(ctx *LogicalOpContext) {}

// EnterComparisonOp is called when production comparisonOp is entered.
func (s *BaseTdatListener) EnterComparisonOp(ctx *ComparisonOpContext) {}

// ExitComparisonOp is called when production comparisonOp is exited.
func (s *BaseTdatListener) ExitComparisonOp(ctx *ComparisonOpContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseTdatListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseTdatListener) ExitLiteral(ctx *LiteralContext) {}

// EnterResultWithElements is called when production ResultWithElements is entered.
func (s *BaseTdatListener) EnterResultWithElements(ctx *ResultWithElementsContext) {}

// ExitResultWithElements is called when production ResultWithElements is exited.
func (s *BaseTdatListener) ExitResultWithElements(ctx *ResultWithElementsContext) {}

// EnterResultWithoutElements is called when production ResultWithoutElements is entered.
func (s *BaseTdatListener) EnterResultWithoutElements(ctx *ResultWithoutElementsContext) {}

// ExitResultWithoutElements is called when production ResultWithoutElements is exited.
func (s *BaseTdatListener) ExitResultWithoutElements(ctx *ResultWithoutElementsContext) {}
