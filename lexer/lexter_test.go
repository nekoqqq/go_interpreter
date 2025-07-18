package lexer

import (
	"go_interpreter/constant"
	"go_interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `+-*/=,;()[]{}`
	tests := []struct {
		expectedType    token.TokenType
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
		{constant.BLANK, "\n"},

		// 第二行: def five=5
		{constant.BLANK, "\t"},
		{constant.KEYWORD, "def"},
		{constant.BLANK, " "},
		{constant.IDENTIFIER, "five"},
		{constant.ASSIGN, "="},
		{constant.LITERAL, "5"},
		{constant.SEMICOLON, ";"},
		{constant.BLANK, "\n"},

		// 第三行: def ten=10
		{constant.BLANK, "\t"},
		{constant.KEYWORD, "def"},
		{constant.BLANK, " "},
		{constant.IDENTIFIER, "ten"},
		{constant.ASSIGN, "="},
		{constant.LITERAL, "10"},
		{constant.SEMICOLON, ";"},
		{constant.BLANK, "\n"},

		// 第四行: def add  = function(x,y){
		{constant.BLANK, "\t"},
		{constant.KEYWORD, "def"},
		{constant.BLANK, " "},
		{constant.IDENTIFIER, "add"},
		{constant.BLANK, " "},
		{constant.ASSIGN, "="},
		{constant.BLANK, " "},
		{constant.KEYWORD, "func"},
		{constant.LPAREN, "("},
		{constant.IDENTIFIER, "x"},
		{constant.COMMA, ","},
		{constant.IDENTIFIER, "y"},
		{constant.RPAREN, ")"},
		{constant.LOPEN, "{"},
		{constant.BLANK, "\n"},

		// 第五行: x+y;
		{constant.BLANK, "\t"},
		{constant.BLANK, "\t"}, // 制表符
		{constant.IDENTIFIER, "x"},
		{constant.PLUS, "+"},
		{constant.IDENTIFIER, "y"},
		{constant.SEMICOLON, ";"},
		{constant.BLANK, "\n"},

		// 第六行: }
		{constant.BLANK, "\t"},
		{constant.ROPEN, "}"},
		{constant.BLANK, "\n"},

		// 第七行: def result = add(five,ten)
		{constant.BLANK, "\t"},
		{constant.KEYWORD, "def"},
		{constant.BLANK, " "},
		{constant.IDENTIFIER, "result"},
		{constant.BLANK, " "},
		{constant.ASSIGN, "="},
		{constant.BLANK, " "},
		{constant.IDENTIFIER, "add"},
		{constant.LPAREN, "("},
		{constant.IDENTIFIER, "five"},
		{constant.COMMA, ","},
		{constant.IDENTIFIER, "ten"},
		{constant.RPAREN, ")"},
		{constant.BLANK, "\n"},

		// 第八行: !/*t;
		{constant.BLANK, "\t"},
		{constant.FACT, "!"},
		{constant.DIV, "/"},
		{constant.TIMES, "*"},
		{constant.LITERAL, "5"},
		{constant.SEMICOLON, ";"},
		{constant.BLANK, "\n"},

		// 第九行: 5< 10 > 5;
		{constant.BLANK, "\t"},
		{constant.LITERAL, "5"},
		{constant.LT, "<"},
		{constant.BLANK, " "},
		{constant.LITERAL, "10"},
		{constant.GT, ">"},
		{constant.BLANK, " "},
		{constant.LITERAL, "5"},
		{constant.SEMICOLON, ";"},
		{constant.BLANK, "\n"},

		// 第十行: if(5<10){
		{constant.BLANK, "\t"},
		{constant.KEYWORD, "if"},
		{constant.LPAREN, "("},
		{constant.LITERAL, "5"},
		{constant.LT, "<"},
		{constant.LITERAL, "10"},
		{constant.RPAREN, ")"},
		{constant.LOPEN, "{"},
		{constant.BLANK, "\n"},

		// 第十一行: return true;
		{constant.BLANK, "\t"}, // 缩进制表符
		{constant.KEYWORD, "return"},
		{constant.BLANK, " "},
		{constant.KEYWORD, "true"},
		{constant.SEMICOLON, ";"},
		{constant.BLANK, "\n"},

		// 第十二行: }elif (10>20){
		{constant.BLANK, "\t"},
		{constant.ROPEN, "}"},
		{constant.KEYWORD, "elif"},
		{constant.BLANK, " "},
		{constant.LPAREN, "("},
		{constant.LITERAL, "10"},
		{constant.GT, ">"},
		{constant.LITERAL, "20"},
		{constant.RPAREN, ")"},
		{constant.LOPEN, "{"},
		{constant.BLANK, "\n"},

		// 第十三行: return false;
		{constant.BLANK, "\t"}, // 缩进制表符
		{constant.KEYWORD, "return"},
		{constant.BLANK, " "},
		{constant.KEYWORD, "false"},
		{constant.SEMICOLON, ";"},
		{constant.BLANK, "\n"},

		// 第十四行: }
		{constant.BLANK, "\t"},
		{constant.ROPEN, "}"},
		{constant.BLANK, "\n"},

		// 第十五行: else return true;
		{constant.BLANK, "\t"},
		{constant.KEYWORD, "else"},
		{constant.BLANK, " "},
		{constant.KEYWORD, "return"},
		{constant.BLANK, " "},
		{constant.KEYWORD, "true"},
		{constant.SEMICOLON, ";"},
		{constant.BLANK, "\n"},

		// 第十六行: 10 == 10 5 != 10
		{constant.BLANK, "\t"},
		{constant.LITERAL, "5"},
		{constant.BLANK, " "},
		{constant.EQ, "=="},
		{constant.BLANK, " "},
		{constant.LITERAL, "10"},
		{constant.BLANK, " "},
		{constant.LITERAL, "5"},
		{constant.BLANK, " "},
		{constant.NEQ, "!="},
		{constant.BLANK, " "},
		{constant.LITERAL, "10"},
		{constant.BLANK, "\n"},

		// 第十七行: '\t'
		{constant.BLANK, "\t"},

		// 文件结束
		{constant.EOF, ""},
	}
	l := New(input)
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
