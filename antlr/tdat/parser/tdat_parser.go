// Code generated from Tdat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tdat

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TdatParser struct {
	*antlr.BaseParser
}

var tdatParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tdatParserInit() {
	staticData := &tdatParserStaticData
	staticData.literalNames = []string{
		"", "':'", "'{'", "'}'", "'=>'", "';'", "'Each'", "'None'", "'Any'",
		"'|'", "','", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'%'", "'roundUp'",
		"'roundDown'", "'abs'", "'or'", "'and'", "'<'", "'>'", "'<='", "'>='",
		"'=='", "'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "ID", "METRIC", "INT",
		"FLOAT", "STRING", "CHAR", "LINE_COMMENT", "COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"prog", "ruleLine", "ruleID", "enumerableFunc", "windowsRange", "conditionExpr",
		"primaryExpr", "arithmeticOp", "builtin", "logicalOp", "comparisonOp",
		"literal", "result",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 127, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 4, 0, 28, 8, 0, 11, 0, 12, 0, 29,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 51, 8, 4, 1, 4, 1, 4, 3,
		4, 55, 8, 4, 1, 4, 3, 4, 58, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 69, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 75, 8,
		5, 10, 5, 12, 5, 78, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 92, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6,
		98, 8, 6, 10, 6, 12, 6, 101, 9, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 117, 8,
		12, 10, 12, 12, 12, 120, 9, 12, 1, 12, 1, 12, 1, 12, 3, 12, 125, 8, 12,
		1, 12, 0, 2, 10, 12, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		0, 6, 1, 0, 6, 8, 1, 0, 13, 17, 1, 0, 18, 20, 1, 0, 21, 22, 1, 0, 23, 28,
		1, 0, 31, 33, 125, 0, 27, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 42, 1, 0, 0,
		0, 6, 44, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12, 91, 1,
		0, 0, 0, 14, 102, 1, 0, 0, 0, 16, 104, 1, 0, 0, 0, 18, 106, 1, 0, 0, 0,
		20, 108, 1, 0, 0, 0, 22, 110, 1, 0, 0, 0, 24, 124, 1, 0, 0, 0, 26, 28,
		3, 2, 1, 0, 27, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0,
		29, 30, 1, 0, 0, 0, 30, 1, 1, 0, 0, 0, 31, 32, 3, 4, 2, 0, 32, 33, 5, 1,
		0, 0, 33, 34, 3, 6, 3, 0, 34, 35, 5, 2, 0, 0, 35, 36, 3, 8, 4, 0, 36, 37,
		3, 10, 5, 0, 37, 38, 5, 3, 0, 0, 38, 39, 5, 4, 0, 0, 39, 40, 3, 24, 12,
		0, 40, 41, 5, 5, 0, 0, 41, 3, 1, 0, 0, 0, 42, 43, 5, 29, 0, 0, 43, 5, 1,
		0, 0, 0, 44, 45, 7, 0, 0, 0, 45, 7, 1, 0, 0, 0, 46, 47, 5, 9, 0, 0, 47,
		58, 5, 9, 0, 0, 48, 50, 5, 9, 0, 0, 49, 51, 5, 31, 0, 0, 50, 49, 1, 0,
		0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 5, 10, 0, 0, 53,
		55, 5, 31, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0,
		0, 0, 56, 58, 5, 9, 0, 0, 57, 46, 1, 0, 0, 0, 57, 48, 1, 0, 0, 0, 58, 9,
		1, 0, 0, 0, 59, 60, 6, 5, -1, 0, 60, 61, 5, 11, 0, 0, 61, 62, 3, 10, 5,
		0, 62, 63, 5, 12, 0, 0, 63, 69, 1, 0, 0, 0, 64, 65, 3, 12, 6, 0, 65, 66,
		3, 20, 10, 0, 66, 67, 3, 12, 6, 0, 67, 69, 1, 0, 0, 0, 68, 59, 1, 0, 0,
		0, 68, 64, 1, 0, 0, 0, 69, 76, 1, 0, 0, 0, 70, 71, 10, 3, 0, 0, 71, 72,
		3, 18, 9, 0, 72, 73, 3, 10, 5, 4, 73, 75, 1, 0, 0, 0, 74, 70, 1, 0, 0,
		0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 11,
		1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 6, 6, -1, 0, 80, 81, 5, 11, 0,
		0, 81, 82, 3, 12, 6, 0, 82, 83, 5, 12, 0, 0, 83, 92, 1, 0, 0, 0, 84, 92,
		5, 30, 0, 0, 85, 86, 3, 16, 8, 0, 86, 87, 5, 11, 0, 0, 87, 88, 3, 12, 6,
		0, 88, 89, 5, 12, 0, 0, 89, 92, 1, 0, 0, 0, 90, 92, 3, 22, 11, 0, 91, 79,
		1, 0, 0, 0, 91, 84, 1, 0, 0, 0, 91, 85, 1, 0, 0, 0, 91, 90, 1, 0, 0, 0,
		92, 99, 1, 0, 0, 0, 93, 94, 10, 4, 0, 0, 94, 95, 3, 14, 7, 0, 95, 96, 3,
		12, 6, 5, 96, 98, 1, 0, 0, 0, 97, 93, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0,
		99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 13, 1, 0, 0, 0, 101, 99,
		1, 0, 0, 0, 102, 103, 7, 1, 0, 0, 103, 15, 1, 0, 0, 0, 104, 105, 7, 2,
		0, 0, 105, 17, 1, 0, 0, 0, 106, 107, 7, 3, 0, 0, 107, 19, 1, 0, 0, 0, 108,
		109, 7, 4, 0, 0, 109, 21, 1, 0, 0, 0, 110, 111, 7, 5, 0, 0, 111, 23, 1,
		0, 0, 0, 112, 113, 5, 11, 0, 0, 113, 118, 5, 33, 0, 0, 114, 115, 5, 10,
		0, 0, 115, 117, 5, 33, 0, 0, 116, 114, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120,
		118, 1, 0, 0, 0, 121, 125, 5, 12, 0, 0, 122, 123, 5, 11, 0, 0, 123, 125,
		5, 12, 0, 0, 124, 112, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 25, 1, 0,
		0, 0, 10, 29, 50, 54, 57, 68, 76, 91, 99, 118, 124,
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

// TdatParserInit initializes any static state used to implement TdatParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTdatParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TdatParserInit() {
	staticData := &tdatParserStaticData
	staticData.once.Do(tdatParserInit)
}

// NewTdatParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTdatParser(input antlr.TokenStream) *TdatParser {
	TdatParserInit()
	this := new(TdatParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tdatParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Tdat.g4"

	return this
}

// TdatParser tokens.
const (
	TdatParserEOF          = antlr.TokenEOF
	TdatParserT__0         = 1
	TdatParserT__1         = 2
	TdatParserT__2         = 3
	TdatParserT__3         = 4
	TdatParserT__4         = 5
	TdatParserT__5         = 6
	TdatParserT__6         = 7
	TdatParserT__7         = 8
	TdatParserT__8         = 9
	TdatParserT__9         = 10
	TdatParserT__10        = 11
	TdatParserT__11        = 12
	TdatParserT__12        = 13
	TdatParserT__13        = 14
	TdatParserT__14        = 15
	TdatParserT__15        = 16
	TdatParserT__16        = 17
	TdatParserT__17        = 18
	TdatParserT__18        = 19
	TdatParserT__19        = 20
	TdatParserT__20        = 21
	TdatParserT__21        = 22
	TdatParserT__22        = 23
	TdatParserT__23        = 24
	TdatParserT__24        = 25
	TdatParserT__25        = 26
	TdatParserT__26        = 27
	TdatParserT__27        = 28
	TdatParserID           = 29
	TdatParserMETRIC       = 30
	TdatParserINT          = 31
	TdatParserFLOAT        = 32
	TdatParserSTRING       = 33
	TdatParserCHAR         = 34
	TdatParserLINE_COMMENT = 35
	TdatParserCOMMENT      = 36
	TdatParserWS           = 37
)

// TdatParser rules.
const (
	TdatParserRULE_prog           = 0
	TdatParserRULE_ruleLine       = 1
	TdatParserRULE_ruleID         = 2
	TdatParserRULE_enumerableFunc = 3
	TdatParserRULE_windowsRange   = 4
	TdatParserRULE_conditionExpr  = 5
	TdatParserRULE_primaryExpr    = 6
	TdatParserRULE_arithmeticOp   = 7
	TdatParserRULE_builtin        = 8
	TdatParserRULE_logicalOp      = 9
	TdatParserRULE_comparisonOp   = 10
	TdatParserRULE_literal        = 11
	TdatParserRULE_result         = 12
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllRuleLine() []IRuleLineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleLineContext); ok {
			len++
		}
	}

	tst := make([]IRuleLineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleLineContext); ok {
			tst[i] = t.(IRuleLineContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) RuleLine(i int) IRuleLineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleLineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleLineContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *TdatParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TdatParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TdatParserID {
		{
			p.SetState(26)
			p.RuleLine()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRuleLineContext is an interface to support dynamic dispatch.
type IRuleLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleLineContext differentiates from other interfaces.
	IsRuleLineContext()
}

type RuleLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleLineContext() *RuleLineContext {
	var p = new(RuleLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_ruleLine
	return p
}

func (*RuleLineContext) IsRuleLineContext() {}

func NewRuleLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleLineContext {
	var p = new(RuleLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_ruleLine

	return p
}

func (s *RuleLineContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleLineContext) RuleID() IRuleIDContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleIDContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleIDContext)
}

func (s *RuleLineContext) EnumerableFunc() IEnumerableFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumerableFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumerableFuncContext)
}

