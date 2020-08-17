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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 172,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 7, 2, 49, 10, 2, 12, 2, 14, 2, 52, 11, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 66, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6, 77, 10, 6,
	13, 6, 14, 6, 78, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 96, 10, 7, 3, 7, 3, 7, 3, 8,
	5, 8, 101, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 11, 5, 11, 115, 10, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 7, 13, 123, 10, 13, 12, 13, 14, 13, 126, 11, 13, 3, 13,
	3, 13, 3, 14, 5, 14, 131, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 143, 10, 15, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 152, 10, 18, 12, 18, 14, 18, 155, 11,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 163, 10, 21, 12, 21,
	14, 21, 166, 11, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 2, 4, 3, 2, 10, 11, 3, 2, 12, 13, 2, 162, 2, 46, 3, 2, 2, 2, 4, 55,
	3, 2, 2, 2, 6, 65, 3, 2, 2, 2, 8, 67, 3, 2, 2, 2, 10, 71, 3, 2, 2, 2, 12,
	82, 3, 2, 2, 2, 14, 100, 3, 2, 2, 2, 16, 102, 3, 2, 2, 2, 18, 104, 3, 2,
	2, 2, 20, 114, 3, 2, 2, 2, 22, 116, 3, 2, 2, 2, 24, 118, 3, 2, 2, 2, 26,
	130, 3, 2, 2, 2, 28, 142, 3, 2, 2, 2, 30, 144, 3, 2, 2, 2, 32, 146, 3,
	2, 2, 2, 34, 148, 3, 2, 2, 2, 36, 156, 3, 2, 2, 2, 38, 158, 3, 2, 2, 2,
	40, 164, 3, 2, 2, 2, 42, 167, 3, 2, 2, 2, 44, 169, 3, 2, 2, 2, 46, 50,
	5, 4, 3, 2, 47, 49, 5, 6, 4, 2, 48, 47, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2,
	50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 53, 3, 2, 2, 2, 52, 50, 3,
	2, 2, 2, 53, 54, 7, 2, 2, 3, 54, 3, 3, 2, 2, 2, 55, 56, 7, 15, 2, 2, 56,
	57, 7, 3, 2, 2, 57, 58, 5, 30, 16, 2, 58, 59, 7, 4, 2, 2, 59, 60, 5, 30,
	16, 2, 60, 61, 7, 5, 2, 2, 61, 5, 3, 2, 2, 2, 62, 66, 5, 8, 5, 2, 63, 66,
	5, 10, 6, 2, 64, 66, 5, 24, 13, 2, 65, 62, 3, 2, 2, 2, 65, 63, 3, 2, 2,
	2, 65, 64, 3, 2, 2, 2, 66, 7, 3, 2, 2, 2, 67, 68, 7, 16, 2, 2, 68, 69,
	5, 34, 18, 2, 69, 70, 7, 5, 2, 2, 70, 9, 3, 2, 2, 2, 71, 72, 7, 18, 2,
	2, 72, 73, 5, 36, 19, 2, 73, 76, 7, 6, 2, 2, 74, 77, 5, 12, 7, 2, 75, 77,
	5, 18, 10, 2, 76, 74, 3, 2, 2, 2, 76, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2,
	2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81,
	7, 7, 2, 2, 81, 11, 3, 2, 2, 2, 82, 83, 7, 19, 2, 2, 83, 84, 5, 44, 23,
	2, 84, 85, 7, 8, 2, 2, 85, 86, 5, 14, 8, 2, 86, 95, 7, 9, 2, 2, 87, 88,
	7, 20, 2, 2, 88, 89, 5, 42, 22, 2, 89, 90, 7, 8, 2, 2, 90, 91, 5, 20, 11,
	2, 91, 92, 7, 9, 2, 2, 92, 93, 7, 21, 2, 2, 93, 94, 5, 16, 9, 2, 94, 96,
	3, 2, 2, 2, 95, 87, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2,
	97, 98, 7, 5, 2, 2, 98, 13, 3, 2, 2, 2, 99, 101, 5, 38, 20, 2, 100, 99,
	3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 15, 3, 2, 2, 2, 102, 103, 9, 2,
	2, 2, 103, 17, 3, 2, 2, 2, 104, 105, 7, 20, 2, 2, 105, 106, 5, 42, 22,
	2, 106, 107, 7, 8, 2, 2, 107, 108, 5, 20, 11, 2, 108, 109, 7, 9, 2, 2,
	109, 110, 7, 21, 2, 2, 110, 111, 5, 22, 12, 2, 111, 112, 7, 5, 2, 2, 112,
	19, 3, 2, 2, 2, 113, 115, 5, 38, 20, 2, 114, 113, 3, 2, 2, 2, 114, 115,
	3, 2, 2, 2, 115, 21, 3, 2, 2, 2, 116, 117, 7, 11, 2, 2, 117, 23, 3, 2,
	2, 2, 118, 119, 7, 17, 2, 2, 119, 120, 5, 38, 20, 2, 120, 124, 7, 6, 2,
	2, 121, 123, 5, 26, 14, 2, 122, 121, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2,
	124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 127, 128, 7, 7, 2, 2, 128, 25, 3, 2, 2, 2, 129, 131, 7,
	22, 2, 2, 130, 129, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 3, 2, 2,
	2, 132, 133, 5, 28, 15, 2, 133, 134, 5, 32, 17, 2, 134, 135, 7, 3, 2, 2,
	135, 136, 7, 25, 2, 2, 136, 137, 7, 5, 2, 2, 137, 27, 3, 2, 2, 2, 138,
	139, 5, 40, 21, 2, 139, 140, 5, 38, 20, 2, 140, 143, 3, 2, 2, 2, 141, 143,
	7, 23, 2, 2, 142, 138, 3, 2, 2, 2, 142, 141, 3, 2, 2, 2, 143, 29, 3, 2,
	2, 2, 144, 145, 9, 3, 2, 2, 145, 31, 3, 2, 2, 2, 146, 147, 7, 24, 2, 2,
	147, 33, 3, 2, 2, 2, 148, 153, 7, 24, 2, 2, 149, 150, 7, 14, 2, 2, 150,
	152, 7, 24, 2, 2, 151, 149, 3, 2, 2, 2, 152, 155, 3, 2, 2, 2, 153, 151,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 35, 3, 2, 2, 2, 155, 153, 3, 2,
	2, 2, 156, 157, 7, 24, 2, 2, 157, 37, 3, 2, 2, 2, 158, 159, 7, 24, 2, 2,
	159, 39, 3, 2, 2, 2, 160, 161, 7, 24, 2, 2, 161, 163, 7, 14, 2, 2, 162,
	160, 3, 2, 2, 2, 163, 166, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164, 165,
	3, 2, 2, 2, 165, 41, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 167, 168, 7, 24,
	2, 2, 168, 43, 3, 2, 2, 2, 169, 170, 7, 24, 2, 2, 170, 45, 3, 2, 2, 2,
	14, 50, 65, 76, 78, 95, 100, 114, 124, 130, 142, 153, 164,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'prosr1'", "';'", "'{'", "'}'", "'('", "')'", "'caller'", "'all'",
	"'''", "'\"'", "'.'", "'syntax'", "'package'", "'message'", "'hub'", "'action'",
	"'calls'", "'on'", "'repeated'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "SYNTAX", "PACKAGE",
	"MESSAGE", "HUB", "ACTION", "CALLS", "ON", "REPEATED", "TYPE", "IDENT",
	"NUMBER", "WHITESPACE",
}

