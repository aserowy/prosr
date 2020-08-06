// Code generated from c:\Users\serow_000\source\repos\prosr\src\prosr.compiler\parser\Prosr1.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Prosr1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Prosr1Listener is a complete listener for a parse tree produced by Prosr1Parser.
type Prosr1Listener interface {
	antlr.ParseTreeListener

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterHub is called when entering the hub production.
	EnterHub(c *HubContext)

	// EnterSending is called when entering the sending production.
	EnterSending(c *SendingContext)

	// EnterSendingTarget is called when entering the sendingTarget production.
	EnterSendingTarget(c *SendingTargetContext)

	// EnterReturning is called when entering the returning production.
	EnterReturning(c *ReturningContext)

	// EnterReturningTarget is called when entering the returningTarget production.
	EnterReturningTarget(c *ReturningTargetContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterTypeIdent is called when entering the typeIdent production.
	EnterTypeIdent(c *TypeIdentContext)

	// EnterQuote is called when entering the quote production.
	EnterQuote(c *QuoteContext)

	// EnterHubIdent is called when entering the hubIdent production.
	EnterHubIdent(c *HubIdentContext)

	// EnterSendingIdent is called when entering the sendingIdent production.
	EnterSendingIdent(c *SendingIdentContext)

	// EnterMessageIdent is called when entering the messageIdent production.
	EnterMessageIdent(c *MessageIdentContext)

	// EnterFieldIdent is called when entering the fieldIdent production.
	EnterFieldIdent(c *FieldIdentContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitHub is called when exiting the hub production.
	ExitHub(c *HubContext)

	// ExitSending is called when exiting the sending production.
	ExitSending(c *SendingContext)

	// ExitSendingTarget is called when exiting the sendingTarget production.
	ExitSendingTarget(c *SendingTargetContext)

	// ExitReturning is called when exiting the returning production.
	ExitReturning(c *ReturningContext)

	// ExitReturningTarget is called when exiting the returningTarget production.
	ExitReturningTarget(c *ReturningTargetContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitTypeIdent is called when exiting the typeIdent production.
	ExitTypeIdent(c *TypeIdentContext)

	// ExitQuote is called when exiting the quote production.
	ExitQuote(c *QuoteContext)

	// ExitHubIdent is called when exiting the hubIdent production.
	ExitHubIdent(c *HubIdentContext)

	// ExitSendingIdent is called when exiting the sendingIdent production.
	ExitSendingIdent(c *SendingIdentContext)

	// ExitMessageIdent is called when exiting the messageIdent production.
	ExitMessageIdent(c *MessageIdentContext)

	// ExitFieldIdent is called when exiting the fieldIdent production.
	ExitFieldIdent(c *FieldIdentContext)
}
