package compiler

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/compiler/rule"
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
	for _, rule := range rule.LexerRules {
		instance.lexer.AddRule(rule)
	}
	instance.lexer.AddTrim(rule.TokenSpace)
	instance.grammar.Build()
	return instance
}

func (c *Complier) Read(source *os.File) (grammar.Phrase, error) {
	tokens, err := c.lexer.Read(source)
	if err != nil {
		return nil, err
	}
	return c.grammar.Read(tokens)
}

func (c *Complier) GetLexer() *lexer.Lexer {
	return c.lexer
}
