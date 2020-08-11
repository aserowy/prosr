// Code generated from c:\Users\serow_000\source\repos\prosr\src\parser\Prosr1.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 139,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 7, 2, 43, 10, 2, 12, 2, 14,
	2, 46, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 5, 4, 59, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 66, 10, 5,
	13, 5, 14, 5, 67, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 5, 7, 88, 10, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	5, 10, 102, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 110,
	10, 12, 12, 12, 14, 12, 113, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 125, 10, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	2, 2, 21, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 2, 4, 3, 2, 10, 11, 3, 2, 12, 13, 2, 127, 2, 40, 3, 2, 2, 2, 4,
	49, 3, 2, 2, 2, 6, 58, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 71, 3, 2, 2,
	2, 12, 87, 3, 2, 2, 2, 14, 89, 3, 2, 2, 2, 16, 91, 3, 2, 2, 2, 18, 101,
	3, 2, 2, 2, 20, 103, 3, 2, 2, 2, 22, 105, 3, 2, 2, 2, 24, 116, 3, 2, 2,
	2, 26, 124, 3, 2, 2, 2, 28, 126, 3, 2, 2, 2, 30, 128, 3, 2, 2, 2, 32, 130,
	3, 2, 2, 2, 34, 132, 3, 2, 2, 2, 36, 134, 3, 2, 2, 2, 38, 136, 3, 2, 2,
	2, 40, 44, 5, 4, 3, 2, 41, 43, 5, 6, 4, 2, 42, 41, 3, 2, 2, 2, 43, 46,
	3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2,
	46, 44, 3, 2, 2, 2, 47, 48, 7, 2, 2, 3, 48, 3, 3, 2, 2, 2, 49, 50, 7, 14,
	2, 2, 50, 51, 7, 3, 2, 2, 51, 52, 5, 28, 15, 2, 52, 53, 7, 4, 2, 2, 53,
	54, 5, 28, 15, 2, 54, 55, 7, 5, 2, 2, 55, 5, 3, 2, 2, 2, 56, 59, 5, 8,
	5, 2, 57, 59, 5, 22, 12, 2, 58, 56, 3, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59,
	7, 3, 2, 2, 2, 60, 61, 7, 16, 2, 2, 61, 62, 5, 30, 16, 2, 62, 65, 7, 6,
	2, 2, 63, 66, 5, 10, 6, 2, 64, 66, 5, 16, 9, 2, 65, 63, 3, 2, 2, 2, 65,
	64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2,
	2, 68, 69, 3, 2, 2, 2, 69, 70, 7, 7, 2, 2, 70, 9, 3, 2, 2, 2, 71, 72, 7,
	17, 2, 2, 72, 73, 5, 38, 20, 2, 73, 74, 7, 8, 2, 2, 74, 75, 5, 12, 7, 2,
	75, 76, 7, 9, 2, 2, 76, 77, 7, 18, 2, 2, 77, 78, 5, 36, 19, 2, 78, 79,
	7, 8, 2, 2, 79, 80, 5, 18, 10, 2, 80, 81, 7, 9, 2, 2, 81, 82, 7, 19, 2,
	2, 82, 83, 5, 14, 8, 2, 83, 84, 3, 2, 2, 2, 84, 85, 7, 5, 2, 2, 85, 11,
	3, 2, 2, 2, 86, 88, 5, 34, 18, 2, 87, 86, 3, 2, 2, 2, 87, 88, 3, 2, 2,
	2, 88, 13, 3, 2, 2, 2, 89, 90, 9, 2, 2, 2, 90, 15, 3, 2, 2, 2, 91, 92,
	7, 18, 2, 2, 92, 93, 5, 36, 19, 2, 93, 94, 7, 8, 2, 2, 94, 95, 5, 18, 10,
	2, 95, 96, 7, 9, 2, 2, 96, 97, 7, 19, 2, 2, 97, 98, 5, 20, 11, 2, 98, 99,
	7, 5, 2, 2, 99, 17, 3, 2, 2, 2, 100, 102, 5, 34, 18, 2, 101, 100, 3, 2,
	2, 2, 101, 102, 3, 2, 2, 2, 102, 19, 3, 2, 2, 2, 103, 104, 7, 11, 2, 2,
	104, 21, 3, 2, 2, 2, 105, 106, 7, 15, 2, 2, 106, 107, 5, 34, 18, 2, 107,
	111, 7, 6, 2, 2, 108, 110, 5, 24, 13, 2, 109, 108, 3, 2, 2, 2, 110, 113,
	3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114, 3, 2,
	2, 2, 113, 111, 3, 2, 2, 2, 114, 115, 7, 7, 2, 2, 115, 23, 3, 2, 2, 2,
	116, 117, 5, 26, 14, 2, 117, 118, 5, 32, 17, 2, 118, 119, 7, 3, 2, 2, 119,
	120, 7, 22, 2, 2, 120, 121, 7, 5, 2, 2, 121, 25, 3, 2, 2, 2, 122, 125,
	5, 34, 18, 2, 123, 125, 7, 20, 2, 2, 124, 122, 3, 2, 2, 2, 124, 123, 3,
	2, 2, 2, 125, 27, 3, 2, 2, 2, 126, 127, 9, 3, 2, 2, 127, 29, 3, 2, 2, 2,
	128, 129, 7, 21, 2, 2, 129, 31, 3, 2, 2, 2, 130, 131, 7, 21, 2, 2, 131,
	33, 3, 2, 2, 2, 132, 133, 7, 21, 2, 2, 133, 35, 3, 2, 2, 2, 134, 135, 7,
	21, 2, 2, 135, 37, 3, 2, 2, 2, 136, 137, 7, 21, 2, 2, 137, 39, 3, 2, 2,
	2, 10, 44, 58, 65, 67, 87, 101, 111, 124,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'prosr1'", "';'", "'{'", "'}'", "'('", "')'", "'caller'", "'all'",
	"'''", "'\"'", "'syntax'", "'message'", "'hub'", "'action'", "'calls'",
	"'on'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "SYNTAX", "MESSAGE", "HUB",
	"ACTION", "CALLS", "ON", "TYPE", "IDENT", "NUMBER", "WHITESPACE",
}

