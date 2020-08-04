// Code generated from c:\Users\serow_000\source\repos\prosr\src\prosr.compiler\parser\Prosr1.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Prosr1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProsr1Listener is a complete listener for a parse tree produced by Prosr1Parser.
type BaseProsr1Listener struct{}

var _ Prosr1Listener = &BaseProsr1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProsr1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProsr1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProsr1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProsr1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMain is called when production main is entered.
func (s *BaseProsr1Listener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseProsr1Listener) ExitMain(ctx *MainContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *BaseProsr1Listener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BaseProsr1Listener) ExitSyntax(ctx *SyntaxContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseProsr1Listener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseProsr1Listener) ExitDefinition(ctx *DefinitionContext) {}

// EnterHub is called when production hub is entered.
func (s *BaseProsr1Listener) EnterHub(ctx *HubContext) {}

// ExitHub is called when production hub is exited.
func (s *BaseProsr1Listener) ExitHub(ctx *HubContext) {}

// EnterSending is called when production sending is entered.
func (s *BaseProsr1Listener) EnterSending(ctx *SendingContext) {}

// ExitSending is called when production sending is exited.
func (s *BaseProsr1Listener) ExitSending(ctx *SendingContext) {}

// EnterReturning is called when production returning is entered.
func (s *BaseProsr1Listener) EnterReturning(ctx *ReturningContext) {}

// ExitReturning is called when production returning is exited.
func (s *BaseProsr1Listener) ExitReturning(ctx *ReturningContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseProsr1Listener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseProsr1Listener) ExitMessage(ctx *MessageContext) {}

// EnterField is called when production field is entered.
func (s *BaseProsr1Listener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseProsr1Listener) ExitField(ctx *FieldContext) {}

// EnterQuote is called when production quote is entered.
func (s *BaseProsr1Listener) EnterQuote(ctx *QuoteContext) {}

// ExitQuote is called when production quote is exited.
func (s *BaseProsr1Listener) ExitQuote(ctx *QuoteContext) {}

// EnterHubIdent is called when production hubIdent is entered.
func (s *BaseProsr1Listener) EnterHubIdent(ctx *HubIdentContext) {}

// ExitHubIdent is called when production hubIdent is exited.
func (s *BaseProsr1Listener) ExitHubIdent(ctx *HubIdentContext) {}

// EnterActionIdent is called when production actionIdent is entered.
func (s *BaseProsr1Listener) EnterActionIdent(ctx *ActionIdentContext) {}

// ExitActionIdent is called when production actionIdent is exited.
func (s *BaseProsr1Listener) ExitActionIdent(ctx *ActionIdentContext) {}

// EnterMessageIdent is called when production messageIdent is entered.
func (s *BaseProsr1Listener) EnterMessageIdent(ctx *MessageIdentContext) {}

// ExitMessageIdent is called when production messageIdent is exited.
func (s *BaseProsr1Listener) ExitMessageIdent(ctx *MessageIdentContext) {}

// EnterFieldIdent is called when production fieldIdent is entered.
func (s *BaseProsr1Listener) EnterFieldIdent(ctx *FieldIdentContext) {}

// ExitFieldIdent is called when production fieldIdent is exited.
func (s *BaseProsr1Listener) ExitFieldIdent(ctx *FieldIdentContext) {}
