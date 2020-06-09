package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ParamSetSeed interface {
	ToLanguage(string, *ParamSet) string
	NewException(string, string) concept.Exception
}

type ParamSet struct {
	*adaptor.ExpressionIndex
	key   concept.String
	value concept.Index
	param concept.Index
	seed  ParamSetSeed
}

func (f *ParamSet) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *ParamSet) ToString(prefix string) string {
	return fmt.Sprintf("%v[%v] = %v", a.param.ToString(prefix), a.key.ToString(prefix), a.value.ToString(prefix))
}

func (a *ParamSet) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return nil, a.seed.NewException("type error", "Only Param can be set in ParamSet")
	}

	preValue, suspend := a.value.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	param.Set(a.key, preValue)
	return preValue, nil
}

type ParamSetCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	ExpressionIndexCreator func(func(concept.Closure) (concept.Variable, concept.Interrupt)) *adaptor.ExpressionIndex
}

type ParamSetCreator struct {
	Seeds map[string]func(string, *ParamSet) string
	param *ParamSetCreatorParam
}

func (s *ParamSetCreator) New(param concept.Index, key concept.String, value concept.Index) *ParamSet {
	back := &ParamSet{
		key:   key,
		value: value,
		param: param,
		seed:  s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back.Exec)
	return back
}

func (s *ParamSetCreator) ToLanguage(language string, instance *ParamSet) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}
func (s *ParamSetCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}
func NewParamSetCreator(param *ParamSetCreatorParam) *ParamSetCreator {
	return &ParamSetCreator{
		Seeds: map[string]func(string, *ParamSet) string{},
		param: param,
	}
}
