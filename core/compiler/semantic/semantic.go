package semantic

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type Semantic struct {
	rules   map[*grammar.Rule]*Rule
	getPage func(path string) (concept.Pipe, error)
	libs    *tree.LibraryManager
}

func NewSemantic(libs *tree.LibraryManager, getPage func(path string) (concept.Pipe, error)) *Semantic {
	return &Semantic{
		rules:   map[*grammar.Rule]*Rule{},
		getPage: getPage,
		libs:    libs,
	}
}

func (s *Semantic) Read(phrase grammar.Phrase, path string, content []byte) (concept.Pipe, error) {
	context := NewContext(s.libs, s.getPage, s.rules, path, content)
	pageIndex, _, err := context.Deal(phrase)
	if err != nil {
		return nil, err
	}
	if len(pageIndex) != 1 {
		return nil, errors.New("Illegal global semantic rule whose result is not unique.")
	}
	return pageIndex[0], nil
}

func (s *Semantic) AddRule(rule *Rule) error {
	if s.rules[rule.GetSource()] != nil {
		return fmt.Errorf("Semantic rule (%v) repeated.", rule.GetSource())
	}
	s.rules[rule.GetSource()] = rule
	return nil
}
