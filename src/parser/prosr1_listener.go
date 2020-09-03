// Code generated from c:\Users\serow_000\source\repos\prosr\src\parser\Prosr1.g4 by ANTLR 4.8. DO NOT EDIT.

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

	// EnterPackageDefinition is called when entering the packageDefinition production.
	EnterPackageDefinition(c *PackageDefinitionContext)

	// EnterHubDefinition is called when entering the hubDefinition production.
	EnterHubDefinition(c *HubDefinitionContext)

	// EnterActionDefinition is called when entering the actionDefinition production.
	EnterActionDefinition(c *ActionDefinitionContext)

	// EnterCallsDefinition is called when entering the callsDefinition production.
	EnterCallsDefinition(c *CallsDefinitionContext)

	// EnterMessageDefinition is called when entering the messageDefinition production.
	EnterMessageDefinition(c *MessageDefinitionContext)

	// EnterFieldDefinition is called when entering the fieldDefinition production.
	EnterFieldDefinition(c *FieldDefinitionContext)

	// EnterQuote is called when entering the quote production.
	EnterQuote(c *QuoteContext)

	// EnterFullIdent is called when entering the fullIdent production.
	EnterFullIdent(c *FullIdentContext)

	// EnterTypeIdent is called when entering the typeIdent production.
	EnterTypeIdent(c *TypeIdentContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitSyntax is called when exiting the syntax production.
	ExitSyntax(c *SyntaxContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitPackageDefinition is called when exiting the packageDefinition production.
	ExitPackageDefinition(c *PackageDefinitionContext)

	// ExitHubDefinition is called when exiting the hubDefinition production.
	ExitHubDefinition(c *HubDefinitionContext)

	// ExitActionDefinition is called when exiting the actionDefinition production.
	ExitActionDefinition(c *ActionDefinitionContext)

	// ExitCallsDefinition is called when exiting the callsDefinition production.
	ExitCallsDefinition(c *CallsDefinitionContext)

	// ExitMessageDefinition is called when exiting the messageDefinition production.
	ExitMessageDefinition(c *MessageDefinitionContext)

	// ExitFieldDefinition is called when exiting the fieldDefinition production.
	ExitFieldDefinition(c *FieldDefinitionContext)

	// ExitQuote is called when exiting the quote production.
	ExitQuote(c *QuoteContext)

	// ExitFullIdent is called when exiting the fullIdent production.
	ExitFullIdent(c *FullIdentContext)

	// ExitTypeIdent is called when exiting the typeIdent production.
	ExitTypeIdent(c *TypeIdentContext)
}
