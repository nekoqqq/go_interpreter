package lexer

import (
	"go_interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `+-*/=,;()[]{}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.TIMES, "*"},
		{token.DIV, "/"},
		{token.ASSIGN, "="},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.LOPEN, "{"},
		{token.ROPEN, "}"},
		{token.EOF, ""},
	}
	l := NewLexer(input)
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
func TestSourceCode(t *testing.T) {
	input := `
	def five=5;
	def ten=10;
	def add = func(x,y){
		x+y;
	}
	def result = add(five,ten)
	!/*5;
	5< 10> 5;
	if(5<10){
	return true;
	}elif (10>20){
	return false;
	}
	else return true;
	5 == 10 5 != 10
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// 第一行: 空行（只有换行符）
		{token.BLANK, "\n"},

		// 第二行: def five=5
		{token.BLANK, "\t"},
		{token.KEYWORD, "def"},
		{token.BLANK, " "},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.LITERAL, "5"},
		{token.SEMICOLON, ";"},
		{token.BLANK, "\n"},

		// 第三行: def ten=10
		{token.BLANK, "\t"},
		{token.KEYWORD, "def"},
		{token.BLANK, " "},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.LITERAL, "10"},
		{token.SEMICOLON, ";"},
		{token.BLANK, "\n"},

		// 第四行: def add  = function(x,y){
		{token.BLANK, "\t"},
		{token.KEYWORD, "def"},
		{token.BLANK, " "},
		{token.IDENTIFIER, "add"},
		{token.BLANK, " "},
		{token.ASSIGN, "="},
		{token.BLANK, " "},
		{token.KEYWORD, "func"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LOPEN, "{"},
		{token.BLANK, "\n"},

		// 第五行: x+y;
		{token.BLANK, "\t"},
		{token.BLANK, "\t"}, // 制表符
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.BLANK, "\n"},

		// 第六行: }
		{token.BLANK, "\t"},
		{token.ROPEN, "}"},
		{token.BLANK, "\n"},

		// 第七行: def result = add(five,ten)
		{token.BLANK, "\t"},
		{token.KEYWORD, "def"},
		{token.BLANK, " "},
		{token.IDENTIFIER, "result"},
		{token.BLANK, " "},
		{token.ASSIGN, "="},
		{token.BLANK, " "},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.BLANK, "\n"},

		// 第八行: !/*t;
		{token.BLANK, "\t"},
		{token.FACT, "!"},
		{token.DIV, "/"},
		{token.TIMES, "*"},
		{token.LITERAL, "5"},
		{token.SEMICOLON, ";"},
		{token.BLANK, "\n"},

		// 第九行: 5< 10 > 5;
		{token.BLANK, "\t"},
		{token.LITERAL, "5"},
		{token.LT, "<"},
		{token.BLANK, " "},
		{token.LITERAL, "10"},
		{token.GT, ">"},
		{token.BLANK, " "},
		{token.LITERAL, "5"},
		{token.SEMICOLON, ";"},
		{token.BLANK, "\n"},

		// 第十行: if(5<10){
		{token.BLANK, "\t"},
		{token.KEYWORD, "if"},
		{token.LPAREN, "("},
		{token.LITERAL, "5"},
		{token.LT, "<"},
		{token.LITERAL, "10"},
		{token.RPAREN, ")"},
		{token.LOPEN, "{"},
		{token.BLANK, "\n"},

		// 第十一行: return true;
		{token.BLANK, "\t"}, // 缩进制表符
		{token.KEYWORD, "return"},
		{token.BLANK, " "},
		{token.KEYWORD, "true"},
		{token.SEMICOLON, ";"},
		{token.BLANK, "\n"},

		// 第十二行: }elif (10>20){
		{token.BLANK, "\t"},
		{token.ROPEN, "}"},
		{token.KEYWORD, "elif"},
		{token.BLANK, " "},
		{token.LPAREN, "("},
		{token.LITERAL, "10"},
		{token.GT, ">"},
		{token.LITERAL, "20"},
		{token.RPAREN, ")"},
		{token.LOPEN, "{"},
		{token.BLANK, "\n"},

		// 第十三行: return false;
		{token.BLANK, "\t"}, // 缩进制表符
		{token.KEYWORD, "return"},
		{token.BLANK, " "},
		{token.KEYWORD, "false"},
		{token.SEMICOLON, ";"},
		{token.BLANK, "\n"},

		// 第十四行: }
		{token.BLANK, "\t"},
		{token.ROPEN, "}"},
		{token.BLANK, "\n"},

		// 第十五行: else return true;
		{token.BLANK, "\t"},
		{token.KEYWORD, "else"},
		{token.BLANK, " "},
		{token.KEYWORD, "return"},
		{token.BLANK, " "},
		{token.KEYWORD, "true"},
		{token.SEMICOLON, ";"},
		{token.BLANK, "\n"},

		// 第十六行: 10 == 10 5 != 10
		{token.BLANK, "\t"},
		{token.LITERAL, "5"},
		{token.BLANK, " "},
		{token.EQ, "=="},
		{token.BLANK, " "},
		{token.LITERAL, "10"},
		{token.BLANK, " "},
		{token.LITERAL, "5"},
		{token.BLANK, " "},
		{token.NEQ, "!="},
		{token.BLANK, " "},
		{token.LITERAL, "10"},
		{token.BLANK, "\n"},

		// 第十七行: '\t'
		{token.BLANK, "\t"},

		// 文件结束
		{token.EOF, ""},
	}
	l := NewLexer(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.expectedType {
			t.Fatalf("test: [%d] - token: %v, tokentype wrong. expected=%q, got=%q", i, tok, test.expectedType, tok.Type)
		}
		if tok.LiteralValue != test.expectedLiteral {
			t.Fatalf("test[%d] - token: %v, literal wrong. expected=%q, got=%q", i, tok, test.expectedLiteral, tok.LiteralValue)
		}
	}
}
