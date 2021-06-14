package lexer

import (
	"fmt"
	"strings"
)

type Lexer struct {
	rules []*Rule
	trims []int
	eof   func() *Token
}

func NewLexer() *Lexer {
	return &Lexer{
		rules: []*Rule{},
	}
}

func (l *Lexer) AddTrim(trims ...int) {
	l.trims = append(l.trims, trims...)
}

func (l *Lexer) SetEof(eof func() *Token) {
	l.eof = eof
}

func (l *Lexer) AddRule(rule *Rule) {
	l.rules = append(l.rules, rule)
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

func (l *Lexer) Read(content []byte, path string) (*TokenList, error) {
	size := len(content)
	cursor := 0
	row := 0
	col := 0
	tokens := NewTokenList()
	tokens.AddTrims(l.trims...)
	startLine := NewLine(path, content)
	startLine.SetStart(0, 0, 0)
	startLine.SetEnd(0, 0, 0)
	tokens.SetStartLine(startLine)
	for cursor < size {
		token := l.next(content, cursor)
		line := NewLine(path, content)
		line.SetStart(row, col, cursor)
		if token == nil {
			return nil, fmt.Errorf("invalid character: '%v'\n%v", string(content[cursor]), line.ToString())
		}
		row, col = l.updateRowCol(row, col, token)
		cursor += token.Size()
		line.SetEnd(row, col, cursor)
		token.SetLine(line)
		tokens.AddToken(token)
	}
	eofLine := NewLine(path, content)
	eofLine.SetStart(row, col, cursor)
	eofLine.SetEnd(row, col, cursor)
	eof := l.eof()
	eof.SetLine(eofLine)
	tokens.SetEof(eof)
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
