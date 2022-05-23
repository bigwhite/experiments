package main

import (
	"fmt"
	"strconv"
	"strings"
	"tdat/parser"
	"tdat/semantic"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Item struct {
	level int
	val   semantic.Value
}

const (
	windowsRangeMax = 300
)

type ReversePolishExprListener struct {
	*parser.BaseTdatListener

	ruleID string

	// for constructing Reverse Polish expression
	//
	// infixExpr:($speed<5)and($temperature<2)or(roundDown($sanility)<600) =>
	//
	// reversePolishExpr:
	// $speed,5,<,$temperature,2,<,and,$sanility,roundDown,600,<,or
	//
	reversePolishExpr []semantic.Value
	s1                semantic.Stack[*Item] // temp stack for constructing reversePolishExpr, for final result
	s2                semantic.Stack[*Item] // temp stack for constructing reversePolishExpr, for operator temporarily

	// for windowsRange
	low  int
	high int

	// for enumerableFunc
	ef string

	// for result
	result []string
}

func NewReversePolishExprListener() *ReversePolishExprListener {
	return &ReversePolishExprListener{}
}

func (l *ReversePolishExprListener) ExitRuleID(c *parser.RuleIDContext) {
	l.ruleID = c.GetText()
}

func (l *ReversePolishExprListener) ExitEnumerableFunc(c *parser.EnumerableFuncContext) {
	s := c.GetText()
	switch s {
	case "None":
		l.ef = "none"
	case "Any":
		l.ef = "any"
	case "Each":
		l.ef = "each"
	}
}

func (l *ReversePolishExprListener) ExitResultWithElements(c *parser.ResultWithElementsContext) {
	s := c.GetText()
	l.result = strings.Split(s, ",")
}

func (l *ReversePolishExprListener) ExitResultWithoutElements(c *parser.ResultWithoutElementsContext) {
	l.result = []string{}
}

func (l *ReversePolishExprListener) ExitWindowsWithZeroIndex(c *parser.WindowsWithZeroIndexContext) {
	l.low = 1
	l.high = 1
}

func (l *ReversePolishExprListener) ExitWindowsWithLowAndHighIndex(c *parser.WindowsWithLowAndHighIndexContext) {
	s := c.GetText()
	s = s[1 : len(s)-1] // remove two '|'

	t := strings.Split(s, ",")

	if t[0] == "" {
		l.low = 1
	} else {
		l.low, _ = strconv.Atoi(t[0])
	}

	if t[1] == "" {
		l.high = windowsRangeMax
	} else {
		l.high, _ = strconv.Atoi(t[1])
	}

	if l.high < l.low {
		panic(fmt.Sprintf("windowsRange: low(%d) > high(%d)", l.low, l.high))
	}
}

func (l *ReversePolishExprListener) ExitConditionExpr(c *parser.ConditionExprContext) {
	// get the rule index of parent context
	if i, ok := c.GetParent().(antlr.RuleContext); ok {
		if i.GetRuleIndex() != parser.TdatParserRULE_ruleLine {
			// ignore this node
			return
		}
	}

	// pop all left in the stack
	for l.s2.Len() != 0 {
		l.s1.Push(l.s2.Pop())
	}

	// fill in the reversePolishExpr
	var vs []semantic.Value
	for l.s1.Len() != 0 {
		vs = append(vs, l.s1.Pop().val)
	}

	for i := len(vs) - 1; i >= 0; i-- {
		l.reversePolishExpr = append(l.reversePolishExpr, vs[i])
	}

	t := l.reversePolishExpr
	fmt.Printf("===>\n")
	for _, i := range t {
		fmt.Printf("%v,", i.Value())
	}
	fmt.Printf("\n<===\n")
}

func getLevel(t antlr.Tree) int {
	level := 0
	t = t.GetParent()
	for t != nil {
		level++
		t = t.GetParent()
	}
	return level
}

func (l *ReversePolishExprListener) handleBinOperator(c *antlr.BaseParserRuleContext) {
	v := c.GetText()
	lvl := getLevel(c)

	for {
		lastOp := l.s2.Top()
		if lastOp == nil {
			l.s2.Push(&Item{
				level: lvl,
				val: &semantic.BinaryOperator{
					Val: v,
				},
			})
			return
		}

		if lvl > lastOp.level {
			l.s2.Push(&Item{
				level: lvl,
				val: &semantic.BinaryOperator{
					Val: v,
				},
			})
			return
		}
		l.s1.Push(l.s2.Pop())
	}
}

func (l *ReversePolishExprListener) ExitLogicalOp(c *parser.LogicalOpContext) {
	l.handleBinOperator(c.BaseParserRuleContext)
}

func (l *ReversePolishExprListener) ExitArithmeticOp(c *parser.ArithmeticOpContext) {
	l.handleBinOperator(c.BaseParserRuleContext)
}

func (l *ReversePolishExprListener) ExitComparisonOp(c *parser.ComparisonOpContext) {
	l.handleBinOperator(c.BaseParserRuleContext)
}

func (l *ReversePolishExprListener) ExitMetricInPrimaryExpr(c *parser.MetricInPrimaryExprContext) {
	v := c.GetText()
	s := &Item{
		val: &semantic.Variable{
			Val: v[1:],
		},
	}
	l.s1.Push(s)
}

func (l *ReversePolishExprListener) ExitLiteral(c *parser.LiteralContext) {
	v := c.GetText()
	t := c.GetStop().GetTokenType()
	switch t {
	case parser.TdatLexerINT:
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		s := &Item{
			val: &semantic.Literal{
				Val: i,
			},
		}
		l.s1.Push(s)

	case parser.TdatLexerFLOAT:
		f, err := strconv.ParseFloat(v, 64) // float64
		if err != nil {
			panic(err)
		}
		s := &Item{
			val: &semantic.Literal{
				Val: f,
			},
		}
		l.s1.Push(s)

	case parser.TdatLexerSTRING:
		s := &Item{
			val: &semantic.Literal{
				Val: v,
			},
		}
		l.s1.Push(s)
	}
}

func (l *ReversePolishExprListener) ExitBuiltin(c *parser.BuiltinContext) {
	v := c.GetText()
	lvl := getLevel(c)

	l.s2.Push(&Item{
		level: lvl,
		val: &semantic.UnaryOperator{
			Val: v,
		},
	})
}
