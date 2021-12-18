// Code generated from GrammarDesc.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // GrammarDesc

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 71, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 23, 10, 2, 12, 2, 14,
	2, 26, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 48,
	10, 6, 12, 6, 14, 6, 51, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 5, 7, 57, 10, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 63, 10, 8, 3, 8, 3, 8, 5, 8, 67, 10, 8, 3,
	9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 2, 2, 72, 2, 24,
	3, 2, 2, 2, 4, 27, 3, 2, 2, 2, 6, 32, 3, 2, 2, 2, 8, 37, 3, 2, 2, 2, 10,
	42, 3, 2, 2, 2, 12, 56, 3, 2, 2, 2, 14, 62, 3, 2, 2, 2, 16, 68, 3, 2, 2,
	2, 18, 23, 5, 4, 3, 2, 19, 23, 5, 6, 4, 2, 20, 23, 5, 8, 5, 2, 21, 23,
	5, 10, 6, 2, 22, 18, 3, 2, 2, 2, 22, 19, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2,
	22, 21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3,
	2, 2, 2, 25, 3, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27, 28, 7, 10, 2, 2, 28,
	29, 7, 3, 2, 2, 29, 30, 7, 12, 2, 2, 30, 31, 7, 4, 2, 2, 31, 5, 3, 2, 2,
	2, 32, 33, 7, 5, 2, 2, 33, 34, 7, 3, 2, 2, 34, 35, 7, 12, 2, 2, 35, 36,
	7, 4, 2, 2, 36, 7, 3, 2, 2, 2, 37, 38, 7, 6, 2, 2, 38, 39, 7, 3, 2, 2,
	39, 40, 7, 11, 2, 2, 40, 41, 7, 4, 2, 2, 41, 9, 3, 2, 2, 2, 42, 43, 7,
	11, 2, 2, 43, 49, 7, 7, 2, 2, 44, 48, 5, 12, 7, 2, 45, 48, 5, 14, 8, 2,
	46, 48, 5, 16, 9, 2, 47, 44, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3,
	2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50,
	52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 7, 4, 2, 2, 53, 11, 3, 2, 2,
	2, 54, 55, 7, 9, 2, 2, 55, 57, 7, 8, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57,
	3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 7, 10, 2, 2, 59, 13, 3, 2, 2, 2,
	60, 61, 7, 9, 2, 2, 61, 63, 7, 8, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3,
	2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 7, 11, 2, 2, 65, 67, 7, 13, 2, 2,
	66, 65, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 15, 3, 2, 2, 2, 68, 69, 7,
	14, 2, 2, 69, 17, 3, 2, 2, 2, 9, 22, 24, 47, 49, 56, 62, 66,
}
var literalNames = []string{
	"", "':'", "';'", "'skip'", "'start'", "'->'", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "ID", "TERM_ID", "NON_TERM_ID", "REGEXP", "ATTR",
	"CODE", "WS",
}

var ruleNames = []string{
	"rules", "lexer_rule", "skip_rule", "start_rule", "parser_rule", "term_parser_rule_item",
	"non_term_parser_rule_item", "code_item",
}

type GrammarDescParser struct {
	*antlr.BaseParser
}

// NewGrammarDescParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GrammarDescParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGrammarDescParser(input antlr.TokenStream) *GrammarDescParser {
	this := new(GrammarDescParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GrammarDesc.g4"

	return this
}

// GrammarDescParser tokens.
const (
	GrammarDescParserEOF         = antlr.TokenEOF
	GrammarDescParserT__0        = 1
	GrammarDescParserT__1        = 2
	GrammarDescParserT__2        = 3
	GrammarDescParserT__3        = 4
	GrammarDescParserT__4        = 5
	GrammarDescParserT__5        = 6
	GrammarDescParserID          = 7
	GrammarDescParserTERM_ID     = 8
	GrammarDescParserNON_TERM_ID = 9
	GrammarDescParserREGEXP      = 10
	GrammarDescParserATTR        = 11
	GrammarDescParserCODE        = 12
	GrammarDescParserWS          = 13
)

// GrammarDescParser rules.
const (
	GrammarDescParserRULE_rules                     = 0
	GrammarDescParserRULE_lexer_rule                = 1
	GrammarDescParserRULE_skip_rule                 = 2
	GrammarDescParserRULE_start_rule                = 3
	GrammarDescParserRULE_parser_rule               = 4
	GrammarDescParserRULE_term_parser_rule_item     = 5
	GrammarDescParserRULE_non_term_parser_rule_item = 6
	GrammarDescParserRULE_code_item                 = 7
)

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarDescParserRULE_rules
	return p
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarDescParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) AllLexer_rule() []ILexer_ruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILexer_ruleContext)(nil)).Elem())
	var tst = make([]ILexer_ruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILexer_ruleContext)
		}
	}

	return tst
}

