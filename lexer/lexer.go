package lexer

import (
	"go_interpreter/constant"
	"go_interpreter/token"
)

type Lexer struct {
	input    string
	position int  // 当前读取位置
	char     byte // 当前读取的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) ReadChar() byte {
	if l.position >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.position]
		l.position++
	}
	return l.char
}

func (l *Lexer) NextToken() token.Token {
	var t *token.Token
	char := string(l.ReadChar())
	switch l.char {
	case '+':
		t = token.NewToken(constant.PLUS, char)
	case '-':
		t = token.NewToken(constant.MINUS, char)
	case '*':
		t = token.NewToken(constant.TIMES, char)
	case '/':
		t = token.NewToken(constant.DIV, char)
	case '=':
		t = token.NewToken(constant.ASSIGN, char)
	case ',':
		t = token.NewToken(constant.COMMA, char)
	case ';':
		t = token.NewToken(constant.SEMICOLON, char)
	case '(':
		t = token.NewToken(constant.LPAREN, char)
	case ')':
		t = token.NewToken(constant.RPAREN, char)
	case '[':
		t = token.NewToken(constant.LBRACKET, char)
	case ']':
		t = token.NewToken(constant.RBRACKET, char)
	case '{':
		t = token.NewToken(constant.LOPEN, char)
	case '}':
		t = token.NewToken(constant.ROPEN, char)
	case 0: // NULL
		t = token.NewToken(constant.EOF, "")
	default:
		t = &token.Token{Type: constant.ILLEGAL, LiteralValue: char}
	}
	return *t
}
