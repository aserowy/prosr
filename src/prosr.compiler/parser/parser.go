package parser

import (
	"prosr.compiler/ast"
	"prosr.compiler/lexer"
	"prosr.compiler/token"
)

// Parser is parsing token of a lexer
type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

// New constructs a parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so currentToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses token and returns root of the AST
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
