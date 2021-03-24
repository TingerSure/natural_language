package semantic

import (
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/tree"
)

type Semantic struct {
	libs    *tree.LibraryManager
	context *Context
}

func NewSemantic(libs *tree.LibraryManager) *Semantic {
	return &Semantic{
		libs:    libs,
		context: NewContext(libs),
	}
}

func (s *Semantic) Read(phrase grammar.Phrase) (*FilePage, error) {
	rule, err := s.context.GetRule(phrase)
	if err != nil {
		return nil, err
	}
	page := NewFilePage(s.libs)
	err = rule.Deal(phrase, s.context, page)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (s *Semantic) AddRule(rule *Rule) {
	s.context.AddRule(rule)
}
