package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type NewReturnSeed interface {
	ToLanguage(string, concept.Pool, *NewReturn) (string, concept.Exception)
	NewNull() concept.Null
	NewReturn() *interrupt.Return
}

type NewReturn struct {
	*adaptor.ExpressionIndex
	seed NewReturnSeed
}

func (f *NewReturn) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewReturn) ToString(prefix string) string {
	return "return"
}

func (a *NewReturn) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return nil, a.seed.NewReturn()
}

type NewReturnCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
	ReturnCreator          func() *interrupt.Return
}

type NewReturnCreator struct {
	Seeds map[string]func(concept.Pool, *NewReturn) (string, concept.Exception)
	param *NewReturnCreatorParam
}

func (s *NewReturnCreator) New() *NewReturn {
	back := &NewReturn{
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewReturnCreator) NewReturn() *interrupt.Return {
	return s.param.ReturnCreator()
}

func (s *NewReturnCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *NewReturnCreator) ToLanguage(language string, space concept.Pool, instance *NewReturn) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewReturnCreator(param *NewReturnCreatorParam) *NewReturnCreator {
	return &NewReturnCreator{
		Seeds: map[string]func(concept.Pool, *NewReturn) (string, concept.Exception){},
		param: param,
	}
}
