package lexer

import (
	"go_interpreter/token"
)

type Lexer struct {
	input    string
	position int  // 下一个要读取位置
	char     byte // 当前位置的字符
}

func NewLexer(input string) *Lexer {
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

func (l *Lexer) peekChar() byte {
	if l.position >= len(l.input) {
		return 0
	}
	return l.input[l.position]
} // 向前多看一个字符

// NextToken TODO UTF-8字符支持，现在没法跳过
func (l *Lexer) NextToken() token.Token {
	var t *token.Token
	l.skipBlank()
	char := string(l.readChar())
	switch l.char {
	case '+':
		t = token.NewToken(token.PLUS, char)
	case '-':
		t = token.NewToken(token.MINUS, char)
	case '*':
		t = token.NewToken(token.TIMES, char)
	case '/':
		t = token.NewToken(token.DIV, char)
	case '=':
		t = l.makeTwoCharToken(l.char, '=', token.ASSIGN, token.EQ)
	case '!':
		t = l.makeTwoCharToken(l.char, '=', token.FACT, token.NEQ)
	case '<':
		t = token.NewToken(token.LT, char)
	case '>':
		t = token.NewToken(token.GT, char)
	case ',':
		t = token.NewToken(token.COMMA, char)
	case ';':
		t = token.NewToken(token.SEMICOLON, char)
	case '(':
		t = token.NewToken(token.LPAREN, char)
	case ')':
		t = token.NewToken(token.RPAREN, char)
	case '[':
		t = token.NewToken(token.LBRACKET, char)
	case ']':
		t = token.NewToken(token.RBRACKET, char)
	case '{':
		t = token.NewToken(token.LOPEN, char)
	case '}':
		t = token.NewToken(token.ROPEN, char)
	case 0: // NULL
		t = token.NewToken(token.EOF, "")
	case ' ', '\t', '\n', '\r':
		t = token.NewToken(token.BLANK, char)
	default:
		if isLetter(l.char) { // 读取标识符
			identifier := l.readWithStrategy(isLetter)
			tokenType := token.LookupIdentifier(identifier)
			t = token.NewToken(tokenType, identifier)
		} else if isNumber(l.char) {
			t = token.NewToken(token.LITERAL, l.readWithStrategy(isNumber))
		} else {
			t = token.NewToken(token.ILLEGAL, char)
		}
	}
	return *t
}
func (l *Lexer) skipBlank() {
	char := l.peekChar()
	for char == '\n' || char == '\t' || char == '\r' || char == ' ' {
		l.readChar()
		char = l.peekChar()
	}
}

func (l *Lexer) makeTwoCharToken(firstChar byte, secondTarget byte, singleType token.TokenType, doubleType token.TokenType) *token.Token {
	var t *token.Token
	if l.peekChar() == secondTarget {
		l.readChar()
		t = token.NewToken(doubleType, string(firstChar)+string(secondTarget))
	} else {
		t = token.NewToken(singleType, string(firstChar))
	}
	return t
}

func (l *Lexer) readWithStrategy(strategy readStrategy) string {
	begChar := l.char
	beg := l.position
	for {
		char := l.peekChar()
		if char == 0 || !strategy(char) {
			break
		}
		l.readChar()
	}
	return string(begChar) + l.input[beg:l.position]
}

type readStrategy func(byte) bool

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}
func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}
