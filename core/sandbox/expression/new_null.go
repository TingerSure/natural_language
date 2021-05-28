package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type NewNullSeed interface {
	ToLanguage(string, concept.Closure, *NewNull) string
	NewNull() concept.Null
}

type NewNull struct {
	*adaptor.ExpressionIndex
	seed NewNullSeed
}

func (f *NewNull) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewNull) ToString(prefix string) string {
	return "null"
}

func (a *NewNull) Anticipate(space concept.Closure) concept.Variable {
	return a.seed.NewNull()
}

func (a *NewNull) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return a.seed.NewNull(), nil
}

type NewNullCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
}

type NewNullCreator struct {
	Seeds map[string]func(string, concept.Closure, *NewNull) string
	param *NewNullCreatorParam
}

func (s *NewNullCreator) New() *NewNull {
	back := &NewNull{
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewNullCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *NewNullCreator) ToLanguage(language string, space concept.Closure, instance *NewNull) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewNewNullCreator(param *NewNullCreatorParam) *NewNullCreator {
	return &NewNullCreator{
		Seeds: map[string]func(string, concept.Closure, *NewNull) string{},
		param: param,
	}
}
