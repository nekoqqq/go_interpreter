package parser

import (
	"go_interpreter/ast"
	"go_interpreter/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}

func (p *Parser) Parse() *ast.TreeNode {
	return nil
}
