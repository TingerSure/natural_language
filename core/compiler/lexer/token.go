package lexer

type Token struct {
	types int
	value string
}

func NewToken(types int, value string) *Token {
	return &Token{
		types: types,
		value: value,
	}
}

func (t *Token) Size() int {
	return len(t.value)
}

func (t *Token) Type() int {
	return t.types
}

func (t *Token) Value() string {
	return t.value
}
