package main

import (
	"fmt"
	"tdat/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TraceListener struct {
	*parser.BaseTdatListener
	p *parser.TdatParser
	t antlr.Tree
}

func NewTraceListener(p *parser.TdatParser, t antlr.Tree) *TraceListener {
	return &TraceListener{
		p: p,
		t: t,
	}
}

func (l *TraceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	printLevelPrefix(ctx)
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	fmt.Printf("==> %s 《 %s 》\n", ruleName, ctx.GetText())
}

func (l *TraceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	printLevelPrefix(ctx)
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	fmt.Println("<==", ruleName)
}

func printLevelPrefix(ctx antlr.ParserRuleContext) {
	level := 0

	t := ctx.GetParent()
	for t != nil {
		level++
		t = t.GetParent()
	}

	for i := 0; i < level; i++ {
		fmt.Printf("\t")
	}
}
