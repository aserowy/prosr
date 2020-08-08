package compiler

import (
	"fmt"

	"prosr.compiler/parser"
)

// Prosr1Listener implements concrete handling of parser events for compiling
type Prosr1Listener struct {
	*parser.BaseProsr1Listener

	ast   *Ast
	stack *Stack
}

// NewProsr1Listener ctor for Prosr1Listener
func NewProsr1Listener() *Prosr1Listener {
	l := new(Prosr1Listener)
	l.ast = NewAst()
	l.stack = NewStack()

	return l
}

// Ast validates and returns the parsed Ast
func (l *Prosr1Listener) Ast() *Ast {
	// TODO: Validate Ast
	return l.ast
}

// ExitHub is called when production hub is exited.
func (l *Prosr1Listener) ExitHub(ctx *parser.HubContext) {
	l.ast.Push(NewHub(ctx, l.stack))
}

// ExitSending is called when production sending is exited.
func (l *Prosr1Listener) ExitSending(ctx *parser.SendingContext) {
	l.stack.Push(NewSending(ctx))
}

// ExitReturning is called when production returning is exited.
func (l *Prosr1Listener) ExitReturning(ctx *parser.ReturningContext) {
	l.stack.Push(NewReturning(ctx))
}

// ExitMessage is called when production message is exited.
func (l *Prosr1Listener) ExitMessage(ctx *parser.MessageContext) {
	l.ast.Push(NewMessage(ctx, l.stack))
}

// ExitField is called when production field is exited.
func (l *Prosr1Listener) ExitField(ctx *parser.FieldContext) {
	l.stack.Push(NewField(ctx))
}

// ExitContent is called when production content is exited.
func (l *Prosr1Listener) ExitContent(ctx *parser.ContentContext) {
	if l.stack.Filled() {
		for l.stack.Filled() {
			n := l.stack.Pop()

			fmt.Println(n.TokenLiteral())
		}

		panic("Stack should be empty but has pending nodes!")
	}
}
