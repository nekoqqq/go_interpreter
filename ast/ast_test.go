package ast

import (
	"fmt"
	"go_interpreter/token"
	"testing"
)

func TestAST(t *testing.T) {
	ast := &DefStatement{
		identifier: token.Token{Type: token.IDENTIFIER, LiteralValue: "x"},
		expression: nil, // 这里可以填入具体的表达式实现
	}
	var f TreeNode = ast
	fmt.Println(f)
}
