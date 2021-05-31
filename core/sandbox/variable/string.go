package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
	"strings"
)

const (
	VariableStringType = "string"
)

type StringSeed interface {
	ToLanguage(string, concept.Closure, *String) string
	Type() string
	New(string) *String
}

type String struct {
	*adaptor.AdaptorVariable
	value   string
	mapping map[string]string
	seed    StringSeed
}

func (o *String) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *String) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (n *String) GetLanguage(language string) string {
	value := n.mapping[language]
	if value == "" {
		return n.value
	}
	return value
}

func (n *String) SetLanguage(language string, value string) {
	n.mapping[language] = value
}

func (n *String) HasLanguage(language string) bool {
	return n.mapping[language] != ""
}

func (n *String) IsLanguage(language string, value string) bool {
	return n.mapping[language] == value
}

func (n *String) Equal(other concept.String) bool {
	return n.value == other.Value()
}

func (n *String) IterateLanguages(on func(string, string) bool) bool {
	for language, value := range n.mapping {
		if on(language, value) {
			return true
		}
	}
	return false
}

func (n *String) ToString(prefix string) string {
	return fmt.Sprintf("\"%v\"", n.value)
}

func (n *String) Value() string {
	return n.value
}

func (s *String) Type() string {
	return s.seed.Type()
}

func (s *String) Clone() concept.String {
	instance := s.seed.New(s.value)
	for language, value := range s.mapping {
		instance.mapping[language] = value
	}
	return instance
}

func (s *String) CloneTo(instance concept.String) {
	for language, value := range s.mapping {
		instance.SetLanguage(language, value)
	}
}

type StringCreatorParam struct {
	NullCreator           func() concept.Null
	ExceptionCreator      func(string, string) concept.Exception
	ParamCreator          func() concept.Param
	StringCreator         func(string) concept.String
	DelayStringCreator    func(string) concept.String
	DelayFunctionCreator  func(func() concept.Function) concept.Function
	SystemFunctionCreator func(
		funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
		anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
		paramNames []concept.String,
		returnNames []concept.String,
	) concept.Function
}

type StringCreator struct {
	Seeds map[string]func(string, concept.Closure, *String) string
	param *StringCreatorParam
}

func (s *StringCreator) New(value string) *String {
	back := &String{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		value:   value,
		mapping: make(map[string]string),
		seed:    s,
	}
	back.SetField(s.param.DelayStringCreator("SetLanguage"), s.param.DelayFunctionCreator(s.fieldSetLanguage(back)))
	back.SetField(s.param.DelayStringCreator("GetLanguage"), s.param.DelayFunctionCreator(s.fieldGetLanguage(back)))
	return back
}

func (s *StringCreator) fieldSetLanguage(target *String) func() concept.Function {
	return func() concept.Function {
		languageParam := s.param.StringCreator("language")
		valueParam := s.param.StringCreator("value")
		return s.param.SystemFunctionCreator(
			func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				languagePre := param.Get(languageParam)
				language, yes := VariableFamilyInstance.IsStringHome(languagePre)
				if !yes {
					return nil, s.param.ExceptionCreator("type error", fmt.Sprintf("Param language is not a string: %v", languagePre))
				}
				valuePre := param.Get(valueParam)
				value, yes := VariableFamilyInstance.IsStringHome(valuePre)
				if !yes {
					return nil, s.param.ExceptionCreator("type error", fmt.Sprintf("Param value is not a string: %v", languagePre))
				}
				target.SetLanguage(language.Value(), value.Value())
				return s.param.ParamCreator(), nil
			},
			func(param concept.Param, _ concept.Variable) concept.Param {
				return s.param.ParamCreator()
			},
			[]concept.String{
				languageParam,
				valueParam,
			},
			[]concept.String{},
		)
	}
}

func (s *StringCreator) fieldGetLanguage(target *String) func() concept.Function {
	return func() concept.Function {
		languageParam := s.param.StringCreator("language")
		valueBack := s.param.StringCreator("value")
		return s.param.SystemFunctionCreator(
			func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				languagePre := input.Get(languageParam)
				language, yes := VariableFamilyInstance.IsStringHome(languagePre)
				if !yes {
					return nil, s.param.ExceptionCreator("type error", fmt.Sprintf("Param language is not a string: %v", languagePre))
				}
				value := target.GetLanguage(language.Value())
				output := s.param.ParamCreator()
				output.Set(valueBack, s.param.StringCreator(value))
				return output, nil
			},
			nil,
			[]concept.String{
				languageParam,
			},
			[]concept.String{
				valueBack,
			},
		)
	}
}

func (s *StringCreator) ToLanguage(language string, space concept.Closure, instance *String) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *StringCreator) Type() string {
	return VariableStringType
}

func NewStringCreator(param *StringCreatorParam) *StringCreator {
	return &StringCreator{
		Seeds: map[string]func(string, concept.Closure, *String) string{},
		param: param,
	}
}

func StringJoin(values []concept.String, separator string) string {
	paramsToString := make([]string, 0, len(values))
	for _, value := range values {
		paramsToString = append(paramsToString, value.Value())
	}
	return strings.Join(paramsToString, separator)
}
