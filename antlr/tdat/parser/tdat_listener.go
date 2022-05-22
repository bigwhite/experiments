// Code generated from Tdat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tdat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TdatListener is a complete listener for a parse tree produced by TdatParser.
type TdatListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterRuleLine is called when entering the ruleLine production.
	EnterRuleLine(c *RuleLineContext)

	// EnterRuleID is called when entering the ruleID production.
	EnterRuleID(c *RuleIDContext)

	// EnterEnumerableFunc is called when entering the enumerableFunc production.
	EnterEnumerableFunc(c *EnumerableFuncContext)

	// EnterWindowsWithZeroIndex is called when entering the WindowsWithZeroIndex production.
	EnterWindowsWithZeroIndex(c *WindowsWithZeroIndexContext)

	// EnterWindowsWithLowAndHighIndex is called when entering the WindowsWithLowAndHighIndex production.
	EnterWindowsWithLowAndHighIndex(c *WindowsWithLowAndHighIndexContext)

	// EnterConditionExpr is called when entering the conditionExpr production.
	EnterConditionExpr(c *ConditionExprContext)

	// EnterBracketExprInPrimaryExpr is called when entering the BracketExprInPrimaryExpr production.
	EnterBracketExprInPrimaryExpr(c *BracketExprInPrimaryExprContext)

	// EnterRightLiteralInPrimaryExpr is called when entering the RightLiteralInPrimaryExpr production.
	EnterRightLiteralInPrimaryExpr(c *RightLiteralInPrimaryExprContext)

	// EnterArithmeticExprInPrimaryExpr is called when entering the ArithmeticExprInPrimaryExpr production.
	EnterArithmeticExprInPrimaryExpr(c *ArithmeticExprInPrimaryExprContext)

	// EnterBuildinExprInPrimaryExpr is called when entering the BuildinExprInPrimaryExpr production.
	EnterBuildinExprInPrimaryExpr(c *BuildinExprInPrimaryExprContext)

	// EnterMetricInPrimaryExpr is called when entering the MetricInPrimaryExpr production.
	EnterMetricInPrimaryExpr(c *MetricInPrimaryExprContext)

	// EnterArithmeticOp is called when entering the arithmeticOp production.
	EnterArithmeticOp(c *ArithmeticOpContext)

	// EnterBuiltin is called when entering the builtin production.
	EnterBuiltin(c *BuiltinContext)

	// EnterLogicalOp is called when entering the logicalOp production.
	EnterLogicalOp(c *LogicalOpContext)

	// EnterComparisonOp is called when entering the comparisonOp production.
	EnterComparisonOp(c *ComparisonOpContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterResultWithElements is called when entering the ResultWithElements production.
	EnterResultWithElements(c *ResultWithElementsContext)

	// EnterResultWithoutElements is called when entering the ResultWithoutElements production.
	EnterResultWithoutElements(c *ResultWithoutElementsContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitRuleLine is called when exiting the ruleLine production.
	ExitRuleLine(c *RuleLineContext)

	// ExitRuleID is called when exiting the ruleID production.
	ExitRuleID(c *RuleIDContext)

	// ExitEnumerableFunc is called when exiting the enumerableFunc production.
	ExitEnumerableFunc(c *EnumerableFuncContext)

	// ExitWindowsWithZeroIndex is called when exiting the WindowsWithZeroIndex production.
	ExitWindowsWithZeroIndex(c *WindowsWithZeroIndexContext)

	// ExitWindowsWithLowAndHighIndex is called when exiting the WindowsWithLowAndHighIndex production.
	ExitWindowsWithLowAndHighIndex(c *WindowsWithLowAndHighIndexContext)

	// ExitConditionExpr is called when exiting the conditionExpr production.
	ExitConditionExpr(c *ConditionExprContext)

	// ExitBracketExprInPrimaryExpr is called when exiting the BracketExprInPrimaryExpr production.
	ExitBracketExprInPrimaryExpr(c *BracketExprInPrimaryExprContext)

	// ExitRightLiteralInPrimaryExpr is called when exiting the RightLiteralInPrimaryExpr production.
	ExitRightLiteralInPrimaryExpr(c *RightLiteralInPrimaryExprContext)

	// ExitArithmeticExprInPrimaryExpr is called when exiting the ArithmeticExprInPrimaryExpr production.
	ExitArithmeticExprInPrimaryExpr(c *ArithmeticExprInPrimaryExprContext)

	// ExitBuildinExprInPrimaryExpr is called when exiting the BuildinExprInPrimaryExpr production.
	ExitBuildinExprInPrimaryExpr(c *BuildinExprInPrimaryExprContext)

	// ExitMetricInPrimaryExpr is called when exiting the MetricInPrimaryExpr production.
	ExitMetricInPrimaryExpr(c *MetricInPrimaryExprContext)

	// ExitArithmeticOp is called when exiting the arithmeticOp production.
	ExitArithmeticOp(c *ArithmeticOpContext)

	// ExitBuiltin is called when exiting the builtin production.
	ExitBuiltin(c *BuiltinContext)

	// ExitLogicalOp is called when exiting the logicalOp production.
	ExitLogicalOp(c *LogicalOpContext)

	// ExitComparisonOp is called when exiting the comparisonOp production.
	ExitComparisonOp(c *ComparisonOpContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitResultWithElements is called when exiting the ResultWithElements production.
	ExitResultWithElements(c *ResultWithElementsContext)

	// ExitResultWithoutElements is called when exiting the ResultWithoutElements production.
	ExitResultWithoutElements(c *ResultWithoutElementsContext)
}