func (s *RuleLineContext) WindowsRange() IWindowsRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowsRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindowsRangeContext)
}

func (s *RuleLineContext) ConditionExpr() IConditionExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionExprContext)
}

func (s *RuleLineContext) Result() IResultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultContext)
}

func (s *RuleLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterRuleLine(s)
	}
}

func (s *RuleLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitRuleLine(s)
	}
}

func (p *TdatParser) RuleLine() (localctx IRuleLineContext) {
	this := p
	_ = this

	localctx = NewRuleLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TdatParserRULE_ruleLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.RuleID()
	}
	{
		p.SetState(32)
		p.Match(TdatParserT__0)
	}
	{
		p.SetState(33)
		p.EnumerableFunc()
	}
	{
		p.SetState(34)
		p.Match(TdatParserT__1)
	}
	{
		p.SetState(35)
		p.WindowsRange()
	}
	{
		p.SetState(36)
		p.conditionExpr(0)
	}
	{
		p.SetState(37)
		p.Match(TdatParserT__2)
	}
	{
		p.SetState(38)
		p.Match(TdatParserT__3)
	}
	{
		p.SetState(39)
		p.Result()
	}
	{
		p.SetState(40)
		p.Match(TdatParserT__4)
	}

	return localctx
}

// IRuleIDContext is an interface to support dynamic dispatch.
type IRuleIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleIDContext differentiates from other interfaces.
	IsRuleIDContext()
}

type RuleIDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleIDContext() *RuleIDContext {
	var p = new(RuleIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_ruleID
	return p
}

func (*RuleIDContext) IsRuleIDContext() {}

func NewRuleIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleIDContext {
	var p = new(RuleIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_ruleID

	return p
}

func (s *RuleIDContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleIDContext) ID() antlr.TerminalNode {
	return s.GetToken(TdatParserID, 0)
}

func (s *RuleIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterRuleID(s)
	}
}

func (s *RuleIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitRuleID(s)
	}
}

func (p *TdatParser) RuleID() (localctx IRuleIDContext) {
	this := p
	_ = this

	localctx = NewRuleIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TdatParserRULE_ruleID)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(TdatParserID)
	}

	return localctx
}

// IEnumerableFuncContext is an interface to support dynamic dispatch.
type IEnumerableFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumerableFuncContext differentiates from other interfaces.
	IsEnumerableFuncContext()
}

type EnumerableFuncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumerableFuncContext() *EnumerableFuncContext {
	var p = new(EnumerableFuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_enumerableFunc
	return p
}

func (*EnumerableFuncContext) IsEnumerableFuncContext() {}

func NewEnumerableFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumerableFuncContext {
	var p = new(EnumerableFuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_enumerableFunc

	return p
}

func (s *EnumerableFuncContext) GetParser() antlr.Parser { return s.parser }
func (s *EnumerableFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumerableFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumerableFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterEnumerableFunc(s)
	}
}

func (s *EnumerableFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitEnumerableFunc(s)
	}
}