func (s *RulesContext) Lexer_rule(i int) ILexer_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILexer_ruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILexer_ruleContext)
}

func (s *RulesContext) AllSkip_rule() []ISkip_ruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISkip_ruleContext)(nil)).Elem())
	var tst = make([]ISkip_ruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISkip_ruleContext)
		}
	}

	return tst
}

func (s *RulesContext) Skip_rule(i int) ISkip_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISkip_ruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISkip_ruleContext)
}

func (s *RulesContext) AllStart_rule() []IStart_ruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStart_ruleContext)(nil)).Elem())
	var tst = make([]IStart_ruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStart_ruleContext)
		}
	}

	return tst
}

func (s *RulesContext) Start_rule(i int) IStart_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStart_ruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStart_ruleContext)
}

func (s *RulesContext) AllParser_rule() []IParser_ruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParser_ruleContext)(nil)).Elem())
	var tst = make([]IParser_ruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParser_ruleContext)
		}
	}

	return tst
}

func (s *RulesContext) Parser_rule(i int) IParser_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParser_ruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParser_ruleContext)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.ExitRules(s)
	}
}

func (s *RulesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarDescVisitor:
		return t.VisitRules(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarDescParser) Rules() (localctx IRulesContext) {
	localctx = NewRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarDescParserRULE_rules)
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
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarDescParserT__2)|(1<<GrammarDescParserT__3)|(1<<GrammarDescParserTERM_ID)|(1<<GrammarDescParserNON_TERM_ID))) != 0 {
		p.SetState(20)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GrammarDescParserTERM_ID:
			{
				p.SetState(16)
				p.Lexer_rule()
			}

		case GrammarDescParserT__2:
			{
				p.SetState(17)
				p.Skip_rule()
			}

		case GrammarDescParserT__3:
			{
				p.SetState(18)
				p.Start_rule()
			}

		case GrammarDescParserNON_TERM_ID:
			{
				p.SetState(19)
				p.Parser_rule()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILexer_ruleContext is an interface to support dynamic dispatch.
type ILexer_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLexer_ruleContext differentiates from other interfaces.
	IsLexer_ruleContext()
}

type Lexer_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLexer_ruleContext() *Lexer_ruleContext {
	var p = new(Lexer_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarDescParserRULE_lexer_rule
	return p
}

func (*Lexer_ruleContext) IsLexer_ruleContext() {}

func NewLexer_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lexer_ruleContext {
	var p = new(Lexer_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarDescParserRULE_lexer_rule

	return p
}

func (s *Lexer_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Lexer_ruleContext) TERM_ID() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserTERM_ID, 0)
}

func (s *Lexer_ruleContext) REGEXP() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserREGEXP, 0)
}

func (s *Lexer_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lexer_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lexer_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.EnterLexer_rule(s)
	}
}

func (s *Lexer_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.ExitLexer_rule(s)
	}
}

func (s *Lexer_ruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarDescVisitor:
		return t.VisitLexer_rule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarDescParser) Lexer_rule() (localctx ILexer_ruleContext) {
	localctx = NewLexer_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarDescParserRULE_lexer_rule)

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
		p.SetState(25)
		p.Match(GrammarDescParserTERM_ID)
	}
	{
		p.SetState(26)
		p.Match(GrammarDescParserT__0)
	}
	{
		p.SetState(27)
		p.Match(GrammarDescParserREGEXP)
	}
	{
		p.SetState(28)
		p.Match(GrammarDescParserT__1)
	}

	return localctx
}

// ISkip_ruleContext is an interface to support dynamic dispatch.
type ISkip_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSkip_ruleContext differentiates from other interfaces.
	IsSkip_ruleContext()
}

type Skip_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkip_ruleContext() *Skip_ruleContext {
	var p = new(Skip_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarDescParserRULE_skip_rule
	return p
}

func (*Skip_ruleContext) IsSkip_ruleContext() {}

func NewSkip_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Skip_ruleContext {
	var p = new(Skip_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarDescParserRULE_skip_rule

	return p
}

func (s *Skip_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Skip_ruleContext) REGEXP() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserREGEXP, 0)
}

func (s *Skip_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Skip_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Skip_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.EnterSkip_rule(s)
	}
}

func (s *Skip_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.ExitSkip_rule(s)
	}
}

func (s *Skip_ruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarDescVisitor:
		return t.VisitSkip_rule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarDescParser) Skip_rule() (localctx ISkip_ruleContext) {
	localctx = NewSkip_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarDescParserRULE_skip_rule)

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
		p.SetState(30)
		p.Match(GrammarDescParserT__2)
	}
	{
		p.SetState(31)
		p.Match(GrammarDescParserT__0)
	}
	{
		p.SetState(32)
		p.Match(GrammarDescParserREGEXP)
	}
	{
		p.SetState(33)
		p.Match(GrammarDescParserT__1)
	}

	return localctx
}

