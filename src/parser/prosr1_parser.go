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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 118,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 7, 2, 29, 10, 2, 12, 2, 14, 2, 32, 11, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 46, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 6, 6, 57, 10,
	6, 13, 6, 14, 6, 58, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 67, 10,
	7, 3, 7, 3, 7, 5, 7, 71, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 88, 10, 9, 12,
	9, 14, 9, 91, 11, 9, 3, 9, 3, 9, 3, 10, 5, 10, 96, 10, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7, 12, 109,
	10, 12, 12, 12, 14, 12, 112, 11, 12, 3, 13, 3, 13, 5, 13, 116, 10, 13,
	3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 4, 3, 2,
	16, 17, 3, 2, 19, 20, 2, 116, 2, 26, 3, 2, 2, 2, 4, 35, 3, 2, 2, 2, 6,
	45, 3, 2, 2, 2, 8, 47, 3, 2, 2, 2, 10, 51, 3, 2, 2, 2, 12, 62, 3, 2, 2,
	2, 14, 74, 3, 2, 2, 2, 16, 83, 3, 2, 2, 2, 18, 95, 3, 2, 2, 2, 20, 103,
	3, 2, 2, 2, 22, 105, 3, 2, 2, 2, 24, 115, 3, 2, 2, 2, 26, 30, 5, 4, 3,
	2, 27, 29, 5, 6, 4, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28,
	3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2,
	33, 34, 7, 2, 2, 3, 34, 3, 3, 2, 2, 2, 35, 36, 7, 3, 2, 2, 36, 37, 7, 4,
	2, 2, 37, 38, 5, 20, 11, 2, 38, 39, 7, 5, 2, 2, 39, 40, 5, 20, 11, 2, 40,
	41, 7, 6, 2, 2, 41, 5, 3, 2, 2, 2, 42, 46, 5, 8, 5, 2, 43, 46, 5, 10, 6,
	2, 44, 46, 5, 16, 9, 2, 45, 42, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 44,
	3, 2, 2, 2, 46, 7, 3, 2, 2, 2, 47, 48, 7, 7, 2, 2, 48, 49, 5, 22, 12, 2,
	49, 50, 7, 6, 2, 2, 50, 9, 3, 2, 2, 2, 51, 52, 7, 8, 2, 2, 52, 53, 7, 24,
	2, 2, 53, 56, 7, 9, 2, 2, 54, 57, 5, 12, 7, 2, 55, 57, 5, 14, 8, 2, 56,
	54, 3, 2, 2, 2, 56, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56, 3, 2, 2,
	2, 58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 7, 10, 2, 2, 61, 11,
	3, 2, 2, 2, 62, 63, 7, 11, 2, 2, 63, 64, 7, 24, 2, 2, 64, 66, 7, 12, 2,
	2, 65, 67, 5, 22, 12, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68,
	3, 2, 2, 2, 68, 70, 7, 13, 2, 2, 69, 71, 5, 14, 8, 2, 70, 69, 3, 2, 2,
	2, 70, 71, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 7, 6, 2, 2, 73, 13,
	3, 2, 2, 2, 74, 75, 7, 14, 2, 2, 75, 76, 7, 24, 2, 2, 76, 77, 7, 12, 2,
	2, 77, 78, 5, 22, 12, 2, 78, 79, 7, 13, 2, 2, 79, 80, 7, 15, 2, 2, 80,
	81, 9, 2, 2, 2, 81, 82, 7, 6, 2, 2, 82, 15, 3, 2, 2, 2, 83, 84, 7, 18,
	2, 2, 84, 85, 7, 24, 2, 2, 85, 89, 7, 9, 2, 2, 86, 88, 5, 18, 10, 2, 87,
	86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2,
	2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 7, 10, 2, 2, 93, 17,
	3, 2, 2, 2, 94, 96, 7, 22, 2, 2, 95, 94, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2,
	96, 97, 3, 2, 2, 2, 97, 98, 5, 24, 13, 2, 98, 99, 5, 22, 12, 2, 99, 100,
	7, 4, 2, 2, 100, 101, 7, 25, 2, 2, 101, 102, 7, 6, 2, 2, 102, 19, 3, 2,
	2, 2, 103, 104, 9, 3, 2, 2, 104, 21, 3, 2, 2, 2, 105, 110, 7, 24, 2, 2,
	106, 107, 7, 21, 2, 2, 107, 109, 7, 24, 2, 2, 108, 106, 3, 2, 2, 2, 109,
	112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 23, 3,
	2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 116, 5, 22, 12, 2, 114, 116, 7, 23,
	2, 2, 115, 113, 3, 2, 2, 2, 115, 114, 3, 2, 2, 2, 116, 25, 3, 2, 2, 2,
	12, 30, 45, 56, 58, 66, 70, 89, 95, 110, 115,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'syntax'", "'='", "'prosr1'", "';'", "'package'", "'hub'", "'{'",
	"'}'", "'action'", "'('", "')'", "'calls'", "'on'", "'caller'", "'all'",
	"'message'", "'''", "'\"'", "'.'", "'repeated'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "REPEATED", "TYPE", "IDENT", "NUMBER", "WHITESPACE",
}

