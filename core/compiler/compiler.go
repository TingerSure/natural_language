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
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
	"path/filepath"
	"strings"
)

type Compiler struct {
	lexer     *lexer.Lexer
	grammar   *grammar.Grammar
	semantic  *semantic.Semantic
	libs      *tree.LibraryManager
	reading   map[string]bool
	roots     []string
	extension string
}

func NewCompiler(libs *tree.LibraryManager) *Compiler {
	instance := &Compiler{
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
	err := instance.grammar.Build()
	if err != nil {
		panic(err.Error())
	}
	instance.semantic = semantic.NewSemantic(libs, func(path string) (concept.Index, error) {
		return instance.GetPage(path)
	})
	for _, rule := range rule.SemanticRules {
		err = instance.semantic.AddRule(rule)
		if err != nil {
			panic(err.Error())
		}
	}
	return instance
}

func (c *Compiler) GetPage(path string) (concept.Index, error) {
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

func (c *Compiler) open(path string) (*os.File, error) {
	for _, root := range c.roots {
		fullPath := filepath.Join(root, path+c.extension)
		_, err := os.Stat(fullPath)
		if os.IsNotExist(err) {
			continue
		}
		return os.Open(fullPath)
	}
	return nil, errors.New(fmt.Sprintf("Path \"%v\" not found in all roots:\n%v", path, strings.Join(c.roots, "\n")))
}

func (c *Compiler) ReadPage(path string) (concept.Index, error) {
	source, err := c.open(path)
	if err != nil {
		return nil, err
	}
	tokens, err := c.lexer.Read(source, path)
	if err != nil {
		return nil, err
	}
	phrase, err := c.grammar.Read(tokens)
	if err != nil {
		return nil, err
	}
	page, err := c.semantic.Read(phrase)
	if err != nil {
		return nil, err
	}
	err = c.initPage(page, path)
	if err != nil {
		return nil, err
	}

	fmt.Println(page.ToString(""))
	return page, nil
}

func (c *Compiler) initPage(pageIndex concept.Index, path string) error {
	initKey := c.libs.Sandbox.Variable.String.New("init")
	page, exception := pageIndex.Get(nil)
	if !nl_interface.IsNil(exception) {
		return errors.New(fmt.Sprintf("Page index error: \"%v\"(\"%v\") is not an index without closure, cannot be used as a page index.", path, pageIndex.Type()))
	}
	init, exception := page.GetField(initKey)
	if !nl_interface.IsNil(exception) {
		return exception.(concept.Exception)
	}

	_, yes := variable.VariableFamilyInstance.IsNull(init)
	if yes {
		return nil
	}
	_, yes = variable.VariableFamilyInstance.IsFunctionHome(init)
	if !yes {
		return errors.New(fmt.Sprintf("\"%v\".init exist but not a function.", path))
	}
	_, exception = page.Call(initKey, c.libs.Sandbox.Variable.Param.New())
	if !nl_interface.IsNil(exception) {
		return exception.(concept.Exception)
	}
	return nil
}

func (c *Compiler) Read(path string) error {
	_, err := c.GetPage(path)
	return err
}

func (c *Compiler) GetLexer() *lexer.Lexer {
	return c.lexer
}

func (c *Compiler) GetGrammar() *grammar.Grammar {
	return c.grammar
}

func (c *Compiler) GetSemantic() *semantic.Semantic {
	return c.semantic
}

func (c *Compiler) SetExtension(extension string) {
	c.extension = extension
}

func (c *Compiler) AddRoots(roots ...string) {
	c.roots = append(c.roots, roots...)
}
