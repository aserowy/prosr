package compiler

import (
	"strconv"

	"prosr.compiler/parser"
)

// Ast represents the whole body of the parsed tree
type Ast struct {
	nodes []Node
}

// NewAst ctor for Ast
func NewAst() *Ast {
	return new(Ast)
}

// Push adds node to toplevel
func (a *Ast) Push(n Node) {
	a.nodes = append(a.nodes, n)
}

// Node is the base structure of the compiler ast
type Node interface {
	TokenLiteral() string
}

// Hub defines the node for hubs
type Hub struct {
	Token string
	Ident string
	Nodes []Node
}

// NewHub ctor for Hub
func NewHub(ctx *parser.HubContext, s *Stack) *Hub {
	h := new(Hub)
	h.Ident = ctx.HubIdent().GetText()

	for {
		if s.PeekToken(SENDING) || s.PeekToken(RETURNING) {
			h.Nodes = append(h.Nodes, s.Pop())
		} else {
			break
		}
	}

	return h
}

// TokenLiteral returns token
func (h *Hub) TokenLiteral() string {
	return h.Token
}

// Sending defines the node for actions
type Sending struct {
	Token     string
	Ident     string
	InputType string
	Returns   Node
}

// NewSending ctor for Sending
func NewSending(ctx *parser.SendingContext) *Sending {
	s := new(Sending)
	s.Token = SENDING
	s.Ident = ctx.SendingIdent().GetText()
	s.InputType = ctx.GetInputType().GetText()

	if ctx.GetOutputType() != nil && ctx.SendingTarget() != nil {
		s.Returns = NewReturningByValues(ctx.GetOutputType().GetText(), ctx.SendingTarget().GetText())
	}

	return s
}

// TokenLiteral returns token
func (s *Sending) TokenLiteral() string {
	return s.Token
}

// Returning defines the node for returns
type Returning struct {
	Token        string
	ResponseType string
	Target       string
}

// NewReturning ctor for Returning
func NewReturning(ctx *parser.ReturningContext) *Returning {
	return NewReturningByValues(ctx.MessageIdent().GetText(), ctx.ReturningTarget().GetText())
}

// NewReturningByValues ctor for Returning
func NewReturningByValues(responseType string, target string) *Returning {
	r := new(Returning)
	r.Token = RETURNING
	r.ResponseType = responseType
	r.Target = target

	return r
}

// TokenLiteral returns token
func (r *Returning) TokenLiteral() string {
	return r.Token
}

// Message defines the node for messages
type Message struct {
	Token string
	Ident string
	Nodes []Node
}

// NewMessage ctor for Message
func NewMessage(ctx *parser.MessageContext, s *Stack) *Message {
	m := new(Message)
	m.Ident = ctx.MessageIdent().GetText()

	for {
		if s.PeekToken(FIELD) {
			m.Nodes = append(m.Nodes, s.Pop())
		} else {
			break
		}
	}

	return m
}

// TokenLiteral returns token
func (m *Message) TokenLiteral() string {
	return m.Token
}

// Field defines the node for fields
type Field struct {
	Token  string
	Type   string
	Ident  string
	Number int
}

// NewField ctor for Field
func NewField(ctx *parser.FieldContext) *Field {
	s := new(Field)
	s.Token = FIELD
	s.Type = ctx.TypeIdent().GetText()
	s.Ident = ctx.FieldIdent().GetText()

	n, err := strconv.Atoi(ctx.NUMBER().GetText())
	if err != nil {
		n = -1
	}
	s.Number = n

	return s
}

// TokenLiteral returns token
func (f *Field) TokenLiteral() string {
	return f.Token
}