// IStart_ruleContext is an interface to support dynamic dispatch.
type IStart_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStart_ruleContext differentiates from other interfaces.
	IsStart_ruleContext()
}

type Start_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStart_ruleContext() *Start_ruleContext {
	var p = new(Start_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarDescParserRULE_start_rule
	return p
}

func (*Start_ruleContext) IsStart_ruleContext() {}

func NewStart_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Start_ruleContext {
	var p = new(Start_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarDescParserRULE_start_rule

	return p
}

func (s *Start_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Start_ruleContext) NON_TERM_ID() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserNON_TERM_ID, 0)
}

func (s *Start_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Start_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Start_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.EnterStart_rule(s)
	}
}

func (s *Start_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.ExitStart_rule(s)
	}
}

func (s *Start_ruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarDescVisitor:
		return t.VisitStart_rule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarDescParser) Start_rule() (localctx IStart_ruleContext) {
	localctx = NewStart_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarDescParserRULE_start_rule)

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
		p.SetState(35)
		p.Match(GrammarDescParserT__3)
	}
	{
		p.SetState(36)
		p.Match(GrammarDescParserT__0)
	}
	{
		p.SetState(37)
		p.Match(GrammarDescParserNON_TERM_ID)
	}
	{
		p.SetState(38)
		p.Match(GrammarDescParserT__1)
	}

	return localctx
}

// IParser_ruleContext is an interface to support dynamic dispatch.
type IParser_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParser_ruleContext differentiates from other interfaces.
	IsParser_ruleContext()
}

type Parser_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParser_ruleContext() *Parser_ruleContext {
	var p = new(Parser_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarDescParserRULE_parser_rule
	return p
}

func (*Parser_ruleContext) IsParser_ruleContext() {}

func NewParser_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parser_ruleContext {
	var p = new(Parser_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarDescParserRULE_parser_rule

	return p
}

func (s *Parser_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Parser_ruleContext) NON_TERM_ID() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserNON_TERM_ID, 0)
}

func (s *Parser_ruleContext) AllTerm_parser_rule_item() []ITerm_parser_rule_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITerm_parser_rule_itemContext)(nil)).Elem())
	var tst = make([]ITerm_parser_rule_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITerm_parser_rule_itemContext)
		}
	}

	return tst
}

func (s *Parser_ruleContext) Term_parser_rule_item(i int) ITerm_parser_rule_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm_parser_rule_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITerm_parser_rule_itemContext)
}

func (s *Parser_ruleContext) AllNon_term_parser_rule_item() []INon_term_parser_rule_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INon_term_parser_rule_itemContext)(nil)).Elem())
	var tst = make([]INon_term_parser_rule_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INon_term_parser_rule_itemContext)
		}
	}

	return tst
}

func (s *Parser_ruleContext) Non_term_parser_rule_item(i int) INon_term_parser_rule_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INon_term_parser_rule_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INon_term_parser_rule_itemContext)
}

func (s *Parser_ruleContext) AllCode_item() []ICode_itemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICode_itemContext)(nil)).Elem())
	var tst = make([]ICode_itemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICode_itemContext)
		}
	}

	return tst
}

func (s *Parser_ruleContext) Code_item(i int) ICode_itemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICode_itemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICode_itemContext)
}

func (s *Parser_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parser_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parser_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.EnterParser_rule(s)
	}
}

func (s *Parser_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.ExitParser_rule(s)
	}
}

func (s *Parser_ruleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarDescVisitor:
		return t.VisitParser_rule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarDescParser) Parser_rule() (localctx IParser_ruleContext) {
	localctx = NewParser_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarDescParserRULE_parser_rule)
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
		p.SetState(40)
		p.Match(GrammarDescParserNON_TERM_ID)
	}
	{
		p.SetState(41)
		p.Match(GrammarDescParserT__4)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarDescParserID)|(1<<GrammarDescParserTERM_ID)|(1<<GrammarDescParserNON_TERM_ID)|(1<<GrammarDescParserCODE))) != 0 {
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(42)
				p.Term_parser_rule_item()
			}

		case 2:
			{
				p.SetState(43)
				p.Non_term_parser_rule_item()
			}

		case 3:
			{
				p.SetState(44)
				p.Code_item()
			}

		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(GrammarDescParserT__1)
	}

	return localctx
}

// ITerm_parser_rule_itemContext is an interface to support dynamic dispatch.
type ITerm_parser_rule_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm_parser_rule_itemContext differentiates from other interfaces.
	IsTerm_parser_rule_itemContext()
}