var ruleNames = []string{
	"content", "syntax", "definition", "pkg", "hub", "sending", "sendingMessageIdent",
	"sendingTarget", "returning", "returningMessageIdent", "returningTarget",
	"message", "field", "typeIdent", "quote", "fieldIdent", "fullIdent", "hubIdent",
	"messageIdent", "packageIdent", "returningIdent", "sendingIdent",
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
	Prosr1ParserT__11      = 12
	Prosr1ParserSYNTAX     = 13
	Prosr1ParserPACKAGE    = 14
	Prosr1ParserMESSAGE    = 15
	Prosr1ParserHUB        = 16
	Prosr1ParserACTION     = 17
	Prosr1ParserCALLS      = 18
	Prosr1ParserON         = 19
	Prosr1ParserREPEATED   = 20
	Prosr1ParserTYPE       = 21
	Prosr1ParserIDENT      = 22
	Prosr1ParserNUMBER     = 23
	Prosr1ParserWHITESPACE = 24
)

// Prosr1Parser rules.
const (
	Prosr1ParserRULE_content               = 0
	Prosr1ParserRULE_syntax                = 1
	Prosr1ParserRULE_definition            = 2
	Prosr1ParserRULE_pkg                   = 3
	Prosr1ParserRULE_hub                   = 4
	Prosr1ParserRULE_sending               = 5
	Prosr1ParserRULE_sendingMessageIdent   = 6
	Prosr1ParserRULE_sendingTarget         = 7
	Prosr1ParserRULE_returning             = 8
	Prosr1ParserRULE_returningMessageIdent = 9
	Prosr1ParserRULE_returningTarget       = 10
	Prosr1ParserRULE_message               = 11
	Prosr1ParserRULE_field                 = 12
	Prosr1ParserRULE_typeIdent             = 13
	Prosr1ParserRULE_quote                 = 14
	Prosr1ParserRULE_fieldIdent            = 15
	Prosr1ParserRULE_fullIdent             = 16
	Prosr1ParserRULE_hubIdent              = 17
	Prosr1ParserRULE_messageIdent          = 18
	Prosr1ParserRULE_packageIdent          = 19
	Prosr1ParserRULE_returningIdent        = 20
	Prosr1ParserRULE_sendingIdent          = 21
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
		p.SetState(44)
		p.Syntax()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Prosr1ParserPACKAGE)|(1<<Prosr1ParserMESSAGE)|(1<<Prosr1ParserHUB))) != 0 {
		{
			p.SetState(45)
			p.Definition()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
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
		p.SetState(53)
		p.Match(Prosr1ParserSYNTAX)
	}
	{
		p.SetState(54)
		p.Match(Prosr1ParserT__0)
	}
	{
		p.SetState(55)
		p.Quote()
	}
	{
		p.SetState(56)
		p.Match(Prosr1ParserT__1)
	}
	{
		p.SetState(57)
		p.Quote()
	}
	{
		p.SetState(58)
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

func (s *DefinitionContext) Pkg() IPkgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPkgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPkgContext)
}

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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Prosr1ParserPACKAGE:
		{
			p.SetState(60)
			p.Pkg()
		}

	case Prosr1ParserHUB:
		{
			p.SetState(61)
			p.Hub()
		}

	case Prosr1ParserMESSAGE:
		{
			p.SetState(62)
			p.Message()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPkgContext is an interface to support dynamic dispatch.
type IPkgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPkgContext differentiates from other interfaces.
	IsPkgContext()
}

type PkgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPkgContext() *PkgContext {
	var p = new(PkgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_pkg
	return p
}

func (*PkgContext) IsPkgContext() {}

func NewPkgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PkgContext {
	var p = new(PkgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_pkg

	return p
}

func (s *PkgContext) GetParser() antlr.Parser { return s.parser }

func (s *PkgContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserPACKAGE, 0)
}

func (s *PkgContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PkgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PkgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterPkg(s)
	}
}

func (s *PkgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitPkg(s)
	}
}

func (p *Prosr1Parser) Pkg() (localctx IPkgContext) {
	localctx = NewPkgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Prosr1ParserRULE_pkg)

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
		p.SetState(65)
		p.Match(Prosr1ParserPACKAGE)
	}
	{
		p.SetState(66)
		p.FullIdent()
	}
	{
		p.SetState(67)
		p.Match(Prosr1ParserT__2)
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
	p.EnterRule(localctx, 8, Prosr1ParserRULE_hub)
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
		p.SetState(69)
		p.Match(Prosr1ParserHUB)
	}
	{
		p.SetState(70)
		p.HubIdent()
	}
	{
		p.SetState(71)
		p.Match(Prosr1ParserT__3)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == Prosr1ParserACTION || _la == Prosr1ParserCALLS {
		p.SetState(74)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Prosr1ParserACTION:
			{
				p.SetState(72)
				p.Sending()
			}

		case Prosr1ParserCALLS:
			{
				p.SetState(73)
				p.Returning()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
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
	p.EnterRule(localctx, 10, Prosr1ParserRULE_sending)
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
		p.SetState(80)
		p.Match(Prosr1ParserACTION)
	}
	{
		p.SetState(81)
		p.SendingIdent()
	}
	{
		p.SetState(82)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(83)
		p.SendingMessageIdent()
	}
	{
		p.SetState(84)
		p.Match(Prosr1ParserT__6)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserCALLS {
		{
			p.SetState(85)
			p.Match(Prosr1ParserCALLS)
		}
		{
			p.SetState(86)
			p.ReturningIdent()
		}
		{
			p.SetState(87)
			p.Match(Prosr1ParserT__5)
		}
		{
			p.SetState(88)
			p.ReturningMessageIdent()
		}
		{
			p.SetState(89)
			p.Match(Prosr1ParserT__6)
		}
		{
			p.SetState(90)
			p.Match(Prosr1ParserON)
		}
		{
			p.SetState(91)
			p.SendingTarget()
		}

	}
	{
		p.SetState(95)
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
	p.EnterRule(localctx, 12, Prosr1ParserRULE_sendingMessageIdent)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserIDENT {
		{
			p.SetState(97)
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
	p.EnterRule(localctx, 14, Prosr1ParserRULE_sendingTarget)
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
		p.SetState(100)
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
	p.EnterRule(localctx, 16, Prosr1ParserRULE_returning)

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
		p.Match(Prosr1ParserCALLS)
	}
	{
		p.SetState(103)
		p.ReturningIdent()
	}
	{
		p.SetState(104)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(105)
		p.ReturningMessageIdent()
	}
	{
		p.SetState(106)
		p.Match(Prosr1ParserT__6)
	}
	{
		p.SetState(107)
		p.Match(Prosr1ParserON)
	}
	{
		p.SetState(108)
		p.ReturningTarget()
	}
	{
		p.SetState(109)
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
	p.EnterRule(localctx, 18, Prosr1ParserRULE_returningMessageIdent)
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
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserIDENT {
		{
			p.SetState(111)
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
	p.EnterRule(localctx, 20, Prosr1ParserRULE_returningTarget)

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
	p.EnterRule(localctx, 22, Prosr1ParserRULE_message)
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
		p.SetState(116)
		p.Match(Prosr1ParserMESSAGE)
	}
	{
		p.SetState(117)
		p.MessageIdent()
	}
	{
		p.SetState(118)
		p.Match(Prosr1ParserT__3)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Prosr1ParserREPEATED)|(1<<Prosr1ParserTYPE)|(1<<Prosr1ParserIDENT))) != 0 {
		{
			p.SetState(119)
			p.Field()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
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

func (s *FieldContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserREPEATED, 0)
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
	p.EnterRule(localctx, 24, Prosr1ParserRULE_field)
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
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserREPEATED {
		{
			p.SetState(127)
			p.Match(Prosr1ParserREPEATED)
		}

	}
	{
		p.SetState(130)
		p.TypeIdent()
	}
	{
		p.SetState(131)
		p.FieldIdent()
	}
	{
		p.SetState(132)
		p.Match(Prosr1ParserT__0)
	}
	{
		p.SetState(133)
		p.Match(Prosr1ParserNUMBER)
	}
	{
		p.SetState(134)
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

func (s *TypeIdentContext) PackageIdent() IPackageIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageIdentContext)
}

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
	p.EnterRule(localctx, 26, Prosr1ParserRULE_typeIdent)

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
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Prosr1ParserIDENT:
		{
			p.SetState(136)
			p.PackageIdent()
		}
		{
			p.SetState(137)
			p.MessageIdent()
		}

	case Prosr1ParserTYPE:
		{
			p.SetState(139)
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
	p.EnterRule(localctx, 28, Prosr1ParserRULE_quote)
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
		p.SetState(142)
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
		p.SetState(144)
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_fullIdent
	return p
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(Prosr1ParserIDENT)
}

func (s *FullIdentContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (p *Prosr1Parser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Prosr1ParserRULE_fullIdent)
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
		p.SetState(146)
		p.Match(Prosr1ParserIDENT)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Prosr1ParserT__11 {
		{
			p.SetState(147)
			p.Match(Prosr1ParserT__11)
		}
		{
			p.SetState(148)
			p.Match(Prosr1ParserIDENT)
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 34, Prosr1ParserRULE_hubIdent)

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
		p.SetState(154)
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
	p.EnterRule(localctx, 36, Prosr1ParserRULE_messageIdent)

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
		p.SetState(156)
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}

// IPackageIdentContext is an interface to support dynamic dispatch.
type IPackageIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageIdentContext differentiates from other interfaces.
	IsPackageIdentContext()
}

type PackageIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageIdentContext() *PackageIdentContext {
	var p = new(PackageIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_packageIdent
	return p
}

func (*PackageIdentContext) IsPackageIdentContext() {}

func NewPackageIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageIdentContext {
	var p = new(PackageIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_packageIdent

	return p
}

func (s *PackageIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageIdentContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(Prosr1ParserIDENT)
}

func (s *PackageIdentContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, i)
}

func (s *PackageIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterPackageIdent(s)
	}
}

func (s *PackageIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitPackageIdent(s)
	}
}

func (p *Prosr1Parser) PackageIdent() (localctx IPackageIdentContext) {
	localctx = NewPackageIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Prosr1ParserRULE_packageIdent)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(158)
				p.Match(Prosr1ParserIDENT)
			}
			{
				p.SetState(159)
				p.Match(Prosr1ParserT__11)
			}

		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 40, Prosr1ParserRULE_returningIdent)

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
		p.SetState(165)
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
	p.EnterRule(localctx, 42, Prosr1ParserRULE_sendingIdent)

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
		p.SetState(167)
		p.Match(Prosr1ParserIDENT)
	}

	return localctx
}
