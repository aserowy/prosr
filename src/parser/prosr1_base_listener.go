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

// EnterSyntaxDefinition is called when production syntaxDefinition is entered.
func (s *BaseProsr1Listener) EnterSyntaxDefinition(ctx *SyntaxDefinitionContext) {}

// ExitSyntaxDefinition is called when production syntaxDefinition is exited.
func (s *BaseProsr1Listener) ExitSyntaxDefinition(ctx *SyntaxDefinitionContext) {}

// EnterBodyDefinition is called when production bodyDefinition is entered.
func (s *BaseProsr1Listener) EnterBodyDefinition(ctx *BodyDefinitionContext) {}

// ExitBodyDefinition is called when production bodyDefinition is exited.
func (s *BaseProsr1Listener) ExitBodyDefinition(ctx *BodyDefinitionContext) {}

// EnterPackageDefinition is called when production packageDefinition is entered.
func (s *BaseProsr1Listener) EnterPackageDefinition(ctx *PackageDefinitionContext) {}

// ExitPackageDefinition is called when production packageDefinition is exited.
func (s *BaseProsr1Listener) ExitPackageDefinition(ctx *PackageDefinitionContext) {}

// EnterHubDefinition is called when production hubDefinition is entered.
func (s *BaseProsr1Listener) EnterHubDefinition(ctx *HubDefinitionContext) {}

// ExitHubDefinition is called when production hubDefinition is exited.
func (s *BaseProsr1Listener) ExitHubDefinition(ctx *HubDefinitionContext) {}

// EnterActionDefinition is called when production actionDefinition is entered.
func (s *BaseProsr1Listener) EnterActionDefinition(ctx *ActionDefinitionContext) {}

// ExitActionDefinition is called when production actionDefinition is exited.
func (s *BaseProsr1Listener) ExitActionDefinition(ctx *ActionDefinitionContext) {}

// EnterCallsDefinition is called when production callsDefinition is entered.
func (s *BaseProsr1Listener) EnterCallsDefinition(ctx *CallsDefinitionContext) {}

// ExitCallsDefinition is called when production callsDefinition is exited.
func (s *BaseProsr1Listener) ExitCallsDefinition(ctx *CallsDefinitionContext) {}

// EnterMessageDefinition is called when production messageDefinition is entered.
func (s *BaseProsr1Listener) EnterMessageDefinition(ctx *MessageDefinitionContext) {}

// ExitMessageDefinition is called when production messageDefinition is exited.
func (s *BaseProsr1Listener) ExitMessageDefinition(ctx *MessageDefinitionContext) {}

// EnterFieldDefinition is called when production fieldDefinition is entered.
func (s *BaseProsr1Listener) EnterFieldDefinition(ctx *FieldDefinitionContext) {}

// ExitFieldDefinition is called when production fieldDefinition is exited.
func (s *BaseProsr1Listener) ExitFieldDefinition(ctx *FieldDefinitionContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *BaseProsr1Listener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *BaseProsr1Listener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterTypeIdent is called when production typeIdent is entered.
func (s *BaseProsr1Listener) EnterTypeIdent(ctx *TypeIdentContext) {}

// ExitTypeIdent is called when production typeIdent is exited.
func (s *BaseProsr1Listener) ExitTypeIdent(ctx *TypeIdentContext) {}
