package compiler

import (
	"prosr/parser"
)

// Prosr1Listener implements concrete handling of parser events for compiling
type Prosr1Listener struct {
	*parser.BaseProsr1Listener

	packages []string
}

// NewProsr1Listener ctor for Prosr1Listener
func NewProsr1Listener() *Prosr1Listener {
	l := new(Prosr1Listener)

	return l
}

// Packages returns registered package idents
func (l *Prosr1Listener) Packages() []string {
	return l.packages
}

// EnterPackageDefinition is called when production packageDefinition is entered.
func (l *Prosr1Listener) EnterPackageDefinition(ctx *parser.PackageDefinitionContext) {
	l.packages = append(l.packages, ctx.FullIdent().GetText())
}
