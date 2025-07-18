package constant

const (
	ILLEGAL = iota
	EOF
	BLANK // 空格, 制表符，换行等字符

	LITERAL    // 字面量, 5, 10
	IDENTIFIER // 用户定义标识符, add, x, y
	// TODO 如何实现一个嵌套的枚举，这里是KEYWORD需要嵌套
	// 关键字
	KEYWORD

	// 运算符
	PLUS   // +
	MINUS  // -
	TIMES  // *
	DIV    // /
	ASSIGN // =
	FAC    // 阶乘!
	LT     // <
	GT     // >

	// 分割符号
	COMMA     // ,
	SEMICOLON // ;

	LPAREN   // (
	RPAREN   // )
	LBRACKET // [
	RBRACKET // ]
	LOPEN    // {
	ROPEN    // }
)

const (
	FUNCTION = iota
	DEF
	TRUE
	FALSE
	IF
	ELIF
	ELSE
	RETURN
)
