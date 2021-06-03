package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type NewNullSeed interface {
	ToLanguage(string, concept.Pool, *NewNull) (string, concept.Exception)
	NewNull() concept.Null
}

type NewNull struct {
	*adaptor.ExpressionIndex
	seed NewNullSeed
}

func (f *NewNull) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewNull) ToString(prefix string) string {
	return "null"
}

func (a *NewNull) Anticipate(space concept.Pool) concept.Variable {
	return a.seed.NewNull()
}

func (a *NewNull) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return a.seed.NewNull(), nil
}

type NewNullCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
}

type NewNullCreator struct {
	Seeds map[string]func(concept.Pool, *NewNull) (string, concept.Exception)
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

func (s *NewNullCreator) ToLanguage(language string, space concept.Pool, instance *NewNull) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewNullCreator(param *NewNullCreatorParam) *NewNullCreator {
	return &NewNullCreator{
		Seeds: map[string]func(concept.Pool, *NewNull) (string, concept.Exception){},
		param: param,
	}
}
