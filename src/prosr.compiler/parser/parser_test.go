package parser

import (
	"strconv"
	"testing"

	"prosr.compiler/ast"
	"prosr.compiler/lexer"
)

func TestMessageStatements(t *testing.T) {
	input := `
	message SearchRequest {
		string query = 1;
		int32 page_number = 2;
		OtherType result_per_page = 3;
	}`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}

	if program.String() != "message SearchRequest { string query = 1; int32 page_number = 2; OtherType result_per_page = 3; }" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestHubStatements(t *testing.T) {
	input := `
	hub SearchHub {
		action Search(SearchRequest) returns (SearchResponse) to caller;
	}
	
	hub CommentHub {
		action AddComment(CommentRequest);

		returns (CommentAddedResponse) to all;
	}`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier      string
		expectedSignatureLength int
	}{
		{"SearchHub", 1},
		{"CommentHub", 2},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testHubStatement(t, stmt, tt.expectedIdentifier, tt.expectedSignatureLength) {
			return
		}
	}
}

func testHubStatement(t *testing.T, s ast.Statement, name string, signatureLength int) bool {
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

	if len(hubStmt.Signature) != signatureLength {
		t.Errorf("len(hubStmt.Signature) not '%s'. got=%s", strconv.Itoa(signatureLength), strconv.Itoa(len(hubStmt.Signature)))
		return false
	}

	return true
}

func TestReturnsStatements(t *testing.T) {
	input := `
	hub Test {
		returns (CommentAddedResponse) to all;
	}`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	stmt, ok := program.Statements[0].(*ast.HubStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.HubStatement. got=%T", program.Statements[0])
	}

	if len(stmt.Signature) != 1 {
		t.Fatalf("stmt.Signature does not contain 1 statements. got=%d", len(stmt.Signature))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"CommentAddedResponse"},
	}

	for i, tt := range tests {
		stmt := stmt.Signature[i]
		if !testReturnsStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

	if program.String() != "hub Test { returns (CommentAddedResponse) to all; }" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func testReturnsStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "returns" {
		t.Errorf("s.TokenLiteral not 'returns'. got=%q", s.TokenLiteral())
		return false
	}

	_, ok := s.(*ast.ReturnsStatement)
	if !ok {
		t.Errorf("s not *ast.ReturnsStatement. got=%T", s)
		return false
	}

	return true
}

func TestActionStatements(t *testing.T) {
	input := `
	hub Test {
		action TestI(TestTypeI) returns (TestTypeII) to all;
		action TestII(TestTypeII);
	}`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	stmt, ok := program.Statements[0].(*ast.HubStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.HubStatement. got=%T", program.Statements[0])
	}

	if len(stmt.Signature) != 2 {
		t.Fatalf("stmt.Signature does not contain 2 statements. got=%d", len(stmt.Signature))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"TestI"},
		{"TestII"},
	}

	for i, tt := range tests {
		stmt := stmt.Signature[i]
		if !testActionStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

	if program.String() != "hub Test { action TestI(TestTypeI) returns (TestTypeII) to all; action TestII(TestTypeII); }" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func testActionStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "action" {
		t.Errorf("s.TokenLiteral not 'returns'. got=%q", s.TokenLiteral())
		return false
	}

	_, ok := s.(*ast.ActionStatement)
	if !ok {
		t.Errorf("s not *ast.ActionStatement. got=%T", s)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
