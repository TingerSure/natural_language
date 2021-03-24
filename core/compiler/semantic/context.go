package semantic

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/tree"
)

type Context struct {
	rules map[*grammar.Rule]*Rule
	libs  *tree.LibraryManager
}

func NewContext(libs *tree.LibraryManager) *Context {
	return &Context{
		libs:  libs,
		rules: map[*grammar.Rule]*Rule{},
	}
}

func (c *Context) GetImport(path string) (tree.Page, error) {
	// TODO
	return nil, nil
}

func (c *Context) Deal(phrase grammar.Phrase, context *Context, page *FilePage) error {
	rule, err := c.GetRule(phrase)
	if err != nil {
		return err
	}
	return rule.Deal(phrase, context, page)
}

func (c *Context) GetRule(phrase grammar.Phrase) (*Rule, error) {
	if phrase.PhraseType() == grammar.PhraseTypeToken {
		return nil, errors.New("Token phrase has no semantic rule.")
	}
	rule := c.rules[phrase.GetRule()]
	if rule == nil {
		return nil, errors.New(fmt.Sprintf("No semantic rule match grammar rule : %v", phrase.GetRule().ToString()))
	}
	return rule, nil
}

func (c *Context) AddRule(rule *Rule) {
	c.rules[rule.GetSource()] = rule
}
