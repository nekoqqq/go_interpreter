package lexer

import (
	"go_interpreter/constant"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `+-*/=,;()[]{}`
	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{constant.PLUS, "+"},
		{constant.MINUS, "-"},
		{constant.TIMES, "*"},
		{constant.DIV, "/"},
		{constant.ASSIGN, "="},
		{constant.COMMA, ","},
		{constant.SEMICOLON, ";"},
		{constant.LPAREN, "("},
		{constant.RPAREN, ")"},
		{constant.LBRACKET, "["},
		{constant.RBRACKET, "]"},
		{constant.LOPEN, "{"},
		{constant.ROPEN, "}"},
		{constant.EOF, ""},
	}
	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, test.expectedType, tok.Type)
		}
		if tok.LiteralValue != test.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, test.expectedLiteral, tok.LiteralValue)
		}
	}

}
