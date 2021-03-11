package lexer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Lexer struct {
	rules []*Rule
	trims []int
	end   int
}

func NewLexer() *Lexer {
	return &Lexer{
		rules: []*Rule{},
	}
}

func (l *Lexer) AddTrim(trims ...int) {
	l.trims = append(l.trims, trims...)
}

func (l *Lexer) SetEnd(end int) {
	l.end = end
}

func (l *Lexer) AddRule(rule *Rule) {
	l.rules = append(l.rules, rule)
}

func (l *Lexer) getIllegalCharacterError(content []byte, cursor int) error {
	return errors.New(fmt.Sprintf("illegal character:'%v'", string(content[cursor])))
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

func (l *Lexer) Read(source *os.File) (*TokenList, error) {
	content, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}
	size := len(content)
	cursor := 0
	tokens := NewTokenList()
	tokens.SetEnd(NewToken(l.end, "EOF"))
	tokens.AddTrims(l.trims...)
	for cursor < size {
		token := l.next(content, cursor)
		if token == nil {
			return nil, l.getIllegalCharacterError(content, cursor)
		}
		cursor += token.Size()
		tokens.AddToken(token)
	}
	return tokens, nil
}
