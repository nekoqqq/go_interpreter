package parser

import (
	"fmt"
	"go_interpreter/ast"
	"go_interpreter/lexer"
	"go_interpreter/token"
)

type Program []ast.Statement // 一个完整的程序是由多个申明构成
type Parser struct {
	l        *lexer.Lexer
	curToken token.Token     // 当前待读取的token
	program  []ast.Statement // TODO interface不需要用指针吗？
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		l: lexer,
	}
}
func (p *Parser) Parse() Program {
	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			p.program = append(p.program, statement)
		}
	}

	return nil
}

func (p *Parser) parseStatement() ast.Statement {
	curToken := p.l.NextToken()
	statement := &ast.DefStatement{}
	// parseDef
	if curToken.Type == token.DEF {
		identifierToken := p.l.NextToken()
		if identifierToken.Type != token.IDENTIFIER {
			fmt.Printf("解析statement的时候[标识符]无法正常解析")
			return nil
		}
		statement.Identifier = curToken

		assginToken := p.l.NextToken()
		if assginToken.Type != token.ASSIGN {
			fmt.Printf("解析statement的时候[赋值符号]无法正常解析")
			return nil
		}

		// TODO 暂时先跳过对表达式的解析
		for curToken.Type != token.SEMICOLON {
			curToken = p.l.NextToken()
		}

	}
	return statement
}

func (p *Parser) advanceToken() {

}
