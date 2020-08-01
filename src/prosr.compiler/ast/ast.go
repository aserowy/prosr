package ast

import (
	"bytes"

	"prosr.compiler/token"
)

// Node is the base object for the AST
type Node interface {
	TokenLiteral() string
	String() string
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

// String returns a string representation of this node
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// HubStatement represents statements with keyword hub
type HubStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral returns
func (hs *HubStatement) TokenLiteral() string {
	return hs.Token.Literal
}

// String returns a string representation of this node
func (hs *HubStatement) String() string {
	var out bytes.Buffer

	out.WriteString(hs.TokenLiteral() + " ")
	out.WriteString(hs.Name.String())
	out.WriteString(" { ")

	if hs.Value != nil {
		out.WriteString(hs.Value.String())
	}

	out.WriteString(" }")

	return out.String()
}

func (hs *HubStatement) statementNode() {}

// ActionStatement represents statements with keyword action
type ActionStatement struct {
	Token     token.Token
	Name      *Identifier
	InputType Expression
	Returns   Statement
}

// TokenLiteral returns
func (rs *ActionStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns a string representation of this node
func (rs *ActionStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	out.WriteString(rs.Name.String())
	out.WriteString("(" + rs.InputType.String() + ")")

	if rs.Returns != nil {
		out.WriteString(" " + rs.Returns.String())
	} else {
		out.WriteString(";")
	}

	return out.String()
}

func (rs *ActionStatement) statementNode() {}

// ReturnsStatement represents statements with keyword returns
type ReturnsStatement struct {
	Token      token.Token
	Name       *Identifier
	OutputType Expression
	Target     Expression
}

// TokenLiteral returns
func (rs *ReturnsStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns a string representation of this node
func (rs *ReturnsStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	out.WriteString("(" + rs.OutputType.String() + ") ")
	out.WriteString("to ")
	out.WriteString(rs.Target.String())
	out.WriteString(";")

	return out.String()
}

func (rs *ReturnsStatement) statementNode() {}

// TargetExpression represents a message type
type TargetExpression struct {
	Token token.Token // Literal
	Name  *Identifier
}

// TokenLiteral returns
func (te *TargetExpression) TokenLiteral() string {
	return te.Token.Literal
}

// String returns a string representation of this node
func (te *TargetExpression) String() string {
	return te.Name.String()
}

func (te *TargetExpression) expressionNode() {}

// TypeExpression represents a message type
type TypeExpression struct {
	Token token.Token // Literal
	Name  *Identifier
}

// TokenLiteral returns
func (te *TypeExpression) TokenLiteral() string {
	return te.Token.Literal
}

// String returns a string representation of this node
func (te *TypeExpression) String() string {
	return te.Name.String()
}

func (te *TypeExpression) expressionNode() {}

// Identifier for statements
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral returns
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String returns a string representation of this node
func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}
