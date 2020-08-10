package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"prosr.compiler/compiler"
	"prosr.compiler/parser"
)

func main() {
	definition := `
	syntax = "prosr1";

	hub SearchHub {
		action Search(SearchRequest) returns (SearchResponse) to caller;
	  
		returns (SearchRequest) to all;
	}

	message SearchRequest {
		string query = 1;
		int32 page_number = 2;
		int32 result_per_page = 3;
	}
	
	message SearchResponse {
		Result result = 1;
	}
	
	message Result {
		string url = 1;
		string title = 2;
		string snippets = 3;
	}`

	is := antlr.NewInputStream(definition)

	l := parser.NewProsr1Lexer(is)
	s := antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)

	p := parser.NewProsr1Parser(s)
	p.BuildParseTrees = true

	pl := compiler.NewProsr1Listener()
	antlr.ParseTreeWalkerDefault.Walk(pl, p.Content())

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	b := compiler.NewBuilder(wd, pl.Ast(), map[string]string{
		"namespace": "TestNamespace.Test",
	})
	b.Build()
}
