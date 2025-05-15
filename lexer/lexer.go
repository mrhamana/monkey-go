package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readposition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readchar()
	return l
}

func (l *Lexer) readchar() {
	if l.readposition >= len((l.input)) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readposition]
	}
	l.position = l.readposition
	l.readposition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readchar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	//skip whitespace
	l.skipWhitespace()

	switch l.ch {

	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readchar()
			tok = token.Token{Type: token.EQ, Literal: string(l.ch) + string(ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '(':
		tok = newToken(token.LPALEN, l.ch)
	case ')':
		tok = newToken(token.RPALEN, l.ch)

	case '{':
		tok = newToken(token.LBRACE, l.ch)

	case '}':

	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {

			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// tok = newToken(token.IDENT, l.ch)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readchar()
	return tok
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readchar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readchar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peekChar() byte {
	if l.readposition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readposition]
	}
}
