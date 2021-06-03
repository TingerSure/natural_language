package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type NewContinueSeed interface {
	ToLanguage(string, concept.Pool, *NewContinue) (string, concept.Exception)
	NewNull() concept.Null
	NewContinue(concept.String) *interrupt.Continue
}

type NewContinue struct {
	*adaptor.ExpressionIndex
	tag  concept.String
	seed NewContinueSeed
}

func (f *NewContinue) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewContinue) ToString(prefix string) string {
	return fmt.Sprintf("continue %v", a.tag.Value())
}

func (a *NewContinue) Anticipate(space concept.Pool) concept.Variable {
	return a.seed.NewNull()
}

func (a *NewContinue) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return nil, a.seed.NewContinue(a.tag)
}

type NewContinueCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
	ContinueCreator        func(concept.String) *interrupt.Continue
}

type NewContinueCreator struct {
	Seeds map[string]func(concept.Pool, *NewContinue) (string, concept.Exception)
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

func (s *NewContinueCreator) ToLanguage(language string, space concept.Pool, instance *NewContinue) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewContinueCreator(param *NewContinueCreatorParam) *NewContinueCreator {
	return &NewContinueCreator{
		Seeds: map[string]func(concept.Pool, *NewContinue) (string, concept.Exception){},
		param: param,
	}
}
