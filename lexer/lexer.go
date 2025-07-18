package lexer

import (
	"go_interpreter/constant"
	"go_interpreter/token"
)

type Lexer struct {
	input    string
	position int  // 下一个要读取位置
	char     byte // 当前读取的字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() byte {
	if l.position >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.position]
		l.position++
	}
	return l.char
}

func (l *Lexer) rollbackChar() { // 回退一个字符
	if l.position <= 0 {
		return
	}
	l.position--
}

func (l *Lexer) NextToken() token.Token {
	var t *token.Token
	char := string(l.readChar())
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
	case ' ', '\t', '\n', '\r':
		t = token.NewToken(constant.BLANK, char)
	default:
		if isLetter(l.char) { // 读取标识符
			t = token.NewToken(constant.IDENTIFIER, l.readIdentifier())
		} else if isNumber(l.char) {
			t = token.NewToken(constant.LITERAL, l.readNumber())
		} else {
			t = token.NewToken(constant.ILLEGAL, char)
		}
	}
	return *t
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}
func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}

// TODO 优化下面的代码，有重复部分
func (l *Lexer) readIdentifier() string {
	l.rollbackChar()
	beg := l.position
	for isLetter(l.readChar()) {
	}
	l.rollbackChar()
	return l.input[beg:l.position]
}

func (l *Lexer) readNumber() string {
	l.rollbackChar()
	beg := l.position
	for isNumber(l.readChar()) {
	}
	l.rollbackChar()
	return l.input[beg:l.position]
}
