package ast

import (
	"fmt"
	"go_interpreter/token"
	"testing"
)

func TestAST(t *testing.T) {
	ast := &DefStatement{
		Identifier: token.Token{Type: token.IDENTIFIER, LiteralValue: "x"},
		Expression: nil, // 这里可以填入具体的表达式实现
	}
	fmt.Println(ast)
}
