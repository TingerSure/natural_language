package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
	"strings"
)

const (
	VariableArrayType = "array"
)

type ArraySeed interface {
	ToLanguage(string, concept.Pool, *Array) (string, concept.Exception)
	Type() string
	NewException(string, string) concept.Exception
}

type Array struct {
	*adaptor.AdaptorVariable
	values []concept.Variable
	seed   ArraySeed
}

func (o *Array) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Array) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *Array) ToString(prefix string) string {
	if len(a.values) == 0 {
		return "[]"
	}
	itemPrefix := fmt.Sprintf("%v\t", prefix)
	valuesToStrings := make([]string, 0, len(a.values))
	for _, value := range a.values {
		valuesToStrings = append(valuesToStrings, value.ToString(itemPrefix))
	}
	return fmt.Sprintf("[%v]", strings.Join(valuesToStrings, ", "))
}

func (a *Array) Type() string {
	return a.seed.Type()
}

func (a *Array) Set(index int, value concept.Variable) concept.Exception {
	if index < 0 || index >= a.Length() {
		return a.seed.NewException("runtime error", fmt.Sprintf("array index out of bounds error -> index/length : %v/%v", index, a.Length()))
	}
	a.values[index] = value
	return nil
}

func (a *Array) Append(value concept.Variable) {
	a.values = append(a.values, value)
}

func (a *Array) Remove(index int) concept.Exception {
	if index < 0 || index >= a.Length() {
		return a.seed.NewException("runtime error", fmt.Sprintf("array index out of bounds error -> index/length : %v/%v", index, a.Length()))
	}
	a.values = append(a.values[:index], a.values[index+1:]...)
	return nil
}

func (a *Array) Get(index int) (concept.Variable, concept.Exception) {
	if index < 0 || index >= a.Length() {
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("array index out of bounds error -> index/length : %v/%v", index, a.Length()))
	}
	return a.values[index], nil
}

func (a *Array) Length() int {
	return len(a.values)
}

type ArrayCreatorParam struct {
	NullCreator           func() concept.Null
	ExceptionCreator      func(string, string) concept.Exception
	ParamCreator          func() concept.Param
	StringCreator         func(string) concept.String
	DelayStringCreator    func(string) concept.String
	NumberCreator         func(float64) concept.Number
	DelayFunctionCreator  func(func() concept.Function) concept.Function
	SystemFunctionCreator func(
		funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
		anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
		paramNames []concept.String,
		returnNames []concept.String,
	) concept.Function
}

type ArrayCreator struct {
	Seeds map[string]func(concept.Pool, *Array) (string, concept.Exception)
	Inits []func(*Array)
	param *ArrayCreatorParam
}

func (s *ArrayCreator) New() *Array {
	array := &Array{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		values: make([]concept.Variable, 0),
		seed:   s,
	}
	for _, init := range s.Inits {
		init(array)
	}
	return array
}

func (s *ArrayCreator) ToLanguage(language string, space concept.Pool, instance *Array) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *ArrayCreator) Type() string {
	return VariableArrayType
}

func (s *ArrayCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewArrayCreator(param *ArrayCreatorParam) *ArrayCreator {
	return &ArrayCreator{
		Seeds: map[string]func(concept.Pool, *Array) (string, concept.Exception){},
		param: param,
	}
}
