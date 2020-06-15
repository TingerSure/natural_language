package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"strings"
)

type NewParamSeed interface {
	ToLanguage(string, *NewParam) string
	NewParam() concept.Param
}

type NewParam struct {
	*adaptor.ExpressionIndex
	values map[concept.String]concept.Index
	seed   NewParamSeed
}

func (f *NewParam) Iterate(on func(concept.String, concept.Index) bool) bool {
	for key, value := range f.values {
		if on(key, value) {
			return true
		}
	}
	return false
}

func (f *NewParam) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewParam) ToString(prefix string) string {
	if 0 == len(a.values) {
		return "{}"
	}
	subPrefix := fmt.Sprintf("%v\t", prefix)
	paramsToString := make([]string, 0, len(a.values))

	for key, value := range a.values {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.ToString(subPrefix)))
	}

	return fmt.Sprintf("{\n%v\n%v}", strings.Join(paramsToString, ",\n"), prefix)
}

func (a *NewParam) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	if len(a.values) == 0 {
		return a.seed.NewParam(), nil
	}
	var suspend concept.Interrupt = nil

	param := a.seed.NewParam().Init(func(on func(concept.String, concept.Variable) bool) bool {
		for key, index := range a.values {
			value, subSuspend := index.Get(space)
			if !nl_interface.IsNil(subSuspend) {
				suspend = subSuspend
				return true
			}
			if on(key, value) {
				return true
			}
		}
		return false

	})

	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return param, nil
}

func (a *NewParam) Init(values map[concept.String]concept.Index) *NewParam {
	a.values = values
	return a
}

type NewParamCreatorParam struct {
	ExpressionIndexCreator func(func(concept.Closure) (concept.Variable, concept.Interrupt)) *adaptor.ExpressionIndex
	ParamCreator           func() concept.Param
}

type NewParamCreator struct {
	Seeds map[string]func(string, *NewParam) string
	param *NewParamCreatorParam
}

func (s *NewParamCreator) New() *NewParam {
	back := &NewParam{
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back.Exec)
	return back
}

func (s *NewParamCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *NewParamCreator) ToLanguage(language string, instance *NewParam) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewParamCreator(param *NewParamCreatorParam) *NewParamCreator {
	return &NewParamCreator{
		Seeds: map[string]func(string, *NewParam) string{},
		param: param,
	}
}
