package compiler

import (
	"bytes"
	"strconv"

	"prosr/parser"
)

// Ast represents the whole body of the parsed tree
type Ast struct {
	Nodes []Node
}

// NewAst ctor for Ast
func NewAst() *Ast {
	return new(Ast)
}

// Push adds node to toplevel
func (a *Ast) Push(n Node) {
	a.Nodes = append(a.Nodes, n)
}

// Node is the base structure of the compiler ast
type Node interface {
	TokenLiteral() string
	String() string
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
	h.Token = HUB
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

// String returns string representation of object
func (h *Hub) String() string {
	b := new(bytes.Buffer)
	b.WriteString(h.Token + " ")
	b.WriteString(h.Ident + "{ ")

	for _, n := range h.Nodes {
		b.WriteString(n.String() + " ")
	}

	b.WriteString("}")

	return b.String()
}

// Sending defines the node for actions
type Sending struct {
	Token     string
	Ident     string
	InputType string
	Calls     Node
}

// NewSending ctor for Sending
func NewSending(ctx *parser.SendingContext) *Sending {
	s := new(Sending)
	s.Token = SENDING
	s.Ident = ctx.SendingIdent().GetText()
	s.InputType = ctx.SendingMessageIdent().GetText()

	if ctx.ReturningIdent() != nil && ctx.SendingTarget() != nil {
		s.Calls = NewReturningByValues(
			ctx.ReturningIdent().GetText(),
			ctx.ReturningMessageIdent().GetText(),
			ctx.SendingTarget().GetText())
	}

	return s
}

// TokenLiteral returns token
func (s *Sending) TokenLiteral() string {
	return s.Token
}

// String returns string representation of object
func (s *Sending) String() string {
	b := new(bytes.Buffer)
	b.WriteString(s.Token + " ")
	b.WriteString(s.Ident + "(")
	b.WriteString(s.InputType + ") ")

	if s.Calls != nil {
		b.WriteString(s.Calls.String())
	} else {
		b.WriteString(";")
	}

	return b.String()
}

// Returning defines the node for calls
type Returning struct {
	Token        string
	Ident        string
	ResponseType string
	Target       string
}

// NewReturning ctor for Returning
func NewReturning(ctx *parser.ReturningContext) *Returning {
	return NewReturningByValues(
		ctx.ReturningIdent().GetText(),
		ctx.ReturningMessageIdent().GetText(),
		ctx.ReturningTarget().GetText())
}

// NewReturningByValues ctor for Returning
func NewReturningByValues(ident string, responseType string, target string) *Returning {
	r := new(Returning)
	r.Token = RETURNING
	r.Ident = ident
	r.ResponseType = responseType
	r.Target = target

	return r
}

// TokenLiteral returns token
func (r *Returning) TokenLiteral() string {
	return r.Token
}

// String returns string representation of object
func (r *Returning) String() string {
	b := new(bytes.Buffer)
	b.WriteString(r.Token + " ")
	b.WriteString(r.Ident + "(")
	b.WriteString(r.ResponseType + ") ")
	b.WriteString("on ")
	b.WriteString(r.Target + ";")

	return b.String()
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
	m.Token = MESSAGE
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

// String returns string representation of object
func (m *Message) String() string {
	b := new(bytes.Buffer)
	b.WriteString(m.Token + " ")
	b.WriteString(m.Ident + "{ ")

	for _, n := range m.Nodes {
		b.WriteString(n.String() + " ")
	}

	b.WriteString("}")

	return b.String()
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

// String returns string representation of object
func (f *Field) String() string {
	b := new(bytes.Buffer)
	b.WriteString(f.Type + " ")
	b.WriteString(f.Ident + " = ")
	b.WriteString(strconv.Itoa(f.Number) + ";")

	return b.String()
}
