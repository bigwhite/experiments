package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type VerboseErrorListener struct {
	*antlr.DefaultErrorListener
	hasError bool
}

func NewVerboseErrorListener() *VerboseErrorListener {
	return new(VerboseErrorListener)
}

func (d *VerboseErrorListener) HasError() bool {
	return d.hasError
}
func (d *VerboseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	p := recognizer.(antlr.Parser)
	stack := p.GetRuleInvocationStack(p.GetParserRuleContext())

	fmt.Printf("rule: %v ", stack[0])
	fmt.Printf("line %d: %d at %v : %s\n", line, column, offendingSymbol, msg)

	d.hasError = true
}
