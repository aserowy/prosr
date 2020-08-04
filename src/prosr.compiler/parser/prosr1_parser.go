// Code generated from c:\Users\serow_000\source\repos\prosr\src\prosr.compiler\parser\Prosr1.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Prosr1

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 111,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 42, 10, 4, 12, 4, 14, 4, 45, 11, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 52, 10, 5, 13, 5, 14, 5, 53, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 7, 8, 85, 10, 8, 12, 8, 14, 8, 88, 11, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 5, 9, 94, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 2, 2,
	15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 3, 3, 2, 10, 11,
	2, 103, 2, 28, 3, 2, 2, 2, 4, 32, 3, 2, 2, 2, 6, 43, 3, 2, 2, 2, 8, 46,
	3, 2, 2, 2, 10, 57, 3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 14, 80, 3, 2, 2, 2,
	16, 93, 3, 2, 2, 2, 18, 100, 3, 2, 2, 2, 20, 102, 3, 2, 2, 2, 22, 104,
	3, 2, 2, 2, 24, 106, 3, 2, 2, 2, 26, 108, 3, 2, 2, 2, 28, 29, 5, 4, 3,
	2, 29, 30, 5, 6, 4, 2, 30, 31, 7, 2, 2, 3, 31, 3, 3, 2, 2, 2, 32, 33, 7,
	12, 2, 2, 33, 34, 7, 3, 2, 2, 34, 35, 5, 18, 10, 2, 35, 36, 7, 4, 2, 2,
	36, 37, 5, 18, 10, 2, 37, 38, 7, 5, 2, 2, 38, 5, 3, 2, 2, 2, 39, 42, 5,
	8, 5, 2, 40, 42, 5, 14, 8, 2, 41, 39, 3, 2, 2, 2, 41, 40, 3, 2, 2, 2, 42,
	45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 7, 3, 2, 2,
	2, 45, 43, 3, 2, 2, 2, 46, 47, 7, 14, 2, 2, 47, 48, 5, 20, 11, 2, 48, 51,
	7, 6, 2, 2, 49, 52, 5, 10, 6, 2, 50, 52, 5, 12, 7, 2, 51, 49, 3, 2, 2,
	2, 51, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 7, 7, 2, 2, 56, 9, 3, 2, 2, 2,
	57, 58, 7, 15, 2, 2, 58, 59, 5, 22, 12, 2, 59, 60, 7, 8, 2, 2, 60, 61,
	5, 24, 13, 2, 61, 62, 7, 9, 2, 2, 62, 63, 7, 17, 2, 2, 63, 64, 7, 8, 2,
	2, 64, 65, 5, 24, 13, 2, 65, 66, 7, 9, 2, 2, 66, 67, 7, 19, 2, 2, 67, 68,
	7, 16, 2, 2, 68, 69, 7, 5, 2, 2, 69, 70, 3, 2, 2, 2, 70, 71, 7, 5, 2, 2,
	71, 11, 3, 2, 2, 2, 72, 73, 7, 17, 2, 2, 73, 74, 7, 8, 2, 2, 74, 75, 5,
	24, 13, 2, 75, 76, 7, 9, 2, 2, 76, 77, 7, 19, 2, 2, 77, 78, 7, 18, 2, 2,
	78, 79, 7, 5, 2, 2, 79, 13, 3, 2, 2, 2, 80, 81, 7, 13, 2, 2, 81, 82, 5,
	24, 13, 2, 82, 86, 7, 6, 2, 2, 83, 85, 5, 16, 9, 2, 84, 83, 3, 2, 2, 2,
	85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3,
	2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90, 7, 7, 2, 2, 90, 15, 3, 2, 2, 2, 91,
	94, 5, 24, 13, 2, 92, 94, 7, 20, 2, 2, 93, 91, 3, 2, 2, 2, 93, 92, 3, 2,
	2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 5, 26, 14, 2, 96, 97, 7, 3, 2, 2, 97,
	98, 7, 22, 2, 2, 98, 99, 7, 5, 2, 2, 99, 17, 3, 2, 2, 2, 100, 101, 9, 2,
	2, 2, 101, 19, 3, 2, 2, 2, 102, 103, 7, 21, 2, 2, 103, 21, 3, 2, 2, 2,
	104, 105, 7, 21, 2, 2, 105, 23, 3, 2, 2, 2, 106, 107, 7, 21, 2, 2, 107,
	25, 3, 2, 2, 2, 108, 109, 7, 21, 2, 2, 109, 27, 3, 2, 2, 2, 8, 41, 43,
	51, 53, 86, 93,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'prosr1'", "';'", "'{'", "'}'", "'('", "')'", "'''", "'\"'",
	"'syntax'", "'message'", "'hub'", "'action'", "", "'returns'", "'all'",
	"'to'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "SYNTAX", "MESSAGE", "HUB", "ACTION",
	"ACTIONTARGET", "RETURNS", "RETURNSTARGET", "TO", "TYPE", "IDENT", "NUMBER",
	"WHITESPACE",
}

var ruleNames = []string{
	"main", "syntax", "definition", "hub", "sending", "returning", "message",
	"field", "quote", "hubIdent", "actionIdent", "messageIdent", "fieldIdent",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Prosr1Parser struct {
	*antlr.BaseParser
}

func NewProsr1Parser(input antlr.TokenStream) *Prosr1Parser {
	this := new(Prosr1Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Prosr1.g4"

	return this
}

// Prosr1Parser tokens.
const (
	Prosr1ParserEOF           = antlr.TokenEOF
	Prosr1ParserT__0          = 1
	Prosr1ParserT__1          = 2
	Prosr1ParserT__2          = 3
	Prosr1ParserT__3          = 4
	Prosr1ParserT__4          = 5
	Prosr1ParserT__5          = 6
	Prosr1ParserT__6          = 7
	Prosr1ParserT__7          = 8
	Prosr1ParserT__8          = 9
	Prosr1ParserSYNTAX        = 10
	Prosr1ParserMESSAGE       = 11
	Prosr1ParserHUB           = 12
	Prosr1ParserACTION        = 13
	Prosr1ParserACTIONTARGET  = 14
	Prosr1ParserRETURNS       = 15
	Prosr1ParserRETURNSTARGET = 16
	Prosr1ParserTO            = 17
	Prosr1ParserTYPE          = 18
	Prosr1ParserIDENT         = 19
	Prosr1ParserNUMBER        = 20
	Prosr1ParserWHITESPACE    = 21
)

// Prosr1Parser rules.
const (
	Prosr1ParserRULE_main         = 0
	Prosr1ParserRULE_syntax       = 1
	Prosr1ParserRULE_definition   = 2
	Prosr1ParserRULE_hub          = 3
	Prosr1ParserRULE_sending      = 4
	Prosr1ParserRULE_returning    = 5
	Prosr1ParserRULE_message      = 6
	Prosr1ParserRULE_field        = 7
	Prosr1ParserRULE_quote        = 8
	Prosr1ParserRULE_hubIdent     = 9
	Prosr1ParserRULE_actionIdent  = 10
	Prosr1ParserRULE_messageIdent = 11
	Prosr1ParserRULE_fieldIdent   = 12
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *MainContext) Definition() IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *MainContext) EOF() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserEOF, 0)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *Prosr1Parser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Prosr1ParserRULE_main)

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
		p.SetState(26)
		p.Syntax()
	}
	{
		p.SetState(27)
		p.Definition()
	}
	{
		p.SetState(28)
		p.Match(Prosr1ParserEOF)
	}

	return localctx
}

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserSYNTAX, 0)
}

func (s *SyntaxContext) AllQuote() []IQuoteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuoteContext)(nil)).Elem())
	var tst = make([]IQuoteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuoteContext)
		}
	}

	return tst
}

func (s *SyntaxContext) Quote(i int) IQuoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuoteContext)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (p *Prosr1Parser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Prosr1ParserRULE_syntax)

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
		p.Match(Prosr1ParserSYNTAX)
	}
	{
		p.SetState(31)
		p.Match(Prosr1ParserT__0)
	}
	{
		p.SetState(32)
		p.Quote()
	}
	{
		p.SetState(33)
		p.Match(Prosr1ParserT__1)
	}
	{
		p.SetState(34)
		p.Quote()
	}
	{
		p.SetState(35)
		p.Match(Prosr1ParserT__2)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) AllHub() []IHubContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHubContext)(nil)).Elem())
	var tst = make([]IHubContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHubContext)
		}
	}

	return tst
}

func (s *DefinitionContext) Hub(i int) IHubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHubContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHubContext)
}

func (s *DefinitionContext) AllMessage() []IMessageContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageContext)(nil)).Elem())
	var tst = make([]IMessageContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageContext)
		}
	}

	return tst
}

func (s *DefinitionContext) Message(i int) IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *Prosr1Parser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Prosr1ParserRULE_definition)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Prosr1ParserMESSAGE || _la == Prosr1ParserHUB {
		p.SetState(39)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Prosr1ParserHUB:
			{
				p.SetState(37)
				p.Hub()
			}

		case Prosr1ParserMESSAGE:
			{
				p.SetState(38)
				p.Message()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHubContext is an interface to support dynamic dispatch.
type IHubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHubContext differentiates from other interfaces.
	IsHubContext()
}

type HubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHubContext() *HubContext {
	var p = new(HubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_hub
	return p
}

func (*HubContext) IsHubContext() {}

func NewHubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HubContext {
	var p = new(HubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_hub

	return p
}

func (s *HubContext) GetParser() antlr.Parser { return s.parser }

func (s *HubContext) HUB() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserHUB, 0)
}

func (s *HubContext) HubIdent() IHubIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHubIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHubIdentContext)
}

func (s *HubContext) AllSending() []ISendingContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISendingContext)(nil)).Elem())
	var tst = make([]ISendingContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISendingContext)
		}
	}

	return tst
}

func (s *HubContext) Sending(i int) ISendingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISendingContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISendingContext)
}

func (s *HubContext) AllReturning() []IReturningContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReturningContext)(nil)).Elem())
	var tst = make([]IReturningContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReturningContext)
		}
	}

	return tst
}

func (s *HubContext) Returning(i int) IReturningContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReturningContext)
}

func (s *HubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterHub(s)
	}
}

func (s *HubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitHub(s)
	}
}

func (p *Prosr1Parser) Hub() (localctx IHubContext) {
	localctx = NewHubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Prosr1ParserRULE_hub)
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
		p.Match(Prosr1ParserHUB)
	}
	{
		p.SetState(45)
		p.HubIdent()
	}
	{
		p.SetState(46)
		p.Match(Prosr1ParserT__3)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == Prosr1ParserACTION || _la == Prosr1ParserRETURNS {
		p.SetState(49)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Prosr1ParserACTION:
			{
				p.SetState(47)
				p.Sending()
			}

		case Prosr1ParserRETURNS:
			{
				p.SetState(48)
				p.Returning()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(Prosr1ParserT__4)
	}

	return localctx
}

// ISendingContext is an interface to support dynamic dispatch.
type ISendingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSendingContext differentiates from other interfaces.
	IsSendingContext()
}

type SendingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendingContext() *SendingContext {
	var p = new(SendingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_sending
	return p
}

func (*SendingContext) IsSendingContext() {}

func NewSendingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendingContext {
	var p = new(SendingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_sending

	return p
}

func (s *SendingContext) GetParser() antlr.Parser { return s.parser }

func (s *SendingContext) ACTION() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserACTION, 0)
}

func (s *SendingContext) ActionIdent() IActionIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IActionIdentContext)
}

func (s *SendingContext) AllMessageIdent() []IMessageIdentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMessageIdentContext)(nil)).Elem())
	var tst = make([]IMessageIdentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMessageIdentContext)
		}
	}

	return tst
}

func (s *SendingContext) MessageIdent(i int) IMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageIdentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMessageIdentContext)
}

func (s *SendingContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserRETURNS, 0)
}

func (s *SendingContext) TO() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserTO, 0)
}

func (s *SendingContext) ACTIONTARGET() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserACTIONTARGET, 0)
}

func (s *SendingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterSending(s)
	}
}

func (s *SendingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitSending(s)
	}
}

func (p *Prosr1Parser) Sending() (localctx ISendingContext) {
	localctx = NewSendingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Prosr1ParserRULE_sending)

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
		p.SetState(55)
		p.Match(Prosr1ParserACTION)
	}
	{
		p.SetState(56)
		p.ActionIdent()
	}
	{
		p.SetState(57)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(58)
		p.MessageIdent()
	}
	{
		p.SetState(59)
		p.Match(Prosr1ParserT__6)
	}

	{
		p.SetState(60)
		p.Match(Prosr1ParserRETURNS)
	}
	{
		p.SetState(61)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(62)
		p.MessageIdent()
	}
	{
		p.SetState(63)
		p.Match(Prosr1ParserT__6)
	}
	{
		p.SetState(64)
		p.Match(Prosr1ParserTO)
	}
	{
		p.SetState(65)
		p.Match(Prosr1ParserACTIONTARGET)
	}
	{
		p.SetState(66)
		p.Match(Prosr1ParserT__2)
	}

	{
		p.SetState(68)
		p.Match(Prosr1ParserT__2)
	}

	return localctx
}

// IReturningContext is an interface to support dynamic dispatch.
type IReturningContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturningContext differentiates from other interfaces.
	IsReturningContext()
}

type ReturningContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturningContext() *ReturningContext {
	var p = new(ReturningContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_returning
	return p
}

func (*ReturningContext) IsReturningContext() {}

func NewReturningContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturningContext {
	var p = new(ReturningContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_returning

	return p
}

func (s *ReturningContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturningContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserRETURNS, 0)
}

func (s *ReturningContext) MessageIdent() IMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageIdentContext)
}

func (s *ReturningContext) TO() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserTO, 0)
}

func (s *ReturningContext) RETURNSTARGET() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserRETURNSTARGET, 0)
}

func (s *ReturningContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturningContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturningContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterReturning(s)
	}
}

func (s *ReturningContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitReturning(s)
	}
}

func (p *Prosr1Parser) Returning() (localctx IReturningContext) {
	localctx = NewReturningContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Prosr1ParserRULE_returning)

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
		p.SetState(70)
		p.Match(Prosr1ParserRETURNS)
	}
	{
		p.SetState(71)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(72)
		p.MessageIdent()
	}
	{
		p.SetState(73)
		p.Match(Prosr1ParserT__6)
	}
	{
		p.SetState(74)
		p.Match(Prosr1ParserTO)
	}
	{
		p.SetState(75)
		p.Match(Prosr1ParserRETURNSTARGET)
	}
	{
		p.SetState(76)
		p.Match(Prosr1ParserT__2)
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserMESSAGE, 0)
}

func (s *MessageContext) MessageIdent() IMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageIdentContext)
}

func (s *MessageContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *MessageContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitMessage(s)
	}
}

func (p *Prosr1Parser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Prosr1ParserRULE_message)
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
		p.SetState(78)
		p.Match(Prosr1ParserMESSAGE)
	}
	{
		p.SetState(79)
		p.MessageIdent()
	}
	{
		p.SetState(80)
		p.Match(Prosr1ParserT__3)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Prosr1ParserTYPE || _la == Prosr1ParserIDENT {
		{
			p.SetState(81)
			p.Field()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
		p.Match(Prosr1ParserT__4)
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
	p.RuleIndex = Prosr1ParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldIdent() IFieldIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldIdentContext)
}

func (s *FieldContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserNUMBER, 0)
}

func (s *FieldContext) MessageIdent() IMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageIdentContext)
}

func (s *FieldContext) TYPE() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserTYPE, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitField(s)
	}
}

func (p *Prosr1Parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Prosr1ParserRULE_field)

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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Prosr1ParserIDENT:
		{
			p.SetState(89)
			p.MessageIdent()
		}

	case Prosr1ParserTYPE:
		{
			p.SetState(90)
			p.Match(Prosr1ParserTYPE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(93)
		p.FieldIdent()
	}
	{
		p.SetState(94)
		p.Match(Prosr1ParserT__0)
	}
	{
		p.SetState(95)
		p.Match(Prosr1ParserNUMBER)
	}
	{
		p.SetState(96)
		p.Match(Prosr1ParserT__2)
	}

	return localctx
}

// IQuoteContext is an interface to support dynamic dispatch.
type IQuoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuoteContext differentiates from other interfaces.
	IsQuoteContext()
}

type QuoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoteContext() *QuoteContext {
	var p = new(QuoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_quote
	return p
}

func (*QuoteContext) IsQuoteContext() {}

func NewQuoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuoteContext {
	var p = new(QuoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_quote

	return p
}

func (s *QuoteContext) GetParser() antlr.Parser { return s.parser }
func (s *QuoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterQuote(s)
	}
}

func (s *QuoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitQuote(s)
	}
}

func (p *Prosr1Parser) Quote() (localctx IQuoteContext) {
	localctx = NewQuoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Prosr1ParserRULE_quote)
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
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Prosr1ParserT__7 || _la == Prosr1ParserT__8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHubIdentContext is an interface to support dynamic dispatch.
type IHubIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHubIdentContext differentiates from other interfaces.
	IsHubIdentContext()
}

type HubIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHubIdentContext() *HubIdentContext {
	var p = new(HubIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_hubIdent
	return p
}

func (*HubIdentContext) IsHubIdentContext() {}

func NewHubIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HubIdentContext {
	var p = new(HubIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_hubIdent

	return p
}

func (s *HubIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *HubIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *HubIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HubIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HubIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterHubIdent(s)
	}
}

func (s *HubIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitHubIdent(s)
	}
}

func (p *Prosr1Parser) HubIdent() (localctx IHubIdentContext) {
	localctx = NewHubIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Prosr1ParserRULE_hubIdent)

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
		p.SetState(100)
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}

// IActionIdentContext is an interface to support dynamic dispatch.
type IActionIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionIdentContext differentiates from other interfaces.
	IsActionIdentContext()
}

type ActionIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionIdentContext() *ActionIdentContext {
	var p = new(ActionIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_actionIdent
	return p
}

func (*ActionIdentContext) IsActionIdentContext() {}

func NewActionIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionIdentContext {
	var p = new(ActionIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_actionIdent

	return p
}

func (s *ActionIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *ActionIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterActionIdent(s)
	}
}

func (s *ActionIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitActionIdent(s)
	}
}

func (p *Prosr1Parser) ActionIdent() (localctx IActionIdentContext) {
	localctx = NewActionIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Prosr1ParserRULE_actionIdent)

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
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}

// IMessageIdentContext is an interface to support dynamic dispatch.
type IMessageIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageIdentContext differentiates from other interfaces.
	IsMessageIdentContext()
}

type MessageIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageIdentContext() *MessageIdentContext {
	var p = new(MessageIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_messageIdent
	return p
}

func (*MessageIdentContext) IsMessageIdentContext() {}

func NewMessageIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageIdentContext {
	var p = new(MessageIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_messageIdent

	return p
}

func (s *MessageIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *MessageIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterMessageIdent(s)
	}
}

func (s *MessageIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitMessageIdent(s)
	}
}

func (p *Prosr1Parser) MessageIdent() (localctx IMessageIdentContext) {
	localctx = NewMessageIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Prosr1ParserRULE_messageIdent)

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
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}

// IFieldIdentContext is an interface to support dynamic dispatch.
type IFieldIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldIdentContext differentiates from other interfaces.
	IsFieldIdentContext()
}

type FieldIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldIdentContext() *FieldIdentContext {
	var p = new(FieldIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_fieldIdent
	return p
}

func (*FieldIdentContext) IsFieldIdentContext() {}

func NewFieldIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldIdentContext {
	var p = new(FieldIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_fieldIdent

	return p
}

func (s *FieldIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *FieldIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterFieldIdent(s)
	}
}

func (s *FieldIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitFieldIdent(s)
	}
}

func (p *Prosr1Parser) FieldIdent() (localctx IFieldIdentContext) {
	localctx = NewFieldIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Prosr1ParserRULE_fieldIdent)

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
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}
