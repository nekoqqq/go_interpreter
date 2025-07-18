package constant

const (
	ILLEGAL = iota
	EOF
	BLANK // 空格, 制表符，换行等字符

	LITERAL    // 字面量, 5, 10
	IDENTIFIER // 标识符, add, x, y

	// 运算符
	PLUS   // +
	MINUS  // -
	TIMES  // *
	DIV    // /
	ASSIGN // =

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
