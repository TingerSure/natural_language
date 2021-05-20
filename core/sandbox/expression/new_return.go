package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type NewReturnSeed interface {
	ToLanguage(string, *NewReturn) string
	NewNull() concept.Null
	NewReturn() *interrupt.Return
}

type NewReturn struct {
	*adaptor.ExpressionIndex
	seed NewReturnSeed
}

func (f *NewReturn) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewReturn) ToString(prefix string) string {
	return "return"
}

func (a *NewReturn) Anticipate(space concept.Closure) concept.Variable {
	return a.seed.NewNull()
}

func (a *NewReturn) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, a.seed.NewReturn()
}

type NewReturnCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
	ReturnCreator          func() *interrupt.Return
}

type NewReturnCreator struct {
	Seeds map[string]func(string, *NewReturn) string
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

func (s *NewReturnCreator) ToLanguage(language string, instance *NewReturn) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewReturnCreator(param *NewReturnCreatorParam) *NewReturnCreator {
	return &NewReturnCreator{
		Seeds: map[string]func(string, *NewReturn) string{},
		param: param,
	}
}
