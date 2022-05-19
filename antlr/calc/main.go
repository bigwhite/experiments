package main

import (
	"calc/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// calc takes a string expression and returns the evaluated result.
func calc(input string) int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()
}

func main() {
	println(calc("1 + 2 * 3"))
	println(calc("12 * 3 / 6"))
}
