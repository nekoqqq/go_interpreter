package ast

import (
	"fmt"
	"go_interpreter/token"
)

type TreeNode interface {
	String() string
	Do()
}

type Statement interface {
	TreeNode
}
type Expression interface {
	TreeNode
}

type DefStatement struct {
	identifier token.Token
	expression Expression
}

func (ds *DefStatement) String() string {
	return ds.identifier.LiteralValue
}

func (ds *DefStatement) Do() {
	fmt.Println("DefStatement executed: ")
}
