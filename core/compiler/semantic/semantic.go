package semantic

import (
	"errors"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type Semantic struct {
	rules   map[*grammar.Rule]*Rule
	getPage func(path string) (concept.Index, error)
	libs    *tree.LibraryManager
}

func NewSemantic(libs *tree.LibraryManager, getPage func(path string) (concept.Index, error)) *Semantic {
	return &Semantic{
		rules:   map[*grammar.Rule]*Rule{},
		getPage: getPage,
		libs:    libs,
	}
}

func (s *Semantic) Read(phrase grammar.Phrase) (concept.Index, error) {
	context := NewContext(s.libs, s.getPage, s.rules)
	pageIndex, err := context.Deal(phrase)
	if err != nil {
		return nil, err
	}
	if len(pageIndex) != 1 {
		return nil, errors.New("Illegal global semantic rule whose result is not unique.")
	}
	return pageIndex[0], nil
}

func (s *Semantic) AddRule(rule *Rule) {
	s.rules[rule.GetSource()] = rule
}
