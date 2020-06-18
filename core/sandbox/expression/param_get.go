package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ParamGetSeed interface {
	ToLanguage(string, *ParamGet) string
	NewNull() concept.Null
	NewException(string, string) concept.Exception
}

type ParamGet struct {
	*adaptor.ExpressionIndex
	key   concept.String
	param concept.Index
	seed  ParamGetSeed
}

func (f *ParamGet) Key() concept.String {
	return f.key
}

func (f *ParamGet) Param() concept.Index {
	return f.param
}

func (f *ParamGet) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)

}

func (a *ParamGet) ToString(prefix string) string {
	return fmt.Sprintf("%v[%v]", a.param.ToString(prefix), a.key.ToString(prefix))
}

func (a *ParamGet) Anticipate(space concept.Closure) concept.Variable {
	preParam := a.param.Anticipate(space)
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return a.seed.NewNull()
	}
	return param.Get(a.key)
}

func (a *ParamGet) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return a.seed.NewNull(), suspend
	}
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return a.seed.NewNull(), a.seed.NewException("type error", "Only Param can be get in ParamGet")
	}

	return param.Get(a.key), nil
}

type ParamGetCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	NullCreator            func() concept.Null
	ExpressionIndexCreator func(func(concept.Closure) (concept.Variable, concept.Interrupt)) *adaptor.ExpressionIndex
}

type ParamGetCreator struct {
	Seeds map[string]func(string, *ParamGet) string
	param *ParamGetCreatorParam
}

func (s *ParamGetCreator) New(param concept.Index, key concept.String) *ParamGet {
	back := &ParamGet{
		key:   key,
		param: param,
		seed:  s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back.Exec)
	return back
}

func (s *ParamGetCreator) ToLanguage(language string, instance *ParamGet) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ParamGetCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *ParamGetCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}
func NewParamGetCreator(param *ParamGetCreatorParam) *ParamGetCreator {
	return &ParamGetCreator{
		Seeds: map[string]func(string, *ParamGet) string{},
		param: param,
	}
}
