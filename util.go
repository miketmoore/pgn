package pgn

func NewToken(t TokenType, v string) Token {
	return Token{Type: t, Value: v}
}