var ruleNames = []string{
	"content", "syntax", "definition", "packageDefinition", "hubDefinition",
	"actionDefinition", "callsDefinition", "messageDefinition", "fieldDefinition",
	"quote", "fullIdent", "typeIdent",
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
	Prosr1ParserT__12      = 13
	Prosr1ParserT__13      = 14
	Prosr1ParserT__14      = 15
	Prosr1ParserT__15      = 16
	Prosr1ParserT__16      = 17
	Prosr1ParserT__17      = 18
	Prosr1ParserT__18      = 19
	Prosr1ParserREPEATED   = 20
	Prosr1ParserTYPE       = 21
	Prosr1ParserIDENT      = 22
	Prosr1ParserNUMBER     = 23
	Prosr1ParserWHITESPACE = 24
)

// Prosr1Parser rules.
const (
	Prosr1ParserRULE_content           = 0
	Prosr1ParserRULE_syntax            = 1
	Prosr1ParserRULE_definition        = 2
	Prosr1ParserRULE_packageDefinition = 3
	Prosr1ParserRULE_hubDefinition     = 4
	Prosr1ParserRULE_actionDefinition  = 5
	Prosr1ParserRULE_callsDefinition   = 6
	Prosr1ParserRULE_messageDefinition = 7
	Prosr1ParserRULE_fieldDefinition   = 8
	Prosr1ParserRULE_quote             = 9
	Prosr1ParserRULE_fullIdent         = 10
	Prosr1ParserRULE_typeIdent         = 11
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
		p.SetState(24)
		p.Syntax()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Prosr1ParserT__4)|(1<<Prosr1ParserT__5)|(1<<Prosr1ParserT__15))) != 0 {
		{
			p.SetState(25)
			p.Definition()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
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
		p.SetState(33)
		p.Match(Prosr1ParserT__0)
	}
	{
		p.SetState(34)
		p.Match(Prosr1ParserT__1)
	}
	{
		p.SetState(35)
		p.Quote()
	}
	{
		p.SetState(36)
		p.Match(Prosr1ParserT__2)
	}
	{
		p.SetState(37)
		p.Quote()
	}
	{
		p.SetState(38)
		p.Match(Prosr1ParserT__3)
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

func (s *DefinitionContext) PackageDefinition() IPackageDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageDefinitionContext)
}

func (s *DefinitionContext) HubDefinition() IHubDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHubDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHubDefinitionContext)
}

func (s *DefinitionContext) MessageDefinition() IMessageDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMessageDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMessageDefinitionContext)
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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Prosr1ParserT__4:
		{
			p.SetState(40)
			p.PackageDefinition()
		}

	case Prosr1ParserT__5:
		{
			p.SetState(41)
			p.HubDefinition()
		}

	case Prosr1ParserT__15:
		{
			p.SetState(42)
			p.MessageDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPackageDefinitionContext is an interface to support dynamic dispatch.
type IPackageDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageDefinitionContext differentiates from other interfaces.
	IsPackageDefinitionContext()
}

type PackageDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageDefinitionContext() *PackageDefinitionContext {
	var p = new(PackageDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_packageDefinition
	return p
}

func (*PackageDefinitionContext) IsPackageDefinitionContext() {}

func NewPackageDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageDefinitionContext {
	var p = new(PackageDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_packageDefinition

	return p
}

func (s *PackageDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageDefinitionContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterPackageDefinition(s)
	}
}

func (s *PackageDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitPackageDefinition(s)
	}
}

func (p *Prosr1Parser) PackageDefinition() (localctx IPackageDefinitionContext) {
	localctx = NewPackageDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Prosr1ParserRULE_packageDefinition)

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
		p.SetState(45)
		p.Match(Prosr1ParserT__4)
	}
	{
		p.SetState(46)
		p.FullIdent()
	}
	{
		p.SetState(47)
		p.Match(Prosr1ParserT__3)
	}

	return localctx
}

// IHubDefinitionContext is an interface to support dynamic dispatch.
type IHubDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHubDefinitionContext differentiates from other interfaces.
	IsHubDefinitionContext()
}

type HubDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHubDefinitionContext() *HubDefinitionContext {
	var p = new(HubDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_hubDefinition
	return p
}

func (*HubDefinitionContext) IsHubDefinitionContext() {}

func NewHubDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HubDefinitionContext {
	var p = new(HubDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_hubDefinition

	return p
}

func (s *HubDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *HubDefinitionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *HubDefinitionContext) AllActionDefinition() []IActionDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IActionDefinitionContext)(nil)).Elem())
	var tst = make([]IActionDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IActionDefinitionContext)
		}
	}

	return tst
}

func (s *HubDefinitionContext) ActionDefinition(i int) IActionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IActionDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IActionDefinitionContext)
}

func (s *HubDefinitionContext) AllCallsDefinition() []ICallsDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICallsDefinitionContext)(nil)).Elem())
	var tst = make([]ICallsDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICallsDefinitionContext)
		}
	}

	return tst
}

func (s *HubDefinitionContext) CallsDefinition(i int) ICallsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallsDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICallsDefinitionContext)
}

func (s *HubDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HubDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HubDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterHubDefinition(s)
	}
}

func (s *HubDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitHubDefinition(s)
	}
}

func (p *Prosr1Parser) HubDefinition() (localctx IHubDefinitionContext) {
	localctx = NewHubDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Prosr1ParserRULE_hubDefinition)
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
		p.SetState(49)
		p.Match(Prosr1ParserT__5)
	}
	{
		p.SetState(50)
		p.Match(Prosr1ParserIDENT)
	}
	{
		p.SetState(51)
		p.Match(Prosr1ParserT__6)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == Prosr1ParserT__8 || _la == Prosr1ParserT__11 {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Prosr1ParserT__8:
			{
				p.SetState(52)
				p.ActionDefinition()
			}

		case Prosr1ParserT__11:
			{
				p.SetState(53)
				p.CallsDefinition()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(Prosr1ParserT__7)
	}

	return localctx
}

// IActionDefinitionContext is an interface to support dynamic dispatch.
type IActionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActionDefinitionContext differentiates from other interfaces.
	IsActionDefinitionContext()
}

type ActionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionDefinitionContext() *ActionDefinitionContext {
	var p = new(ActionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_actionDefinition
	return p
}

func (*ActionDefinitionContext) IsActionDefinitionContext() {}

func NewActionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionDefinitionContext {
	var p = new(ActionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_actionDefinition

	return p
}

func (s *ActionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionDefinitionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *ActionDefinitionContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ActionDefinitionContext) CallsDefinition() ICallsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallsDefinitionContext)
}

func (s *ActionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterActionDefinition(s)
	}
}

func (s *ActionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitActionDefinition(s)
	}
}

func (p *Prosr1Parser) ActionDefinition() (localctx IActionDefinitionContext) {
	localctx = NewActionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Prosr1ParserRULE_actionDefinition)
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
		p.SetState(60)
		p.Match(Prosr1ParserT__8)
	}
	{
		p.SetState(61)
		p.Match(Prosr1ParserIDENT)
	}
	{
		p.SetState(62)
		p.Match(Prosr1ParserT__9)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserIDENT {
		{
			p.SetState(63)
			p.FullIdent()
		}

	}
	{
		p.SetState(66)
		p.Match(Prosr1ParserT__10)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserT__11 {
		{
			p.SetState(67)
			p.CallsDefinition()
		}

	}
	{
		p.SetState(70)
		p.Match(Prosr1ParserT__3)
	}

	return localctx
}

// ICallsDefinitionContext is an interface to support dynamic dispatch.
type ICallsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTarget returns the target token.
	GetTarget() antlr.Token

	// SetTarget sets the target token.
	SetTarget(antlr.Token)

	// IsCallsDefinitionContext differentiates from other interfaces.
	IsCallsDefinitionContext()
}

type CallsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	target antlr.Token
}

func NewEmptyCallsDefinitionContext() *CallsDefinitionContext {
	var p = new(CallsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_callsDefinition
	return p
}

func (*CallsDefinitionContext) IsCallsDefinitionContext() {}

func NewCallsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallsDefinitionContext {
	var p = new(CallsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_callsDefinition

	return p
}

func (s *CallsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallsDefinitionContext) GetTarget() antlr.Token { return s.target }

func (s *CallsDefinitionContext) SetTarget(v antlr.Token) { s.target = v }

func (s *CallsDefinitionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *CallsDefinitionContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *CallsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterCallsDefinition(s)
	}
}

func (s *CallsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitCallsDefinition(s)
	}
}

