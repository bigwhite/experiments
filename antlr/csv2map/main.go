package main

import (
	"csvparser/parser"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	csvFile := os.Args[1]
	is, err := antlr.NewFileStream(csvFile)
	if err != nil {
		fmt.Printf("new file stream error: %s\n", err)
		return
	}

	// Create the Lexer
	lexer := parser.NewCSVLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCSVParser(stream)

	// Finally parse the expression
	l := &CSVListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, p.CsvFile())
	fmt.Printf("%s\n", l.String())
}
