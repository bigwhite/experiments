// Code generated from CSV.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CSVLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var csvlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func csvlexerLexerInit() {
	staticData := &csvlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "','", "'\\r'", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "TEXT", "STRING",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "TEXT", "STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 33, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 4, 3, 19, 8, 3, 11,
		3, 12, 3, 20, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 27, 8, 4, 10, 4, 12, 4, 30,
		9, 4, 1, 4, 1, 4, 0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0, 2, 4, 0,
		10, 10, 13, 13, 34, 34, 44, 44, 1, 0, 34, 34, 35, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1,
		11, 1, 0, 0, 0, 3, 13, 1, 0, 0, 0, 5, 15, 1, 0, 0, 0, 7, 18, 1, 0, 0, 0,
		9, 22, 1, 0, 0, 0, 11, 12, 5, 44, 0, 0, 12, 2, 1, 0, 0, 0, 13, 14, 5, 13,
		0, 0, 14, 4, 1, 0, 0, 0, 15, 16, 5, 10, 0, 0, 16, 6, 1, 0, 0, 0, 17, 19,
		8, 0, 0, 0, 18, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0,
		20, 21, 1, 0, 0, 0, 21, 8, 1, 0, 0, 0, 22, 28, 5, 34, 0, 0, 23, 24, 5,
		34, 0, 0, 24, 27, 5, 34, 0, 0, 25, 27, 8, 1, 0, 0, 26, 23, 1, 0, 0, 0,
		26, 25, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1,
		0, 0, 0, 29, 31, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 5, 34, 0, 0, 32,
		10, 1, 0, 0, 0, 4, 0, 20, 26, 28, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CSVLexerInit initializes any static state used to implement CSVLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCSVLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CSVLexerInit() {
	staticData := &csvlexerLexerStaticData
	staticData.once.Do(csvlexerLexerInit)
}

// NewCSVLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCSVLexer(input antlr.CharStream) *CSVLexer {
	CSVLexerInit()
	l := new(CSVLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &csvlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "CSV.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CSVLexer tokens.
const (
	CSVLexerT__0   = 1
	CSVLexerT__1   = 2
	CSVLexerT__2   = 3
	CSVLexerTEXT   = 4
	CSVLexerSTRING = 5
)
