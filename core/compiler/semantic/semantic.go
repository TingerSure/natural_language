package semantic

import (
	"errors"
	"github.com/TingerSure/natural_language/core/compiler/grammar"
)

type Semantic struct {
	context *Context
}

func NewSemantic(context *Context) *Semantic {
	return &Semantic{
		context: context,
	}
}

func (s *Semantic) Read(phrase grammar.Phrase) (*FilePage, error) {
	pageIndex, err := s.context.Deal(phrase)
	if err != nil {
		return nil, err
	}
	if len(pageIndex) != 1 {
		return nil, errors.New("Illegal global semantic rule whose result is not unique.")
	}
	page, ok := pageIndex[0].(*FilePage)
	if !ok {
		return nil, errors.New("Illegal global semantic rule whose result is not allowed.")
	}
	return page, nil
}

func (s *Semantic) AddRule(rule *Rule) {
	s.context.AddRule(rule)
}
