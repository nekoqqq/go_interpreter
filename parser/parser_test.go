package parser

import (
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
	if program == nil {
		t.Fatalf("解析程序失败\n")
	}
	// if len(program.) != 3 {
	// 	t.Fatalf("程序不包含3个声明, 现在有: %d\n", len(program.Statements))
	// }

	// tests := []struct {
	// 	expectedIdentifier string
	// }{
	// 	{"x"},
	// 	{"y"},
	// 	{"var"},
	// }

}
