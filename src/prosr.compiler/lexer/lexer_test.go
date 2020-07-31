package lexer

import (
	"testing"

	"prosr.compiler/token"
)

func TestNextToken(t *testing.T) {
	input := `
	hub SearchHub {
		action Search(SearchRequest) returns (SearchResponse) to caller;
	}
	
	message SearchResponse {
		Result result = 1;
	}
	  
	message Result {
		string url = 1;
		string title = 2;
	}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.HUB, "hub"},
		{token.IDENT, "SearchHub"},
		{token.LBRACE, "{"},
		{token.ACTION, "action"},
		{token.IDENT, "Search"},
		{token.LPAREN, "("},
		{token.IDENT, "SearchRequest"},
		{token.RPAREN, ")"},
		{token.RETURNS, "returns"},
		{token.LPAREN, "("},
		{token.IDENT, "SearchResponse"},
		{token.RPAREN, ")"},
		{token.TO, "to"},
		{token.IDENT, "caller"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.MESSAGE, "message"},
		{token.IDENT, "SearchResponse"},
		{token.LBRACE, "{"},
		{token.IDENT, "Result"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.INT32, "1"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.MESSAGE, "message"},
		{token.IDENT, "Result"},
		{token.LBRACE, "{"},
		{token.IDENT, "string"},
		{token.IDENT, "url"},
		{token.ASSIGN, "="},
		{token.INT32, "1"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "string"},
		{token.IDENT, "title"},
		{token.ASSIGN, "="},
		{token.INT32, "2"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
