package semantic

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/tree"
)

type Semantic struct {
	rules map[*grammar.Rule]*Rule
	libs  *tree.LibraryManager
}

func NewSemantic(libs *tree.LibraryManager) *Semantic {
	return &Semantic{
		rules: map[*grammar.Rule]*Rule{},
		libs:  libs,
	}
}

func (s *Semantic) Read(phrase grammar.Phrase) (*FilePage, error) {
	if phrase.PhraseType() == grammar.PhraseTypeToken {
		return nil, errors.New("Grammar rules integrity error.")
	}
	rule := s.rules[phrase.GetRule()]
	if rule == nil {
		return nil, errors.New(fmt.Sprintf("No semantic rule match grammar rule : %v", phrase.GetRule().ToString()))
	}
	page := NewFilePage(s.libs)
	err := rule.Deal(phrase, page)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (s *Semantic) AddRule(rule *Rule) {
	s.rules[rule.GetSource()] = rule
}
