package semantic

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type Context struct {
	rules   map[*grammar.Rule]*Rule
	getPage func(path string) (concept.Pipe, error)
	libs    *tree.LibraryManager
}

func NewContext(libs *tree.LibraryManager, getPage func(path string) (concept.Pipe, error), rules map[*grammar.Rule]*Rule) *Context {
	return &Context{
		rules:   rules,
		getPage: getPage,
		libs:    libs,
	}
}

func (c *Context) FormatSymbolString(value string) string {
	return value[1 : len(value)-1]
}

func (c *Context) GetLibraryManager() *tree.LibraryManager {
	return c.libs
}

func (c *Context) GetPage(path string) (concept.Pipe, error) {
	return c.getPage(path)
}

func (c *Context) Deal(phrase grammar.Phrase) ([]concept.Pipe, error) {
	rule, err := c.GetRule(phrase)
	if err != nil {
		return nil, err
	}
	return rule.Deal(phrase, c)
}

func (c *Context) GetRule(phrase grammar.Phrase) (*Rule, error) {
	if phrase.PhraseType() == grammar.PhraseTypeToken {
		return nil, errors.New("Token phrase has no semantic rule.")
	}
	rule := c.rules[phrase.GetRule()]
	if rule == nil {
		return nil, fmt.Errorf("No semantic rule match grammar rule : %v", phrase.GetRule().ToString())
	}
	return rule, nil
}
