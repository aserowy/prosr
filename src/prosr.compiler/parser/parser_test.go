package parser

import (
	"testing"

	"prosr.compiler/ast"
	"prosr.compiler/lexer"
)

func TestHubStatements(t *testing.T) {
	input := `
	hub SearchHub {
		action Search(SearchRequest) returns (SearchResponse) to caller;
	}
	
	hub CommentHub {
		action AddComment(CommentRequest);

		returns CommentAdded(CommentResponse) to all;
	}`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"SearchHub"},
		{"CommentHub"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testHubStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testHubStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "hub" {
		t.Errorf("s.TokenLiteral not 'hub'. got=%q", s.TokenLiteral())
		return false
	}

	hubStmt, ok := s.(*ast.HubStatement)
	if !ok {
		t.Errorf("s not *ast.HubStatement. got=%T", s)
		return false
	}

	if hubStmt.Name.Value != name {
		t.Errorf("hubStmt.Name.Value not '%s'. got=%s", name, hubStmt.Name.Value)
		return false
	}

	if hubStmt.Name.TokenLiteral() != name {
		t.Errorf("hubStmt.Name.TokenLiteral() not '%s'. got=%s", name, hubStmt.Name.TokenLiteral())
		return false
	}

	return true
}
