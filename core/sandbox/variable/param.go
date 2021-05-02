package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
	"strings"
)

const (
	VariableParamType = "param"
	ParamDefaultKey   = "default"
)

type ParamSeed interface {
	ToLanguage(string, *Param) string
	Type() string
	NewNull() concept.Null
	New() *Param
}

type Param struct {
	*adaptor.AdaptorVariable
	values *concept.Mapping
	seed   ParamSeed
}

func (o *Param) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Param) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Param) ToString(prefix string) string {
	if 0 == a.values.Size() {
		return "{}"
	}
	subPrefix := fmt.Sprintf("%v\t", prefix)
	paramsToString := make([]string, 0, a.values.Size())

	a.values.Iterate(func(key concept.String, value interface{}) bool {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.(concept.ToString).ToString(subPrefix)))
		return false
	})

	return fmt.Sprintf("{\n%v\n%v}", strings.Join(paramsToString, ",\n"), prefix)
}

func (o *Param) Type() string {
	return o.seed.Type()
}

func (o *Param) Set(key concept.String, value concept.Variable) {
	o.values.Set(key, value)
}

func (o *Param) Get(key concept.String) concept.Variable {
	return o.values.Get(key).(concept.Variable)
}

func (o *Param) Iterate(on func(concept.String, concept.Variable) bool) bool {
	return o.values.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Variable))
	})
}

func (o *Param) Copy() concept.Param {
	param := o.seed.New()
	o.values.Iterate(func(key concept.String, value interface{}) bool {
		param.Set(key, value.(concept.Variable))
		return false
	})
	return param
}

type ParamCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type ParamCreator struct {
	Seeds map[string]func(string, *Param) string
	param *ParamCreatorParam
}

func (s *ParamCreator) New() *Param {
	return &Param{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		values: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewNull(),
		}),
		seed: s,
	}
}

func (s *ParamCreator) ToLanguage(language string, instance *Param) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ParamCreator) Type() string {
	return VariableParamType
}

func (s *ParamCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewParamCreator(param *ParamCreatorParam) *ParamCreator {
	return &ParamCreator{
		Seeds: map[string]func(string, *Param) string{},
		param: param,
	}
}
