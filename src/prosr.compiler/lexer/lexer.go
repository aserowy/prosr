package lexer

import "prosr.compiler/token"

// Lexer which handles input and resolves token for prosr
type Lexer struct {
	input           string
	position        int
	currentPosition int
	currentChar     byte
}

// New is the ctor for Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken iterates over tokens parsed out of input
func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var tok token.Token

	switch l.currentChar {
	case '=':
		tok = newToken(token.ASSIGN, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isValidLiteral(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupTokenType(tok.Literal)
			return tok
		} else if isValidNumber(l.currentChar) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT32
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.currentPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.currentPosition]
	}
	l.position = l.currentPosition
	l.currentPosition++
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isValidLiteral(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isValidNumber(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isValidLiteral(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isValidNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
