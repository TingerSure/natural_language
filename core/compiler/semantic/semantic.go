package semantic

import (
	"errors"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Semantic struct {
	context *Context
}

func NewSemantic(context *Context) *Semantic {
	return &Semantic{
		context: context,
	}
}

func (s *Semantic) Read(phrase grammar.Phrase) (concept.Index, error) {
	pageIndex, err := s.context.Deal(phrase)
	if err != nil {
		return nil, err
	}
	if len(pageIndex) != 1 {
		return nil, errors.New("Illegal global semantic rule whose result is not unique.")
	}
	return pageIndex[0], nil
}

func (s *Semantic) AddRule(rule *Rule) {
	s.context.AddRule(rule)
}
