package lexer

type Token struct {
	types int
	name  string
	value string
	line  *Line
}

func NewToken(types int, name string, value string) *Token {
	return &Token{
		types: types,
		value: value,
		name:  name,
	}
}

func (t *Token) Size() int {
	return len(t.value)
}

func (t *Token) SetLine(line *Line) {
	t.line = line
}

func (t *Token) Type() int {
	return t.types
}

func (t *Token) Value() string {
	return t.value
}

func (t *Token) Name() string {
	return t.name
}

func (t *Token) Line() *Line {
	return t.line
}