var ruleNames = []string{
	"content", "syntax", "definition", "hub", "sending", "sendingMessageIdent",
	"sendingTarget", "returning", "returningMessageIdent", "returningTarget",
	"message", "field", "typeIdent", "quote", "hubIdent", "fieldIdent", "messageIdent",
	"returningIdent", "sendingIdent",
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
	Prosr1ParserEOF        = antlr.TokenEOF
	Prosr1ParserT__0       = 1
	Prosr1ParserT__1       = 2
	Prosr1ParserT__2       = 3
	Prosr1ParserT__3       = 4
	Prosr1ParserT__4       = 5
	Prosr1ParserT__5       = 6
	Prosr1ParserT__6       = 7
	Prosr1ParserT__7       = 8
	Prosr1ParserT__8       = 9
	Prosr1ParserT__9       = 10
	Prosr1ParserT__10      = 11
	Prosr1ParserSYNTAX     = 12
	Prosr1ParserMESSAGE    = 13
	Prosr1ParserHUB        = 14
	Prosr1ParserACTION     = 15
	Prosr1ParserCALLS      = 16
	Prosr1ParserON         = 17
	Prosr1ParserTYPE       = 18
	Prosr1ParserIDENT      = 19
	Prosr1ParserNUMBER     = 20
	Prosr1ParserWHITESPACE = 21
)

// Prosr1Parser rules.
const (
	Prosr1ParserRULE_content               = 0
	Prosr1ParserRULE_syntax                = 1
	Prosr1ParserRULE_definition            = 2
	Prosr1ParserRULE_hub                   = 3
	Prosr1ParserRULE_sending               = 4
	Prosr1ParserRULE_sendingMessageIdent   = 5
	Prosr1ParserRULE_sendingTarget         = 6
	Prosr1ParserRULE_returning             = 7
	Prosr1ParserRULE_returningMessageIdent = 8
	Prosr1ParserRULE_returningTarget       = 9
	Prosr1ParserRULE_message               = 10
	Prosr1ParserRULE_field                 = 11
	Prosr1ParserRULE_typeIdent             = 12
	Prosr1ParserRULE_quote                 = 13
	Prosr1ParserRULE_hubIdent              = 14
	Prosr1ParserRULE_fieldIdent            = 15
	Prosr1ParserRULE_messageIdent          = 16
	Prosr1ParserRULE_returningIdent        = 17
	Prosr1ParserRULE_sendingIdent          = 18
)

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) Syntax() ISyntaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyntaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ContentContext) EOF() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserEOF, 0)
}

