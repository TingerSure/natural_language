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
	ToLanguage(string, concept.Pool, *Param) (string, concept.Exception)
	Type() string
	NewNull() concept.Null
	NewString(string) concept.String
	New() *Param
}

type Param struct {
	*adaptor.AdaptorVariable
	list  []concept.Variable
	types int
	seed  ParamSeed
}

func (o *Param) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Param) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Param) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	if a.types == concept.ParamTypeList {
		if a.SizeIndex() == 0 {
			return ""
		}
		paramsToString := make([]string, 0, len(a.list))
		for _, value := range a.list {
			paramsToString = append(paramsToString, value.ToString(subPrefix))
		}
		return strings.Join(paramsToString, ", ")
	}
	if a.types == concept.ParamTypeKeyValue {
		if a.SizeField() == 0 {
			return ""
		}
		paramsToString := make([]string, 0, a.SizeField())
		a.Iterate(func(key concept.String, value concept.Variable) bool {
			paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.Value(), value.ToString(subPrefix)))
			return false
		})
		return strings.Join(paramsToString, ",\n")
	}
	return ""
}

func (o *Param) Type() string {
	return o.seed.Type()
}

func (o *Param) Set(key concept.String, value concept.Variable) {
	o.types = concept.ParamTypeKeyValue
	o.SetField(key, value)
}

func (o *Param) Get(key concept.String) concept.Variable {
	value, _ := o.GetField(key)
	return value
}

func (o *Param) SetOriginal(key string, value concept.Variable) {
	o.Set(o.seed.NewString(key), value)
}

func (o *Param) GetOriginal(key string) concept.Variable {
	return o.Get(o.seed.NewString(key))
}

func (o *Param) SizeIndex() int {
	return len(o.list)
}

func (o *Param) AppendIndex(value concept.Variable) {
	o.types = concept.ParamTypeList
	o.list = append(o.list, value)
}

func (o *Param) SetIndex(index int, value concept.Variable) {
	o.types = concept.ParamTypeList
	o.list[index] = value
}

func (o *Param) GetIndex(index int) concept.Variable {
	return o.list[index]
}

func (o *Param) ParamType() int {
	return o.types
}

type ParamCreatorParam struct {
	NullCreator      func() concept.Null
	StringCreator    func(string) concept.String
	ExceptionCreator func(string, string) concept.Exception
}

type ParamCreator struct {
	Seeds map[string]func(concept.Pool, *Param) (string, concept.Exception)
	param *ParamCreatorParam
}

func (s *ParamCreator) New() *Param {
	return &Param{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		types: concept.ParamTypeKeyValue,
		seed:  s,
	}
}

func (s *ParamCreator) ToLanguage(language string, space concept.Pool, instance *Param) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *ParamCreator) Type() string {
	return VariableParamType
}

func (s *ParamCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *ParamCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func NewParamCreator(param *ParamCreatorParam) *ParamCreator {
	return &ParamCreator{
		Seeds: map[string]func(concept.Pool, *Param) (string, concept.Exception){},
		param: param,
	}
}