type Term_parser_rule_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm_parser_rule_itemContext() *Term_parser_rule_itemContext {
	var p = new(Term_parser_rule_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarDescParserRULE_term_parser_rule_item
	return p
}

func (*Term_parser_rule_itemContext) IsTerm_parser_rule_itemContext() {}

func NewTerm_parser_rule_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term_parser_rule_itemContext {
	var p = new(Term_parser_rule_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarDescParserRULE_term_parser_rule_item

	return p
}

func (s *Term_parser_rule_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Term_parser_rule_itemContext) TERM_ID() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserTERM_ID, 0)
}

func (s *Term_parser_rule_itemContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserID, 0)
}

func (s *Term_parser_rule_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term_parser_rule_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term_parser_rule_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.EnterTerm_parser_rule_item(s)
	}
}

func (s *Term_parser_rule_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.ExitTerm_parser_rule_item(s)
	}
}

func (s *Term_parser_rule_itemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarDescVisitor:
		return t.VisitTerm_parser_rule_item(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarDescParser) Term_parser_rule_item() (localctx ITerm_parser_rule_itemContext) {
	localctx = NewTerm_parser_rule_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarDescParserRULE_term_parser_rule_item)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarDescParserID {
		{
			p.SetState(52)
			p.Match(GrammarDescParserID)
		}
		{
			p.SetState(53)
			p.Match(GrammarDescParserT__5)
		}

	}
	{
		p.SetState(56)
		p.Match(GrammarDescParserTERM_ID)
	}

	return localctx
}

// INon_term_parser_rule_itemContext is an interface to support dynamic dispatch.
type INon_term_parser_rule_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNon_term_parser_rule_itemContext differentiates from other interfaces.
	IsNon_term_parser_rule_itemContext()
}

type Non_term_parser_rule_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNon_term_parser_rule_itemContext() *Non_term_parser_rule_itemContext {
	var p = new(Non_term_parser_rule_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarDescParserRULE_non_term_parser_rule_item
	return p
}

func (*Non_term_parser_rule_itemContext) IsNon_term_parser_rule_itemContext() {}

func NewNon_term_parser_rule_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Non_term_parser_rule_itemContext {
	var p = new(Non_term_parser_rule_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarDescParserRULE_non_term_parser_rule_item

	return p
}

func (s *Non_term_parser_rule_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Non_term_parser_rule_itemContext) NON_TERM_ID() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserNON_TERM_ID, 0)
}

func (s *Non_term_parser_rule_itemContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserID, 0)
}

func (s *Non_term_parser_rule_itemContext) ATTR() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserATTR, 0)
}

func (s *Non_term_parser_rule_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Non_term_parser_rule_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Non_term_parser_rule_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.EnterNon_term_parser_rule_item(s)
	}
}

func (s *Non_term_parser_rule_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.ExitNon_term_parser_rule_item(s)
	}
}

func (s *Non_term_parser_rule_itemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarDescVisitor:
		return t.VisitNon_term_parser_rule_item(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarDescParser) Non_term_parser_rule_item() (localctx INon_term_parser_rule_itemContext) {
	localctx = NewNon_term_parser_rule_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarDescParserRULE_non_term_parser_rule_item)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarDescParserID {
		{
			p.SetState(58)
			p.Match(GrammarDescParserID)
		}
		{
			p.SetState(59)
			p.Match(GrammarDescParserT__5)
		}

	}
	{
		p.SetState(62)
		p.Match(GrammarDescParserNON_TERM_ID)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarDescParserATTR {
		{
			p.SetState(63)
			p.Match(GrammarDescParserATTR)
		}

	}

	return localctx
}

// ICode_itemContext is an interface to support dynamic dispatch.
type ICode_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCode_itemContext differentiates from other interfaces.
	IsCode_itemContext()
}

type Code_itemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCode_itemContext() *Code_itemContext {
	var p = new(Code_itemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarDescParserRULE_code_item
	return p
}

func (*Code_itemContext) IsCode_itemContext() {}

func NewCode_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Code_itemContext {
	var p = new(Code_itemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarDescParserRULE_code_item

	return p
}

func (s *Code_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Code_itemContext) CODE() antlr.TerminalNode {
	return s.GetToken(GrammarDescParserCODE, 0)
}

func (s *Code_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Code_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Code_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.EnterCode_item(s)
	}
}

func (s *Code_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarDescListener); ok {
		listenerT.ExitCode_item(s)
	}
}

func (s *Code_itemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarDescVisitor:
		return t.VisitCode_item(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarDescParser) Code_item() (localctx ICode_itemContext) {
	localctx = NewCode_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarDescParserRULE_code_item)

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
		p.SetState(66)
		p.Match(GrammarDescParserCODE)
	}

	return localctx
}
