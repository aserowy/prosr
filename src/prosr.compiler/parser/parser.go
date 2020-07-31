package parser

import (
	"fmt"

	"prosr.compiler/ast"
	"prosr.compiler/lexer"
	"prosr.compiler/token"
)

// Parser is parsing token of a lexer
type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

// New constructs a parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so currentToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram parses token and returns root of the AST
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// Errors returns all found errors while building the AST
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.HUB:
		return p.parseHubStatement()
	default:
		return nil
	}
}

func (p *Parser) parseHubStatement() *ast.HubStatement {
	stmt := &ast.HubStatement{Token: p.currentToken}

	if !p.expectPeekAndAdvance(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	if !p.expectPeekAndAdvance(token.LBRACE) {
		return nil
	}

	// TODO: We're skipping the expressions until }
	for !p.currentTokenIs(token.RBRACE) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeekAndAdvance(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