func (s *ContentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *ContentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitContent(s)
	}
}

func (p *Prosr1Parser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Prosr1ParserRULE_content)
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
		p.SetState(38)
		p.Syntax()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Prosr1ParserMESSAGE || _la == Prosr1ParserHUB {
		{
			p.SetState(39)
			p.Definition()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(45)
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
		p.SetState(47)
		p.Match(Prosr1ParserSYNTAX)
	}
	{
		p.SetState(48)
		p.Match(Prosr1ParserT__0)
	}
	{
		p.SetState(49)
		p.Quote()
	}
	{
		p.SetState(50)
		p.Match(Prosr1ParserT__1)
	}
	{
		p.SetState(51)
		p.Quote()
	}
	{
		p.SetState(52)
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

func (s *DefinitionContext) Hub() IHubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHubContext)
}

func (s *DefinitionContext) Message() IMessageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageContext)(nil)).Elem(), 0)

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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Prosr1ParserHUB:
		{
			p.SetState(54)
			p.Hub()
		}

	case Prosr1ParserMESSAGE:
		{
			p.SetState(55)
			p.Message()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
		p.SetState(58)
		p.Match(Prosr1ParserHUB)
	}
	{
		p.SetState(59)
		p.HubIdent()
	}
	{
		p.SetState(60)
		p.Match(Prosr1ParserT__3)
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == Prosr1ParserACTION || _la == Prosr1ParserCALLS {
		p.SetState(63)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Prosr1ParserACTION:
			{
				p.SetState(61)
				p.Sending()
			}

		case Prosr1ParserCALLS:
			{
				p.SetState(62)
				p.Returning()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
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

func (s *SendingContext) SendingIdent() ISendingIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISendingIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISendingIdentContext)
}

func (s *SendingContext) SendingMessageIdent() ISendingMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISendingMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISendingMessageIdentContext)
}

func (s *SendingContext) CALLS() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserCALLS, 0)
}

func (s *SendingContext) ReturningIdent() IReturningIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningIdentContext)
}

func (s *SendingContext) ReturningMessageIdent() IReturningMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningMessageIdentContext)
}

func (s *SendingContext) ON() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserON, 0)
}

func (s *SendingContext) SendingTarget() ISendingTargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISendingTargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISendingTargetContext)
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
		p.SetState(69)
		p.Match(Prosr1ParserACTION)
	}
	{
		p.SetState(70)
		p.SendingIdent()
	}
	{
		p.SetState(71)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(72)
		p.SendingMessageIdent()
	}
	{
		p.SetState(73)
		p.Match(Prosr1ParserT__6)
	}

	{
		p.SetState(74)
		p.Match(Prosr1ParserCALLS)
	}
	{
		p.SetState(75)
		p.ReturningIdent()
	}
	{
		p.SetState(76)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(77)
		p.ReturningMessageIdent()
	}
	{
		p.SetState(78)
		p.Match(Prosr1ParserT__6)
	}
	{
		p.SetState(79)
		p.Match(Prosr1ParserON)
	}
	{
		p.SetState(80)
		p.SendingTarget()
	}

	{
		p.SetState(82)
		p.Match(Prosr1ParserT__2)
	}

	return localctx
}

// ISendingMessageIdentContext is an interface to support dynamic dispatch.
type ISendingMessageIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSendingMessageIdentContext differentiates from other interfaces.
	IsSendingMessageIdentContext()
}

type SendingMessageIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendingMessageIdentContext() *SendingMessageIdentContext {
	var p = new(SendingMessageIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_sendingMessageIdent
	return p
}

func (*SendingMessageIdentContext) IsSendingMessageIdentContext() {}

func NewSendingMessageIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendingMessageIdentContext {
	var p = new(SendingMessageIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_sendingMessageIdent

	return p
}

func (s *SendingMessageIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *SendingMessageIdentContext) MessageIdent() IMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageIdentContext)
}

func (s *SendingMessageIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendingMessageIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendingMessageIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterSendingMessageIdent(s)
	}
}

func (s *SendingMessageIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitSendingMessageIdent(s)
	}
}

func (p *Prosr1Parser) SendingMessageIdent() (localctx ISendingMessageIdentContext) {
	localctx = NewSendingMessageIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Prosr1ParserRULE_sendingMessageIdent)
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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserIDENT {
		{
			p.SetState(84)
			p.MessageIdent()
		}

	}

	return localctx
}

