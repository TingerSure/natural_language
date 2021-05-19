package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type NewBreakSeed interface {
	ToLanguage(string, *NewBreak) string
	NewNull() concept.Null
	NewBreak(concept.String) *interrupt.Break
}

type NewBreak struct {
	*adaptor.ExpressionIndex
	tag  concept.String
	seed NewBreakSeed
}

func (f *NewBreak) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewBreak) ToString(prefix string) string {
	return fmt.Sprintf("break %v", a.tag.ToString(prefix))
}

func (a *NewBreak) Anticipate(space concept.Closure) concept.Variable {
	return a.seed.NewNull()
}

func (a *NewBreak) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, a.seed.NewBreak(a.tag)
}

type NewBreakCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
	BreakCreator           func(concept.String) *interrupt.Break
}

type NewBreakCreator struct {
	Seeds map[string]func(string, *NewBreak) string
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

func (s *NewBreakCreator) ToLanguage(language string, instance *NewBreak) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewBreakCreator(param *NewBreakCreatorParam) *NewBreakCreator {
	return &NewBreakCreator{
		Seeds: map[string]func(string, *NewBreak) string{},
		param: param,
	}
}
