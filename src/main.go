package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"

	"prosr/compiler"
	"prosr/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/markbates/pkger"
)

var (
	path     = flag.String("p", "", "Path to .prosr file.")
	language = flag.String("l", "csharp", "Language used for client and hub generation. Currently supported: csharp (default)")
)

func main() {
	// Including content under templates to ensure templates are available after build
	pkger.Include("/templates/")

	flag.Parse()

	fp, err := filepath.Abs((*path))
	check(err)

	is, err := antlr.NewFileStream(fp)
	check(err)

	l := parser.NewProsr1Lexer(is)
	s := antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)

	p := parser.NewProsr1Parser(s)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	pl := compiler.NewProsr1Listener()
	c := p.Content()
	antlr.ParseTreeWalkerDefault.Walk(pl, c)

	b := compiler.NewBuilder((*language), pl.Packages(), c)
	fe, out := b.Build()

	d, fn := filepath.Split(fp)
	nfn := d + strings.Replace(fn, ".prosr", *fe, -1)

	err = ioutil.WriteFile(nfn, []byte(out), 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