// ISendingTargetContext is an interface to support dynamic dispatch.
type ISendingTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSendingTargetContext differentiates from other interfaces.
	IsSendingTargetContext()
}

type SendingTargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendingTargetContext() *SendingTargetContext {
	var p = new(SendingTargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_sendingTarget
	return p
}

func (*SendingTargetContext) IsSendingTargetContext() {}

func NewSendingTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendingTargetContext {
	var p = new(SendingTargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_sendingTarget

	return p
}

func (s *SendingTargetContext) GetParser() antlr.Parser { return s.parser }
func (s *SendingTargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendingTargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendingTargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterSendingTarget(s)
	}
}

func (s *SendingTargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitSendingTarget(s)
	}
}

func (p *Prosr1Parser) SendingTarget() (localctx ISendingTargetContext) {
	localctx = NewSendingTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Prosr1ParserRULE_sendingTarget)
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
		p.SetState(87)
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

func (s *ReturningContext) CALLS() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserCALLS, 0)
}

func (s *ReturningContext) ReturningIdent() IReturningIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningIdentContext)
}

func (s *ReturningContext) ReturningMessageIdent() IReturningMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningMessageIdentContext)
}

func (s *ReturningContext) ON() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserON, 0)
}

func (s *ReturningContext) ReturningTarget() IReturningTargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningTargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningTargetContext)
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
	p.EnterRule(localctx, 14, Prosr1ParserRULE_returning)

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
		p.SetState(89)
		p.Match(Prosr1ParserCALLS)
	}
	{
		p.SetState(90)
		p.ReturningIdent()
	}
	{
		p.SetState(91)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(92)
		p.ReturningMessageIdent()
	}
	{
		p.SetState(93)
		p.Match(Prosr1ParserT__6)
	}
	{
		p.SetState(94)
		p.Match(Prosr1ParserON)
	}
	{
		p.SetState(95)
		p.ReturningTarget()
	}
	{
		p.SetState(96)
		p.Match(Prosr1ParserT__2)
	}

	return localctx
}

// IReturningMessageIdentContext is an interface to support dynamic dispatch.
type IReturningMessageIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturningMessageIdentContext differentiates from other interfaces.
	IsReturningMessageIdentContext()
}

type ReturningMessageIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturningMessageIdentContext() *ReturningMessageIdentContext {
	var p = new(ReturningMessageIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_returningMessageIdent
	return p
}

func (*ReturningMessageIdentContext) IsReturningMessageIdentContext() {}

func NewReturningMessageIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturningMessageIdentContext {
	var p = new(ReturningMessageIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_returningMessageIdent

	return p
}

func (s *ReturningMessageIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturningMessageIdentContext) MessageIdent() IMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageIdentContext)
}

func (s *ReturningMessageIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturningMessageIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturningMessageIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterReturningMessageIdent(s)
	}
}

func (s *ReturningMessageIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitReturningMessageIdent(s)
	}
}

func (p *Prosr1Parser) ReturningMessageIdent() (localctx IReturningMessageIdentContext) {
	localctx = NewReturningMessageIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Prosr1ParserRULE_returningMessageIdent)
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
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserIDENT {
		{
			p.SetState(98)
			p.MessageIdent()
		}

	}

	return localctx
}

// IReturningTargetContext is an interface to support dynamic dispatch.
type IReturningTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturningTargetContext differentiates from other interfaces.
	IsReturningTargetContext()
}

type ReturningTargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturningTargetContext() *ReturningTargetContext {
	var p = new(ReturningTargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_returningTarget
	return p
}

func (*ReturningTargetContext) IsReturningTargetContext() {}

func NewReturningTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturningTargetContext {
	var p = new(ReturningTargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_returningTarget

	return p
}

func (s *ReturningTargetContext) GetParser() antlr.Parser { return s.parser }
func (s *ReturningTargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturningTargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturningTargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterReturningTarget(s)
	}
}

func (s *ReturningTargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitReturningTarget(s)
	}
}

