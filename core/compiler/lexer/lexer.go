package lexer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Lexer struct {
	rules []*Rule
	trims []int
	end   *Token
}

func NewLexer() *Lexer {
	return &Lexer{
		rules: []*Rule{},
	}
}

func (l *Lexer) AddTrim(trims ...int) {
	l.trims = append(l.trims, trims...)
}

func (l *Lexer) SetEnd(end *Token) {
	l.end = end
}

func (l *Lexer) AddRule(rule *Rule) {
	l.rules = append(l.rules, rule)
}

func (l *Lexer) getIllegalCharacterError(content []byte, cursor int, path string, row, col int) error {
	prev := strings.LastIndex(string(content[:cursor]), "\n") + 1
	next := strings.Index(string(content[cursor:]), "\n") + cursor
	line := string(content[prev:next])
	return errors.New(fmt.Sprintf(
		"invalid character: '%v'\n%v:%v:%v: \n%v\n%v^",
		string(content[cursor]),
		path,
		row+1,
		col+1,
		line,
		strings.Repeat(" ", cursor-prev),
	))
}

func (l *Lexer) next(content []byte, cursor int) *Token {
	for _, rule := range l.rules {
		value := rule.GetMatcher().Find(content[cursor:])
		if value == nil {
			continue
		}
		return rule.CreateToken(value)
	}
	return nil
}

func (l *Lexer) Read(source *os.File, path string) (*TokenList, error) {
	content, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}
	size := len(content)
	cursor := 0
	row := 0
	col := 0
	tokens := NewTokenList()
	tokens.SetEnd(l.end)
	tokens.AddTrims(l.trims...)
	for cursor < size {
		token := l.next(content, cursor)
		if token == nil {
			return nil, l.getIllegalCharacterError(content, cursor, path, row, col)
		}
		token.SetRow(row)
		token.SetCol(col)
		row, col = l.updateRowCol(row, col, token)
		cursor += token.Size()
		tokens.AddToken(token)
	}
	return tokens, nil
}

func (l *Lexer) updateRowCol(row, col int, token *Token) (int, int) {
	last := strings.LastIndex(token.Value(), "\n")
	if last < 0 {
		col += token.Size()
		return row, col
	}
	count := strings.Count(token.Value(), "\n")
	row += count
	col = token.Size() - last - 1
	return row, col
}
