package parser

import (
	"go_interpreter/ast"
	"go_interpreter/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	input := `
	def x=5;
	def y=10;
	def var=996;
	`
	l := lexer.NewLexer(input)
	parser := NewParser(l)
	program := parser.Parse()
	checkParserErrors(t, parser)
	if program == nil {
		t.Fatalf("解析程序失败\n")
	}
	if len(program) != 3 {
		t.Fatalf("程序不包含3个声明, 现在有: %d个声明\n", len(program))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"var"},
	}
	for i, test := range tests {
		statement := program[i]
		if statement.String() != "def" {
			t.Errorf("当前statement的token string不是def, 当前是: %q\n", statement.String())
			continue
		}
		defStatement, ok := statement.(*ast.DefStatement)
		if !ok {
			t.Errorf("当前的statement类型不是def Statement, 当前类型是: %T\n", statement)
			continue
		}
		if defStatement.Identifier.LiteralValue != test.expectedIdentifier {
			t.Errorf("当前的statement的字符串名称不一致, 当前是: %s\n", defStatement.Identifier.LiteralValue)
			continue
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errros := p.errors
	if len(errros) == 0 {
		return
	}
	for i, errMsg := range p.errors {
		t.Errorf("[%d]解析失败, 失败原因: %q", i+1, errMsg)
	}
	t.FailNow()
}