func (p *TdatParser) EnumerableFunc() (localctx IEnumerableFuncContext) {
	this := p
	_ = this

	localctx = NewEnumerableFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TdatParserRULE_enumerableFunc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TdatParserT__5)|(1<<TdatParserT__6)|(1<<TdatParserT__7))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWindowsRangeContext is an interface to support dynamic dispatch.
type IWindowsRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowsRangeContext differentiates from other interfaces.
	IsWindowsRangeContext()
}

type WindowsRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowsRangeContext() *WindowsRangeContext {
	var p = new(WindowsRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_windowsRange
	return p
}

func (*WindowsRangeContext) IsWindowsRangeContext() {}

func NewWindowsRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowsRangeContext {
	var p = new(WindowsRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_windowsRange

	return p
}

func (s *WindowsRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowsRangeContext) CopyFrom(ctx *WindowsRangeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *WindowsRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowsRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WindowsWithZeroIndexContext struct {
	*WindowsRangeContext
}

func NewWindowsWithZeroIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WindowsWithZeroIndexContext {
	var p = new(WindowsWithZeroIndexContext)

	p.WindowsRangeContext = NewEmptyWindowsRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WindowsRangeContext))

	return p
}

func (s *WindowsWithZeroIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowsWithZeroIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterWindowsWithZeroIndex(s)
	}
}

func (s *WindowsWithZeroIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitWindowsWithZeroIndex(s)
	}
}

type WindowsWithLowAndHighIndexContext struct {
	*WindowsRangeContext
}

func NewWindowsWithLowAndHighIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WindowsWithLowAndHighIndexContext {
	var p = new(WindowsWithLowAndHighIndexContext)

	p.WindowsRangeContext = NewEmptyWindowsRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WindowsRangeContext))

	return p
}

func (s *WindowsWithLowAndHighIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowsWithLowAndHighIndexContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(TdatParserINT)
}

func (s *WindowsWithLowAndHighIndexContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(TdatParserINT, i)
}

func (s *WindowsWithLowAndHighIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterWindowsWithLowAndHighIndex(s)
	}
}

func (s *WindowsWithLowAndHighIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitWindowsWithLowAndHighIndex(s)
	}
}

func (p *TdatParser) WindowsRange() (localctx IWindowsRangeContext) {
	this := p
	_ = this

	localctx = NewWindowsRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TdatParserRULE_windowsRange)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewWindowsWithZeroIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(TdatParserT__8)
		}
		{
			p.SetState(47)
			p.Match(TdatParserT__8)
		}

	case 2:
		localctx = NewWindowsWithLowAndHighIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(TdatParserT__8)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TdatParserINT {
			{
				p.SetState(49)
				p.Match(TdatParserINT)
			}

		}
		{
			p.SetState(52)
			p.Match(TdatParserT__9)
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TdatParserINT {
			{
				p.SetState(53)
				p.Match(TdatParserINT)
			}

		}
		{
			p.SetState(56)
			p.Match(TdatParserT__8)
		}

	}

	return localctx
}

// IConditionExprContext is an interface to support dynamic dispatch.
type IConditionExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionExprContext differentiates from other interfaces.
	IsConditionExprContext()
}

type ConditionExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionExprContext() *ConditionExprContext {
	var p = new(ConditionExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_conditionExpr
	return p
}

func (*ConditionExprContext) IsConditionExprContext() {}

func NewConditionExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionExprContext {
	var p = new(ConditionExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_conditionExpr

	return p
}

func (s *ConditionExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionExprContext) AllConditionExpr() []IConditionExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionExprContext); ok {
			len++
		}
	}

	tst := make([]IConditionExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionExprContext); ok {
			tst[i] = t.(IConditionExprContext)
			i++
		}
	}

	return tst
}

func (s *ConditionExprContext) ConditionExpr(i int) IConditionExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionExprContext)
}

