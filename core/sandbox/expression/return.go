package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type ReturnSeed interface {
	NewNull() concept.Null
	ToLanguage(string, *Return) string
}

type Return struct {
	*adaptor.ExpressionIndex
	key    concept.String
	result concept.Index
	seed   ReturnSeed
}

func (f *Return) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)

}

func (a *Return) Key() concept.String {
	return a.key
}

func (a *Return) ToString(prefix string) string {
	return fmt.Sprintf("return[%v] %v", a.key.ToString(prefix), a.result.ToString(prefix))
}

func (a *Return) Anticipate(space concept.Closure) concept.Variable {
	return a.result.Anticipate(space)
}

func (a *Return) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	result, suspend := a.result.Get(space)

	if !nl_interface.IsNil(suspend) {
		return a.seed.NewNull(), suspend
	}
	space.SetReturn(a.key, result)
	return result, nil
}

type ReturnCreatorParam struct {
	NullCreator            func() concept.Null
	ExpressionIndexCreator func(func(concept.Closure) (concept.Variable, concept.Interrupt)) *adaptor.ExpressionIndex
}

type ReturnCreator struct {
	Seeds map[string]func(string, *Return) string
	param *ReturnCreatorParam
}

func (s *ReturnCreator) New(key concept.String, result concept.Index) *Return {
	back := &Return{
		key:    key,
		result: result,
		seed:   s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back.Exec)
	return back
}

func (s *ReturnCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *ReturnCreator) ToLanguage(language string, instance *Return) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewReturnCreator(param *ReturnCreatorParam) *ReturnCreator {
	return &ReturnCreator{
		Seeds: map[string]func(string, *Return) string{},
		param: param,
	}
}
