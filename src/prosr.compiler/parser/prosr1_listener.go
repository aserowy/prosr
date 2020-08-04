// Code generated from c:\Users\serow_000\source\repos\prosr\src\prosr.compiler\parser\Prosr1.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Prosr1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Prosr1Listener is a complete listener for a parse tree produced by Prosr1Parser.
type Prosr1Listener interface {
	antlr.ParseTreeListener

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterSyntax is called when entering the syntax production.
	EnterSyntax(c *SyntaxContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterHub is called when entering the hub production.
	EnterHub(c *HubContext)

	// EnterSending is called when entering the sending production.
	EnterSending(c *SendingContext)

	// EnterReturning is called when entering the returning production.
	EnterReturning(c *ReturningContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterQuote is called when entering the quote production.
	EnterQuote(c *QuoteContext)

	// EnterHubIdent is called when entering the hubIdent production.
	EnterHubIdent(c *HubIdentContext)

	// EnterActionIdent is called when entering the actionIdent production.
	EnterActionIdent(c *ActionIdentContext)

	// EnterMessageIdent is called when entering the messageIdent production.
	EnterMessageIdent(c *MessageIdentContext)

	// EnterFieldIdent is called when entering the fieldIdent production.
	EnterFieldIdent(c *FieldIdentContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitHub is called when exiting the hub production.
	ExitHub(c *HubContext)

	// ExitSending is called when exiting the sending production.
	ExitSending(c *SendingContext)

	// ExitReturning is called when exiting the returning production.
	ExitReturning(c *ReturningContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitQuote is called when exiting the quote production.
	ExitQuote(c *QuoteContext)

	// ExitHubIdent is called when exiting the hubIdent production.
	ExitHubIdent(c *HubIdentContext)

	// ExitActionIdent is called when exiting the actionIdent production.
	ExitActionIdent(c *ActionIdentContext)

	// ExitMessageIdent is called when exiting the messageIdent production.
	ExitMessageIdent(c *MessageIdentContext)

	// ExitFieldIdent is called when exiting the fieldIdent production.
	ExitFieldIdent(c *FieldIdentContext)
}