func (s *ConditionExprContext) AllPrimaryExpr() []IPrimaryExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryExprContext); ok {
			tst[i] = t.(IPrimaryExprContext)
			i++
		}
	}

	return tst
}

func (s *ConditionExprContext) PrimaryExpr(i int) IPrimaryExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ConditionExprContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *ConditionExprContext) LogicalOp() ILogicalOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOpContext)
}

func (s *ConditionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterConditionExpr(s)
	}
}

func (s *ConditionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitConditionExpr(s)
	}
}

func (p *TdatParser) ConditionExpr() (localctx IConditionExprContext) {
	return p.conditionExpr(0)
}

func (p *TdatParser) conditionExpr(_p int) (localctx IConditionExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewConditionExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, TdatParserRULE_conditionExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(60)
			p.Match(TdatParserT__10)
		}
		{
			p.SetState(61)
			p.conditionExpr(0)
		}
		{
			p.SetState(62)
			p.Match(TdatParserT__11)
		}

	case 2:
		{
			p.SetState(64)
			p.primaryExpr(0)
		}
		{
			p.SetState(65)
			p.ComparisonOp()
		}
		{
			p.SetState(66)
			p.primaryExpr(0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConditionExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TdatParserRULE_conditionExpr)
			p.SetState(70)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(71)
				p.LogicalOp()
			}
			{
				p.SetState(72)
				p.conditionExpr(4)
			}

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) CopyFrom(ctx *PrimaryExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BracketExprInPrimaryExprContext struct {
	*PrimaryExprContext
}

func NewBracketExprInPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketExprInPrimaryExprContext {
	var p = new(BracketExprInPrimaryExprContext)

	p.PrimaryExprContext = NewEmptyPrimaryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExprContext))

	return p
}

func (s *BracketExprInPrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExprInPrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *BracketExprInPrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterBracketExprInPrimaryExpr(s)
	}
}

func (s *BracketExprInPrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitBracketExprInPrimaryExpr(s)
	}
}

type RightLiteralInPrimaryExprContext struct {
	*PrimaryExprContext
}

func NewRightLiteralInPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RightLiteralInPrimaryExprContext {
	var p = new(RightLiteralInPrimaryExprContext)

	p.PrimaryExprContext = NewEmptyPrimaryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExprContext))

	return p
}

func (s *RightLiteralInPrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightLiteralInPrimaryExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *RightLiteralInPrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterRightLiteralInPrimaryExpr(s)
	}
}

func (s *RightLiteralInPrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitRightLiteralInPrimaryExpr(s)
	}
}

type ArithmeticExprInPrimaryExprContext struct {
	*PrimaryExprContext
}

func NewArithmeticExprInPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticExprInPrimaryExprContext {
	var p = new(ArithmeticExprInPrimaryExprContext)

	p.PrimaryExprContext = NewEmptyPrimaryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExprContext))

	return p
}

func (s *ArithmeticExprInPrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticExprInPrimaryExprContext) AllPrimaryExpr() []IPrimaryExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryExprContext); ok {
			tst[i] = t.(IPrimaryExprContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticExprInPrimaryExprContext) PrimaryExpr(i int) IPrimaryExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ArithmeticExprInPrimaryExprContext) ArithmeticOp() IArithmeticOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticOpContext)
}

func (s *ArithmeticExprInPrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterArithmeticExprInPrimaryExpr(s)
	}
}

func (s *ArithmeticExprInPrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitArithmeticExprInPrimaryExpr(s)
	}
}

type BuildinExprInPrimaryExprContext struct {
	*PrimaryExprContext
}

func NewBuildinExprInPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuildinExprInPrimaryExprContext {
	var p = new(BuildinExprInPrimaryExprContext)

	p.PrimaryExprContext = NewEmptyPrimaryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExprContext))

	return p
}

func (s *BuildinExprInPrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuildinExprInPrimaryExprContext) Builtin() IBuiltinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltinContext)
}

func (s *BuildinExprInPrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *BuildinExprInPrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterBuildinExprInPrimaryExpr(s)
	}
}

func (s *BuildinExprInPrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitBuildinExprInPrimaryExpr(s)
	}
}

type MetricInPrimaryExprContext struct {
	*PrimaryExprContext
}

func NewMetricInPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MetricInPrimaryExprContext {
	var p = new(MetricInPrimaryExprContext)

	p.PrimaryExprContext = NewEmptyPrimaryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExprContext))

	return p
}

func (s *MetricInPrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetricInPrimaryExprContext) METRIC() antlr.TerminalNode {
	return s.GetToken(TdatParserMETRIC, 0)
}

func (s *MetricInPrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterMetricInPrimaryExpr(s)
	}
}

func (s *MetricInPrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitMetricInPrimaryExpr(s)
	}
}

func (p *TdatParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *TdatParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, TdatParserRULE_primaryExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TdatParserT__10:
		localctx = NewBracketExprInPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(80)
			p.Match(TdatParserT__10)
		}
		{
			p.SetState(81)
			p.primaryExpr(0)
		}
		{
			p.SetState(82)
			p.Match(TdatParserT__11)
		}

	case TdatParserMETRIC:
		localctx = NewMetricInPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Match(TdatParserMETRIC)
		}

	case TdatParserT__17, TdatParserT__18, TdatParserT__19:
		localctx = NewBuildinExprInPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(85)
			p.Builtin()
		}
		{
			p.SetState(86)
			p.Match(TdatParserT__10)
		}
		{
			p.SetState(87)
			p.primaryExpr(0)
		}
		{
			p.SetState(88)
			p.Match(TdatParserT__11)
		}

	case TdatParserINT, TdatParserFLOAT, TdatParserSTRING:
		localctx = NewRightLiteralInPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArithmeticExprInPrimaryExprContext(p, NewPrimaryExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, TdatParserRULE_primaryExpr)
			p.SetState(93)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(94)
				p.ArithmeticOp()
			}
			{
				p.SetState(95)
				p.primaryExpr(5)
			}

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IArithmeticOpContext is an interface to support dynamic dispatch.
type IArithmeticOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithmeticOpContext differentiates from other interfaces.
	IsArithmeticOpContext()
}

type ArithmeticOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmeticOpContext() *ArithmeticOpContext {
	var p = new(ArithmeticOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_arithmeticOp
	return p
}

func (*ArithmeticOpContext) IsArithmeticOpContext() {}

func NewArithmeticOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticOpContext {
	var p = new(ArithmeticOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_arithmeticOp

	return p
}

func (s *ArithmeticOpContext) GetParser() antlr.Parser { return s.parser }
func (s *ArithmeticOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterArithmeticOp(s)
	}
}

func (s *ArithmeticOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitArithmeticOp(s)
	}
}

func (p *TdatParser) ArithmeticOp() (localctx IArithmeticOpContext) {
	this := p
	_ = this

	localctx = NewArithmeticOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TdatParserRULE_arithmeticOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TdatParserT__12)|(1<<TdatParserT__13)|(1<<TdatParserT__14)|(1<<TdatParserT__15)|(1<<TdatParserT__16))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBuiltinContext is an interface to support dynamic dispatch.
type IBuiltinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinContext differentiates from other interfaces.
	IsBuiltinContext()
}

type BuiltinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinContext() *BuiltinContext {
	var p = new(BuiltinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_builtin
	return p
}

func (*BuiltinContext) IsBuiltinContext() {}

func NewBuiltinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinContext {
	var p = new(BuiltinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_builtin

	return p
}

func (s *BuiltinContext) GetParser() antlr.Parser { return s.parser }
func (s *BuiltinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterBuiltin(s)
	}
}

func (s *BuiltinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitBuiltin(s)
	}
}

func (p *TdatParser) Builtin() (localctx IBuiltinContext) {
	this := p
	_ = this

	localctx = NewBuiltinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TdatParserRULE_builtin)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TdatParserT__17)|(1<<TdatParserT__18)|(1<<TdatParserT__19))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogicalOpContext is an interface to support dynamic dispatch.
type ILogicalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOpContext differentiates from other interfaces.
	IsLogicalOpContext()
}

type LogicalOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOpContext() *LogicalOpContext {
	var p = new(LogicalOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_logicalOp
	return p
}

func (*LogicalOpContext) IsLogicalOpContext() {}

func NewLogicalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOpContext {
	var p = new(LogicalOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_logicalOp

	return p
}

func (s *LogicalOpContext) GetParser() antlr.Parser { return s.parser }
func (s *LogicalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterLogicalOp(s)
	}
}

func (s *LogicalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitLogicalOp(s)
	}
}

