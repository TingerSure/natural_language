package compiler

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"github.com/TingerSure/natural_language/core/compiler/rule"
	"github.com/TingerSure/natural_language/core/compiler/semantic"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
	"path/filepath"
	"strings"
)

type Complier struct {
	lexer    *lexer.Lexer
	grammar  *grammar.Grammar
	semantic *semantic.Semantic
	context  *semantic.Context
	libs     *tree.LibraryManager
	reading  map[string]bool
	roots    []string
}

func NewComplier(libs *tree.LibraryManager) *Complier {
	instance := &Complier{
		lexer:   lexer.NewLexer(),
		grammar: grammar.NewGrammar(),
		libs:    libs,
		reading: map[string]bool{},
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
	instance.context = semantic.NewContext(libs, func(path string) (concept.Index, error) {
		return instance.GetPage(path)
	})
	instance.semantic = semantic.NewSemantic(instance.context)
	for _, rule := range rule.SemanticRules {
		instance.semantic.AddRule(rule)
	}
	return instance
}

func (c *Complier) GetPage(path string) (concept.Index, error) {
	page := c.libs.GetPage(path)
	if !nl_interface.IsNil(page) {
		return page, nil
	}
	if c.reading[path] {
		return nil, errors.New(fmt.Sprintf("Import cycle: \"%v\".", path))
	}
	c.reading[path] = true
	page, err := c.ReadPage(path)
	if err != nil {
		return nil, err
	}
	c.libs.AddPage(path, page)
	c.reading[path] = false
	return page, nil
}

func (c *Complier) open(path string) (*os.File, error) {
	for _, root := range c.roots {
		fullPath := filepath.Join(root, path)
		_, err := os.Stat(fullPath)
		if os.IsNotExist(err) {
			continue
		}
		return os.Open(fullPath)
	}
	return nil, errors.New(fmt.Sprintf("Path not found in all roots: \"%v\".\n%v", path, strings.Join(c.roots, "\n")))
}

func (c *Complier) ReadPage(path string) (concept.Index, error) {
	source, err := c.open(path)
	if err != nil {
		return nil, err
	}
	tokens, err := c.lexer.Read(source)
	if err != nil {
		return nil, err
	}
	phrase, err := c.grammar.Read(tokens)
	if err != nil {
		return nil, err
	}
	return c.semantic.Read(phrase)
}

func (c *Complier) Read(path string) error {
	_, err := c.GetPage(path)
	// TODO run init()
	return err
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
