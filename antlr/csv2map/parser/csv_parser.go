// Code generated from CSV.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CSV

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

type CSVParser struct {
	*antlr.BaseParser
}

var csvParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func csvParserInit() {
	staticData := &csvParserStaticData
	staticData.literalNames = []string{
		"", "','", "'\\r'", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "TEXT", "STRING",
	}
	staticData.ruleNames = []string{
		"csvFile", "hdr", "row", "field",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 35, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1, 0,
		4, 0, 11, 8, 0, 11, 0, 12, 0, 12, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 20,
		8, 2, 10, 2, 12, 2, 23, 9, 2, 1, 2, 3, 2, 26, 8, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 3, 3, 33, 8, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6, 0, 0, 35, 0, 8, 1,
		0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 16, 1, 0, 0, 0, 6, 32, 1, 0, 0, 0, 8, 10,
		3, 2, 1, 0, 9, 11, 3, 4, 2, 0, 10, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12,
		10, 1, 0, 0, 0, 12, 13, 1, 0, 0, 0, 13, 1, 1, 0, 0, 0, 14, 15, 3, 4, 2,
		0, 15, 3, 1, 0, 0, 0, 16, 21, 3, 6, 3, 0, 17, 18, 5, 1, 0, 0, 18, 20, 3,
		6, 3, 0, 19, 17, 1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21,
		22, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 24, 26, 5, 2, 0,
		0, 25, 24, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28,
		5, 3, 0, 0, 28, 5, 1, 0, 0, 0, 29, 33, 5, 4, 0, 0, 30, 33, 5, 5, 0, 0,
		31, 33, 1, 0, 0, 0, 32, 29, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 31, 1,
		0, 0, 0, 33, 7, 1, 0, 0, 0, 4, 12, 21, 25, 32,
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

// CSVParserInit initializes any static state used to implement CSVParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCSVParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CSVParserInit() {
	staticData := &csvParserStaticData
	staticData.once.Do(csvParserInit)
}

// NewCSVParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCSVParser(input antlr.TokenStream) *CSVParser {
	CSVParserInit()
	this := new(CSVParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &csvParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CSV.g4"

	return this
}

// CSVParser tokens.
const (
	CSVParserEOF    = antlr.TokenEOF
	CSVParserT__0   = 1
	CSVParserT__1   = 2
	CSVParserT__2   = 3
	CSVParserTEXT   = 4
	CSVParserSTRING = 5
)

// CSVParser rules.
const (
	CSVParserRULE_csvFile = 0
	CSVParserRULE_hdr     = 1
	CSVParserRULE_row     = 2
	CSVParserRULE_field   = 3
)

// ICsvFileContext is an interface to support dynamic dispatch.
type ICsvFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCsvFileContext differentiates from other interfaces.
	IsCsvFileContext()
}

type CsvFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCsvFileContext() *CsvFileContext {
	var p = new(CsvFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_csvFile
	return p
}

func (*CsvFileContext) IsCsvFileContext() {}

func NewCsvFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CsvFileContext {
	var p = new(CsvFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_csvFile

	return p
}

func (s *CsvFileContext) GetParser() antlr.Parser { return s.parser }

func (s *CsvFileContext) Hdr() IHdrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHdrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHdrContext)
}

func (s *CsvFileContext) AllRow() []IRowContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRowContext); ok {
			len++
		}
	}

	tst := make([]IRowContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRowContext); ok {
			tst[i] = t.(IRowContext)
			i++
		}
	}

	return tst
}

func (s *CsvFileContext) Row(i int) IRowContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRowContext); ok {
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

	return t.(IRowContext)
}

func (s *CsvFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CsvFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CsvFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterCsvFile(s)
	}
}

func (s *CsvFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitCsvFile(s)
	}
}

func (p *CSVParser) CsvFile() (localctx ICsvFileContext) {
	this := p
	_ = this

	localctx = NewCsvFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CSVParserRULE_csvFile)
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
		p.SetState(8)
		p.Hdr()
	}
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CSVParserT__0)|(1<<CSVParserT__1)|(1<<CSVParserT__2)|(1<<CSVParserTEXT)|(1<<CSVParserSTRING))) != 0) {
		{
			p.SetState(9)
			p.Row()
		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHdrContext is an interface to support dynamic dispatch.
type IHdrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHdrContext differentiates from other interfaces.
	IsHdrContext()
}

type HdrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHdrContext() *HdrContext {
	var p = new(HdrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_hdr
	return p
}

func (*HdrContext) IsHdrContext() {}

func NewHdrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HdrContext {
	var p = new(HdrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_hdr

	return p
}

func (s *HdrContext) GetParser() antlr.Parser { return s.parser }

func (s *HdrContext) Row() IRowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *HdrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HdrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HdrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterHdr(s)
	}
}

func (s *HdrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitHdr(s)
	}
}

func (p *CSVParser) Hdr() (localctx IHdrContext) {
	this := p
	_ = this

	localctx = NewHdrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CSVParserRULE_hdr)

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
		p.SetState(14)
		p.Row()
	}

	return localctx
}

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *RowContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterRow(s)
	}
}

func (s *RowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitRow(s)
	}
}

func (p *CSVParser) Row() (localctx IRowContext) {
	this := p
	_ = this

	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CSVParserRULE_row)
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
		p.SetState(16)
		p.Field()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CSVParserT__0 {
		{
			p.SetState(17)
			p.Match(CSVParserT__0)
		}
		{
			p.SetState(18)
			p.Field()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CSVParserT__1 {
		{
			p.SetState(24)
			p.Match(CSVParserT__1)
		}

	}
	{
		p.SetState(27)
		p.Match(CSVParserT__2)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) TEXT() antlr.TerminalNode {
	return s.GetToken(CSVParserTEXT, 0)
}

func (s *FieldContext) STRING() antlr.TerminalNode {
	return s.GetToken(CSVParserSTRING, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *CSVParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CSVParserRULE_field)

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

	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CSVParserTEXT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(CSVParserTEXT)
		}

	case CSVParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(CSVParserSTRING)
		}

	case CSVParserT__0, CSVParserT__1, CSVParserT__2:
		p.EnterOuterAlt(localctx, 3)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