func (p *Prosr1Parser) ReturningTarget() (localctx IReturningTargetContext) {
	localctx = NewReturningTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Prosr1ParserRULE_returningTarget)

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
		p.SetState(101)
		p.Match(Prosr1ParserT__8)
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
	p.EnterRule(localctx, 20, Prosr1ParserRULE_message)
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
		p.SetState(103)
		p.Match(Prosr1ParserMESSAGE)
	}
	{
		p.SetState(104)
		p.MessageIdent()
	}
	{
		p.SetState(105)
		p.Match(Prosr1ParserT__3)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Prosr1ParserTYPE || _la == Prosr1ParserIDENT {
		{
			p.SetState(106)
			p.Field()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
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

func (s *FieldContext) TypeIdent() ITypeIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeIdentContext)
}

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
	p.EnterRule(localctx, 22, Prosr1ParserRULE_field)

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
		p.SetState(114)
		p.TypeIdent()
	}
	{
		p.SetState(115)
		p.FieldIdent()
	}
	{
		p.SetState(116)
		p.Match(Prosr1ParserT__0)
	}
	{
		p.SetState(117)
		p.Match(Prosr1ParserNUMBER)
	}
	{
		p.SetState(118)
		p.Match(Prosr1ParserT__2)
	}

	return localctx
}

// ITypeIdentContext is an interface to support dynamic dispatch.
type ITypeIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeIdentContext differentiates from other interfaces.
	IsTypeIdentContext()
}

type TypeIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeIdentContext() *TypeIdentContext {
	var p = new(TypeIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_typeIdent
	return p
}

func (*TypeIdentContext) IsTypeIdentContext() {}

func NewTypeIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeIdentContext {
	var p = new(TypeIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_typeIdent

	return p
}

func (s *TypeIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeIdentContext) MessageIdent() IMessageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageIdentContext)
}

func (s *TypeIdentContext) TYPE() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserTYPE, 0)
}

func (s *TypeIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterTypeIdent(s)
	}
}

func (s *TypeIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitTypeIdent(s)
	}
}

func (p *Prosr1Parser) TypeIdent() (localctx ITypeIdentContext) {
	localctx = NewTypeIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Prosr1ParserRULE_typeIdent)

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
	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Prosr1ParserIDENT:
		{
			p.SetState(120)
			p.MessageIdent()
		}

	case Prosr1ParserTYPE:
		{
			p.SetState(121)
			p.Match(Prosr1ParserTYPE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 26, Prosr1ParserRULE_quote)
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
		p.SetState(124)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Prosr1ParserT__9 || _la == Prosr1ParserT__10) {
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
	p.EnterRule(localctx, 28, Prosr1ParserRULE_hubIdent)

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
		p.SetState(126)
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
	p.EnterRule(localctx, 30, Prosr1ParserRULE_fieldIdent)

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
		p.SetState(128)
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
	p.EnterRule(localctx, 32, Prosr1ParserRULE_messageIdent)

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
		p.SetState(130)
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}

// IReturningIdentContext is an interface to support dynamic dispatch.
type IReturningIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturningIdentContext differentiates from other interfaces.
	IsReturningIdentContext()
}

type ReturningIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturningIdentContext() *ReturningIdentContext {
	var p = new(ReturningIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_returningIdent
	return p
}

func (*ReturningIdentContext) IsReturningIdentContext() {}

func NewReturningIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturningIdentContext {
	var p = new(ReturningIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_returningIdent

	return p
}

func (s *ReturningIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturningIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *ReturningIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturningIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturningIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterReturningIdent(s)
	}
}

func (s *ReturningIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitReturningIdent(s)
	}
}

func (p *Prosr1Parser) ReturningIdent() (localctx IReturningIdentContext) {
	localctx = NewReturningIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Prosr1ParserRULE_returningIdent)

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
		p.SetState(132)
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}

// ISendingIdentContext is an interface to support dynamic dispatch.
type ISendingIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSendingIdentContext differentiates from other interfaces.
	IsSendingIdentContext()
}

type SendingIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendingIdentContext() *SendingIdentContext {
	var p = new(SendingIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_sendingIdent
	return p
}

func (*SendingIdentContext) IsSendingIdentContext() {}

func NewSendingIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendingIdentContext {
	var p = new(SendingIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_sendingIdent

	return p
}

func (s *SendingIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *SendingIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *SendingIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendingIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendingIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterSendingIdent(s)
	}
}

func (s *SendingIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitSendingIdent(s)
	}
}

func (p *Prosr1Parser) SendingIdent() (localctx ISendingIdentContext) {
	localctx = NewSendingIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Prosr1ParserRULE_sendingIdent)

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
		p.SetState(134)
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}
