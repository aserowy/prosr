package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"prosr/compiler"
	"prosr/parser"
)

var (
	path      = flag.String("p", "", "Path to .prosr file.")
	language  = flag.String("l", "csharp", "Language used for client and hub generation. The given language must have a valid .tmpl file under /compiler/.")
	namespace = flag.String("namespace", "", "Optional: Specify namespace to be able to use this value in Template. E.g. {{resolveOption .Options \"namespace\"}}")
)

func main() {
	flag.Parse()

	fp, err := filepath.Abs((*path))
	if err != nil {
		panic(err)
	}

	is, err := antlr.NewFileStream(fp)
	if err != nil {
		panic(err)
	}

	l := parser.NewProsr1Lexer(is)
	s := antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)

	p := parser.NewProsr1Parser(s)
	p.BuildParseTrees = true

	pl := compiler.NewProsr1Listener()
	antlr.ParseTreeWalkerDefault.Walk(pl, p.Content())

	b := compiler.NewBuilder("csharp", pl.Ast(), map[string]string{
		"namespace": (*namespace),
	})
	fe, out := b.Build()

	d, fn := filepath.Split(fp)
	nfn := d + strings.Replace(fn, ".prosr", *fe, -1)

	err = ioutil.WriteFile(nfn, []byte(out), 0644)
	if err != nil {
		panic(err)
	}
}
