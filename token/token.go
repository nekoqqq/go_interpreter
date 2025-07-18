package token

type Token struct {
	Type         int
	LiteralValue string
}

func NewToken(Type int, LiteralValue string) *Token {
	return &Token{
		Type:         Type,
		LiteralValue: LiteralValue,
	}
}
