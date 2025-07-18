package constant

const (
	ILLEGAL = iota
	EOF

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

	// 关键字
	DEF      // def
	FUNCTION // function
)
