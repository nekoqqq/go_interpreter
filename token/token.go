package token

import "go_interpreter/constant"

type TokenType int
type Token struct {
	Type         TokenType
	LiteralValue string
}

func NewToken(Type TokenType, LiteralValue string) *Token {
	return &Token{
		Type:         Type,
		LiteralValue: LiteralValue,
	}
}

var keywords = map[string]TokenType{
	"func":   constant.FUNCTION,
	"def":    constant.DEF,
	"true":   constant.TRUE,
	"false":  constant.FALSE,
	"if":     constant.IF,
	"elif":   constant.ELIF,
	"else":   constant.ELIF,
	"return": constant.RETURN,
}

func LookupIdentifier(identifier string) TokenType {
	if _, ok := keywords[identifier]; ok {
		return constant.KEYWORD
	}
	return constant.IDENTIFIER
}
