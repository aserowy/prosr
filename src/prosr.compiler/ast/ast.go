package ast

import "prosr.compiler/token"

// Node is the base object for the AST
type Node interface {
	TokenLiteral() string
}

// Statement is an AST object
type Statement interface {
	Node
	statementNode()
}

// Expression is an AST object
type Expression interface {
	Node
	expressionNode()
}

// Program defines the root of each AST
type Program struct {
	Statements []Statement
}

// TokenLiteral returns
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// HubStatement represents statements with keyword hub
type HubStatement struct {
	Token token.Token // the token.HUB token
	Name  *Identifier
	Value Expression
}

// TokenLiteral returns
func (ls *HubStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *HubStatement) statementNode() {}

// Identifier for statements
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// TokenLiteral returns
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

// ReturnsStatement represents statements with keyword returns
type ReturnsStatement struct {
	Token   token.Token
	Name    *Identifier
	Returns Expression
	Target  Expression
}

// TokenLiteral returns
func (rs *ReturnsStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnsStatement) statementNode() {}
