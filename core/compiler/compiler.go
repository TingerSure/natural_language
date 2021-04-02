package compiler

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/compiler/rule"
	"github.com/TingerSure/natural_language/core/compiler/semantic"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
	"strings"
)

type Complier struct {
	lexer    *lexer.Lexer
	grammar  *grammar.Grammar
	semantic *semantic.Semantic
	context  *semantic.Context
	libs     *tree.LibraryManager
	roots    []string
}

func NewComplier(libs *tree.LibraryManager) *Complier {
	instance := &Complier{
		lexer:   lexer.NewLexer(),
		grammar: grammar.NewGrammar(),
		libs:    libs,
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
	instance.context = semantic.NewContext(libs, func(path string) (tree.Page, error) {
		return instance.GetPage(path)
	})
	instance.semantic = semantic.NewSemantic(instance.context)
	for _, rule := range rule.SemanticRules {
		instance.semantic.AddRule(rule)
	}
	return instance
}

func (c *Complier) GetPage(path string) (tree.Page, error) {
	//TODO
	return nil, nil
}

func (c *Complier) PathDecompose(path string) (string, string, error) {
	if len(path) < 3 {
		return "", "", errors.New(fmt.Sprintf("invalid path : %v", path))
	}
	path = path[1 : len(path)-1]
	index := strings.LastIndex(path, "/")
	if index <= 0 {
		return "/", path[index+1:], nil
	}
	return path[0:index], path[index+1:], nil
}

const (
	fileLibrary = "file"
)

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
	lib := c.libs.GetLibrary(fileLibrary)
	if lib == nil {
		lib = tree.NewLibraryAdaptor()
		c.libs.AddLibrary(fileLibrary, lib)
	}
	lib.SetPage(page.GetName(), page)
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

func (c *Complier) AddRoots(roots ...string) {
	c.roots = append(c.roots, roots...)
}
