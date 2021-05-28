package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type NewContinueSeed interface {
	ToLanguage(string, concept.Closure, *NewContinue) string
	NewNull() concept.Null
	NewContinue(concept.String) *interrupt.Continue
}

type NewContinue struct {
	*adaptor.ExpressionIndex
	tag  concept.String
	seed NewContinueSeed
}

func (f *NewContinue) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewContinue) ToString(prefix string) string {
	return fmt.Sprintf("continue %v", a.tag.Value())
}

func (a *NewContinue) Anticipate(space concept.Closure) concept.Variable {
	return a.seed.NewNull()
}

func (a *NewContinue) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, a.seed.NewContinue(a.tag)
}

type NewContinueCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
	ContinueCreator        func(concept.String) *interrupt.Continue
}

type NewContinueCreator struct {
	Seeds map[string]func(string, concept.Closure, *NewContinue) string
	param *NewContinueCreatorParam
}

func (s *NewContinueCreator) New(tag concept.String) *NewContinue {
	back := &NewContinue{
		seed: s,
		tag:  tag,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewContinueCreator) NewContinue(tag concept.String) *interrupt.Continue {
	return s.param.ContinueCreator(tag)
}

func (s *NewContinueCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *NewContinueCreator) ToLanguage(language string, space concept.Closure, instance *NewContinue) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewNewContinueCreator(param *NewContinueCreatorParam) *NewContinueCreator {
	return &NewContinueCreator{
		Seeds: map[string]func(string, concept.Closure, *NewContinue) string{},
		param: param,
	}
}