func (p *Prosr1Parser) CallsDefinition() (localctx ICallsDefinitionContext) {
	localctx = NewCallsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Prosr1ParserRULE_callsDefinition)
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
		p.SetState(72)
		p.Match(Prosr1ParserT__11)
	}
	{
		p.SetState(73)
		p.Match(Prosr1ParserIDENT)
	}
	{
		p.SetState(74)
		p.Match(Prosr1ParserT__9)
	}
	{
		p.SetState(75)
		p.FullIdent()
	}
	{
		p.SetState(76)
		p.Match(Prosr1ParserT__10)
	}
	{
		p.SetState(77)
		p.Match(Prosr1ParserT__12)
	}
	{
		p.SetState(78)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CallsDefinitionContext).target = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == Prosr1ParserT__13 || _la == Prosr1ParserT__14) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CallsDefinitionContext).target = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(79)
		p.Match(Prosr1ParserT__3)
	}

	return localctx
}

// IMessageDefinitionContext is an interface to support dynamic dispatch.
type IMessageDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageDefinitionContext differentiates from other interfaces.
	IsMessageDefinitionContext()
}

type MessageDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDefinitionContext() *MessageDefinitionContext {
	var p = new(MessageDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_messageDefinition
	return p
}

func (*MessageDefinitionContext) IsMessageDefinitionContext() {}

func NewMessageDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDefinitionContext {
	var p = new(MessageDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_messageDefinition

	return p
}

func (s *MessageDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDefinitionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserIDENT, 0)
}

func (s *MessageDefinitionContext) AllFieldDefinition() []IFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem())
	var tst = make([]IFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldDefinitionContext)
		}
	}

	return tst
}

func (s *MessageDefinitionContext) FieldDefinition(i int) IFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldDefinitionContext)
}

func (s *MessageDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterMessageDefinition(s)
	}
}

func (s *MessageDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitMessageDefinition(s)
	}
}

func (p *Prosr1Parser) MessageDefinition() (localctx IMessageDefinitionContext) {
	localctx = NewMessageDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Prosr1ParserRULE_messageDefinition)
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
		p.SetState(81)
		p.Match(Prosr1ParserT__15)
	}
	{
		p.SetState(82)
		p.Match(Prosr1ParserIDENT)
	}
	{
		p.SetState(83)
		p.Match(Prosr1ParserT__6)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Prosr1ParserREPEATED)|(1<<Prosr1ParserTYPE)|(1<<Prosr1ParserIDENT))) != 0 {
		{
			p.SetState(84)
			p.FieldDefinition()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(Prosr1ParserT__7)
	}

	return localctx
}

// IFieldDefinitionContext is an interface to support dynamic dispatch.
type IFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDefinitionContext differentiates from other interfaces.
	IsFieldDefinitionContext()
}

type FieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefinitionContext() *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Prosr1ParserRULE_fieldDefinition
	return p
}

func (*FieldDefinitionContext) IsFieldDefinitionContext() {}

func NewFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Prosr1ParserRULE_fieldDefinition

	return p
}

func (s *FieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefinitionContext) TypeIdent() ITypeIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeIdentContext)
}

func (s *FieldDefinitionContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *FieldDefinitionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserNUMBER, 0)
}

func (s *FieldDefinitionContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Prosr1ParserREPEATED, 0)
}

func (s *FieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.EnterFieldDefinition(s)
	}
}

func (s *FieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Prosr1Listener); ok {
		listenerT.ExitFieldDefinition(s)
	}
}

func (p *Prosr1Parser) FieldDefinition() (localctx IFieldDefinitionContext) {
	localctx = NewFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Prosr1ParserRULE_fieldDefinition)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Prosr1ParserREPEATED {
		{
			p.SetState(92)
			p.Match(Prosr1ParserREPEATED)
		}

	}
	{
		p.SetState(95)
		p.TypeIdent()
	}
	{
		p.SetState(96)
		p.FullIdent()
	}
	{
		p.SetState(97)
		p.Match(Prosr1ParserT__1)
	}
	{
		p.SetState(98)
		p.Match(Prosr1ParserNUMBER)
	}
	{
		p.SetState(99)
		p.Match(Prosr1ParserT__3)
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
	p.EnterRule(localctx, 18, Prosr1ParserRULE_quote)
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
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Prosr1ParserT__16 || _la == Prosr1ParserT__17) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 20, Prosr1ParserRULE_fullIdent)
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
		p.Match(Prosr1ParserIDENT)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Prosr1ParserT__18 {
		{
			p.SetState(104)
			p.Match(Prosr1ParserT__18)
		}
		{
			p.SetState(105)
			p.Match(Prosr1ParserIDENT)
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

func (s *TypeIdentContext) FullIdent() IFullIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
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
	p.EnterRule(localctx, 22, Prosr1ParserRULE_typeIdent)

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

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Prosr1ParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.FullIdent()
		}

	case Prosr1ParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(Prosr1ParserTYPE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