func (p *TdatParser) LogicalOp() (localctx ILogicalOpContext) {
	this := p
	_ = this

	localctx = NewLogicalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TdatParserRULE_logicalOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TdatParserT__20 || _la == TdatParserT__21) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonOpContext is an interface to support dynamic dispatch.
type IComparisonOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOpContext differentiates from other interfaces.
	IsComparisonOpContext()
}

type ComparisonOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOpContext() *ComparisonOpContext {
	var p = new(ComparisonOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_comparisonOp
	return p
}

func (*ComparisonOpContext) IsComparisonOpContext() {}

func NewComparisonOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOpContext {
	var p = new(ComparisonOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_comparisonOp

	return p
}

func (s *ComparisonOpContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterComparisonOp(s)
	}
}

func (s *ComparisonOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitComparisonOp(s)
	}
}

func (p *TdatParser) ComparisonOp() (localctx IComparisonOpContext) {
	this := p
	_ = this

	localctx = NewComparisonOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TdatParserRULE_comparisonOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TdatParserT__22)|(1<<TdatParserT__23)|(1<<TdatParserT__24)|(1<<TdatParserT__25)|(1<<TdatParserT__26)|(1<<TdatParserT__27))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(TdatParserINT, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TdatParserFLOAT, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(TdatParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *TdatParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TdatParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(TdatParserINT-31))|(1<<(TdatParserFLOAT-31))|(1<<(TdatParserSTRING-31)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IResultContext is an interface to support dynamic dispatch.
type IResultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResultContext differentiates from other interfaces.
	IsResultContext()
}

type ResultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultContext() *ResultContext {
	var p = new(ResultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TdatParserRULE_result
	return p
}

func (*ResultContext) IsResultContext() {}

func NewResultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultContext {
	var p = new(ResultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TdatParserRULE_result

	return p
}

func (s *ResultContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultContext) CopyFrom(ctx *ResultContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ResultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ResultWithElementsContext struct {
	*ResultContext
}

func NewResultWithElementsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ResultWithElementsContext {
	var p = new(ResultWithElementsContext)

	p.ResultContext = NewEmptyResultContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ResultContext))

	return p
}

func (s *ResultWithElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultWithElementsContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(TdatParserSTRING)
}

func (s *ResultWithElementsContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(TdatParserSTRING, i)
}

func (s *ResultWithElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterResultWithElements(s)
	}
}

func (s *ResultWithElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitResultWithElements(s)
	}
}

type ResultWithoutElementsContext struct {
	*ResultContext
}

func NewResultWithoutElementsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ResultWithoutElementsContext {
	var p = new(ResultWithoutElementsContext)

	p.ResultContext = NewEmptyResultContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ResultContext))

	return p
}

func (s *ResultWithoutElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultWithoutElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.EnterResultWithoutElements(s)
	}
}

func (s *ResultWithoutElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TdatListener); ok {
		listenerT.ExitResultWithoutElements(s)
	}
}

func (p *TdatParser) Result() (localctx IResultContext) {
	this := p
	_ = this

	localctx = NewResultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TdatParserRULE_result)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewResultWithElementsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(TdatParserT__10)
		}
		{
			p.SetState(113)
			p.Match(TdatParserSTRING)
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TdatParserT__9 {
			{
				p.SetState(114)
				p.Match(TdatParserT__9)
			}
			{
				p.SetState(115)
				p.Match(TdatParserSTRING)
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(121)
			p.Match(TdatParserT__11)
		}

	case 2:
		localctx = NewResultWithoutElementsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(TdatParserT__10)
		}
		{
			p.SetState(123)
			p.Match(TdatParserT__11)
		}

	}

	return localctx
}

func (p *TdatParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ConditionExprContext = nil
		if localctx != nil {
			t = localctx.(*ConditionExprContext)
		}
		return p.ConditionExpr_Sempred(t, predIndex)

	case 6:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TdatParser) ConditionExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TdatParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
