package main

import (
	"fmt"
	"os"

	"tdat/parser"
	"tdat/semantic"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func extractReversePolishExpr(listener antlr.ParseTreeListener, t antlr.Tree) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("semantic tree assembly error: %v", x)
		}
	}()

	antlr.ParseTreeWalkerDefault.Walk(listener, t)

	return nil
}

func main() {
	println("input file:", os.Args[1])
	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		panic(err)
	}

	lexer := parser.NewTdatLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewTdatParser(stream)

	p.RemoveErrorListeners()
	el := NewVerboseErrorListener()
	p.AddErrorListener(el)

	tree := p.Prog()

	if el.HasError() {
		return
	}

	l := NewReversePolishExprListener()
	err = extractReversePolishExpr(l, tree)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	processor := &Processor{
		name:  l.ruleID,
		model: semantic.NewModel(l.reversePolishExpr, semantic.NewWindowsRange(l.low, l.high), l.ef, l.result),
	}

	// r0006: Each { |1,3| ($speed < 50) and (($temperature + 1) < 4) or ((roundDown($salinity) <= 600.0) or (roundUp($ph) > 8.0)) } => ();

	in := []map[string]interface{}{
		{
			"speed":       30,
			"temperature": 6,
			"salinity":    500.0,
			"ph":          7.0,
		},
		{
			"speed":       31,
			"temperature": 7,
			"salinity":    501.0,
			"ph":          7.1,
		},
		{
			"speed":       30,
			"temperature": 6,
			"salinity":    498.0,
			"ph":          6.9,
		},
	}

	out, err := processor.Exec(in)
	if err != nil {
		fmt.Printf("processor execute error: %v", err)
		return
	}
	fmt.Printf("%v\n", out)

}
