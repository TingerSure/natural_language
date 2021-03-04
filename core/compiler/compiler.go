package compiler

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/compiler/manager"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
)

type Complier struct {
	lexer   *lexer.Lexer
	grammar *grammar.Grammar
}

func NewComplier() *Complier {
	instance := &Complier{
		lexer:   lexer.NewLexer(),
		grammar: grammar.NewGrammar(),
	}
	for _, rule := range manager.LexerRules {
		instance.lexer.AddRule(rule)
	}
	return instance
}

func (c *Complier) Read(source *os.File) (tree.Page, error) {
	tokens, err := c.lexer.Read(source)
	if err != nil {
		return nil, err
	}
	return c.grammar.Read(c.TokenTrim(tokens))
}

func (c *Complier) TokenTrim(tokens []*lexer.Token) []*lexer.Token {
	cursor := 0
	for index, token := range tokens {
		if token.Type() == manager.Space {
			continue
		}
		tokens[cursor] = tokens[index]
		cursor++
	}
	return tokens[:cursor]
}

func (c *Complier) GetLexer() *lexer.Lexer {
	return c.lexer
}
