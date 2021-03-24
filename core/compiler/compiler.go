package compiler

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/compiler/rule"
	"github.com/TingerSure/natural_language/core/compiler/semantic"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
)

type Complier struct {
	lexer    *lexer.Lexer
	grammar  *grammar.Grammar
	semantic *semantic.Semantic
	libs     *tree.LibraryManager
}

func NewComplier(libs *tree.LibraryManager) *Complier {
	instance := &Complier{
		lexer:    lexer.NewLexer(),
		grammar:  grammar.NewGrammar(),
		semantic: semantic.NewSemantic(libs),
		libs:     libs,
	}
	for _, rule := range rule.LexerRules {
		instance.lexer.AddRule(rule)
	}
	instance.lexer.AddTrim(rule.LexerTrim...)
	instance.lexer.SetEnd(rule.LexerEnd)
	for _, rule := range rule.GrammarRules {
		instance.grammar.AddRule(rule)
	}
	instance.grammar.SetEnd(rule.GrammarEnd)
	instance.grammar.SetGlobal(rule.GrammarGlobal)
	instance.grammar.Build()
	for _, rule := range rule.SemanticRules {
		instance.semantic.AddRule(rule)
	}
	return instance
}

func (c *Complier) Read(path string) error {
	source, err := os.Open(path)
	if err != nil {
		return err
	}
	tokens, err := c.lexer.Read(source)
	if err != nil {
		return err
	}
	phrase, err := c.grammar.Read(tokens)
	if err != nil {
		return err
	}
	page, err := c.semantic.Read(phrase)
	if err != nil {
		return err
	}
	lib := tree.NewLibraryAdaptor()
	lib.SetPage(page.GetName(), page)
	c.libs.AddLibrary("test", lib)

	return nil
}

func (c *Complier) GetLexer() *lexer.Lexer {
	return c.lexer
}

func (c *Complier) GetGrammar() *grammar.Grammar {
	return c.grammar
}

func (c *Complier) GetSemantic() *semantic.Semantic {
	return c.semantic
}
