package lexer

type Token struct {
	types int
	name  string
	value string
	row   int
	col   int
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

func (t *Token) SetRow(row int) {
	t.row = row
}

func (t *Token) SetCol(col int) {
	t.col = col
}

func (t *Token) Row() int {
	return t.row
}

func (t *Token) Col() int {
	return t.col
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
