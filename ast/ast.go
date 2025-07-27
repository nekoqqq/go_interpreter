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
	Identifier      *token.Token
	IdentifierValue Expression
}

type RetStatement struct {
	RetToken *token.Token // Ret的token
	RetValue Expression   // Ret返回的表达式
}

func (ds *DefStatement) String() string {
	return ds.Identifier.LiteralValue
}

func (ds *DefStatement) Do() {
	fmt.Println("DefStatement executed: ")
}
