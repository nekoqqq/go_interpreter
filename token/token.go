package token

type Token struct {
	Type         TokenType
	LiteralValue string
}

type TokenType string

const (
	ILLEGAL TokenType = "illegal"
	EOF               = "eof"
	BLANK             = "blank" // 空格, 制表符，换行等字符

	LITERAL    = "literal"    // 字面量, 5, 10
	IDENTIFIER = "identifier" // 用户定义标识符, add, x, y

	// 运算符
	PLUS   = "operator" // +
	MINUS               // -
	TIMES               // *
	DIV                 // /
	ASSIGN              // =
	FACT                // 阶乘!
	LT                  // <
	GT                  // >
	EQ                  // ==
	NEQ                 // !=
	// 分割符号
	COMMA     = "separator" // ,
	SEMICOLON               // ;

	LPAREN   = "brace" // (
	RPAREN             // )
	LBRACKET           // [
	RBRACKET           // ]
	LOPEN              // {
	ROPEN              // }
	// 关键字
	// TODO 如何实现一个嵌套的枚举，这里是KEYWORD需要嵌套，当前一版是将keyword放在最后
	KEYWORD = "keyword"
)

const (
	FUNCTION = "function"
	DEF      = "def"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELIF     = "elif"
	ELSE     = "else"
	RETURN   = "return"
)

func NewToken(Type TokenType, LiteralValue string) *Token {
	return &Token{
		Type:         Type,
		LiteralValue: LiteralValue,
	}
}

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"def":    DEF,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"elif":   ELIF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifier(identifier string) TokenType {
	if _, ok := keywords[identifier]; ok {
		return KEYWORD
	}
	return IDENTIFIER
}
