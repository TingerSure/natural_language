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
	instance.lexer.AddTrim(rule.TypeSpace)
	instance.lexer.SetEnd(rule.LexerEnd)
	for _, rule := range rule.GrammarRules {
		instance.grammar.AddRule(rule)
	}
	instance.grammar.SetAccept(rule.SymbolEnd)
	instance.grammar.SetGlobal(rule.SymbolPage)
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

func (c *Complier) GetGrammar() *grammar.Grammar {
	return c.grammar
}
