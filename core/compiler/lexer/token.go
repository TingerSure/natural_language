package lexer

import (
	"fmt"
	"strings"
)

type Token struct {
	types   int
	name    string
	value   string
	path    string
	row     int
	col     int
	cursor  int
	content []byte
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

func (t *Token) SetContent(content []byte) {
	t.content = content
}

func (t *Token) SetCursor(cursor int) {
	t.cursor = cursor
}

func (t *Token) SetPath(path string) {
	t.path = path
}

func (t *Token) SetRow(row int) {
	t.row = row
}

func (t *Token) SetCol(col int) {
	t.col = col
}

func (t *Token) Content() []byte {
	return t.content
}

func (t *Token) Cursor() int {
	return t.cursor
}

func (t *Token) Path() string {
	return t.path
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

func (t *Token) ToLine() string {
	return t.ToLineStatic(t.path, t.row, t.col, t.cursor, t.content)
}

const (
	TokenTabCharactorWidth = 4
)

func (t *Token) ToLineStatic(path string, row int, col int, cursor int, content []byte) string {
	prev := strings.LastIndex(string(content[:cursor]), "\n") + 1
	next := strings.Index(string(content[cursor:]), "\n") + cursor
	tabSize := strings.Count(string(content[prev:cursor]), "\t")
	spaceSize := cursor - prev - tabSize + tabSize*TokenTabCharactorWidth
	return fmt.Sprintf("%v:%v:%v:\n%v\n%v^", path, row+1, col+1, string(content[prev:next]), strings.Repeat(" ", spaceSize))
}
