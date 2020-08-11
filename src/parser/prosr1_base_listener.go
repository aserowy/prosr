// Code generated from c:\Users\serow_000\source\repos\prosr\src\parser\Prosr1.g4 by ANTLR 4.8. DO NOT EDIT.

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

// EnterContent is called when production content is entered.
func (s *BaseProsr1Listener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseProsr1Listener) ExitContent(ctx *ContentContext) {}

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

// EnterSendingMessageIdent is called when production sendingMessageIdent is entered.
func (s *BaseProsr1Listener) EnterSendingMessageIdent(ctx *SendingMessageIdentContext) {}

// ExitSendingMessageIdent is called when production sendingMessageIdent is exited.
func (s *BaseProsr1Listener) ExitSendingMessageIdent(ctx *SendingMessageIdentContext) {}

// EnterSendingTarget is called when production sendingTarget is entered.
func (s *BaseProsr1Listener) EnterSendingTarget(ctx *SendingTargetContext) {}

// ExitSendingTarget is called when production sendingTarget is exited.
func (s *BaseProsr1Listener) ExitSendingTarget(ctx *SendingTargetContext) {}

// EnterReturning is called when production returning is entered.
func (s *BaseProsr1Listener) EnterReturning(ctx *ReturningContext) {}

// ExitReturning is called when production returning is exited.
func (s *BaseProsr1Listener) ExitReturning(ctx *ReturningContext) {}

// EnterReturningMessageIdent is called when production returningMessageIdent is entered.
func (s *BaseProsr1Listener) EnterReturningMessageIdent(ctx *ReturningMessageIdentContext) {}

// ExitReturningMessageIdent is called when production returningMessageIdent is exited.
func (s *BaseProsr1Listener) ExitReturningMessageIdent(ctx *ReturningMessageIdentContext) {}

// EnterReturningTarget is called when production returningTarget is entered.
func (s *BaseProsr1Listener) EnterReturningTarget(ctx *ReturningTargetContext) {}

// ExitReturningTarget is called when production returningTarget is exited.
func (s *BaseProsr1Listener) ExitReturningTarget(ctx *ReturningTargetContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseProsr1Listener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseProsr1Listener) ExitMessage(ctx *MessageContext) {}

// EnterField is called when production field is entered.
func (s *BaseProsr1Listener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseProsr1Listener) ExitField(ctx *FieldContext) {}

// EnterTypeIdent is called when production typeIdent is entered.
func (s *BaseProsr1Listener) EnterTypeIdent(ctx *TypeIdentContext) {}

// ExitTypeIdent is called when production typeIdent is exited.
func (s *BaseProsr1Listener) ExitTypeIdent(ctx *TypeIdentContext) {}

// EnterQuote is called when production quote is entered.
func (s *BaseProsr1Listener) EnterQuote(ctx *QuoteContext) {}

// ExitQuote is called when production quote is exited.
func (s *BaseProsr1Listener) ExitQuote(ctx *QuoteContext) {}

// EnterHubIdent is called when production hubIdent is entered.
func (s *BaseProsr1Listener) EnterHubIdent(ctx *HubIdentContext) {}

// ExitHubIdent is called when production hubIdent is exited.
func (s *BaseProsr1Listener) ExitHubIdent(ctx *HubIdentContext) {}

// EnterFieldIdent is called when production fieldIdent is entered.
func (s *BaseProsr1Listener) EnterFieldIdent(ctx *FieldIdentContext) {}

// ExitFieldIdent is called when production fieldIdent is exited.
func (s *BaseProsr1Listener) ExitFieldIdent(ctx *FieldIdentContext) {}

// EnterMessageIdent is called when production messageIdent is entered.
func (s *BaseProsr1Listener) EnterMessageIdent(ctx *MessageIdentContext) {}

// ExitMessageIdent is called when production messageIdent is exited.
func (s *BaseProsr1Listener) ExitMessageIdent(ctx *MessageIdentContext) {}

// EnterReturningIdent is called when production returningIdent is entered.
func (s *BaseProsr1Listener) EnterReturningIdent(ctx *ReturningIdentContext) {}

// ExitReturningIdent is called when production returningIdent is exited.
func (s *BaseProsr1Listener) ExitReturningIdent(ctx *ReturningIdentContext) {}

// EnterSendingIdent is called when production sendingIdent is entered.
func (s *BaseProsr1Listener) EnterSendingIdent(ctx *SendingIdentContext) {}

// ExitSendingIdent is called when production sendingIdent is exited.
func (s *BaseProsr1Listener) ExitSendingIdent(ctx *SendingIdentContext) {}
