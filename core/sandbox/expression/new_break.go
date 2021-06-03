package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type NewBreakSeed interface {
	ToLanguage(string, concept.Pool, *NewBreak) (string, concept.Exception)
	NewNull() concept.Null
	NewBreak(concept.String) *interrupt.Break
}

type NewBreak struct {
	*adaptor.ExpressionIndex
	tag  concept.String
	seed NewBreakSeed
}

func (f *NewBreak) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewBreak) ToString(prefix string) string {
	return fmt.Sprintf("break %v", a.tag.Value())
}

func (a *NewBreak) Anticipate(space concept.Pool) concept.Variable {
	return a.seed.NewNull()
}

func (a *NewBreak) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return nil, a.seed.NewBreak(a.tag)
}

type NewBreakCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
	BreakCreator           func(concept.String) *interrupt.Break
}

type NewBreakCreator struct {
	Seeds map[string]func(concept.Pool, *NewBreak) (string, concept.Exception)
	param *NewBreakCreatorParam
}

func (s *NewBreakCreator) New(tag concept.String) *NewBreak {
	back := &NewBreak{
		seed: s,
		tag:  tag,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewBreakCreator) NewBreak(tag concept.String) *interrupt.Break {
	return s.param.BreakCreator(tag)
}

func (s *NewBreakCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *NewBreakCreator) ToLanguage(language string, space concept.Pool, instance *NewBreak) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewBreakCreator(param *NewBreakCreatorParam) *NewBreakCreator {
	return &NewBreakCreator{
		Seeds: map[string]func(concept.Pool, *NewBreak) (string, concept.Exception){},
		param: param,
	}
}
