package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableValueLanguageFunctionType = "value_language_function"
	FunctionValueLanguageFunctionType = "value_language"
)

type ValueLanguageFunctionSeed interface {
	ToLanguage(string, concept.Pool, *ValueLanguageFunction) (string, concept.Exception)
	Type() string
}

type ValueLanguageFunction struct {
	*adaptor.AdaptorFunction
	languageOnCallDefaultSeed func(string, concept.Function, concept.Pool, string, concept.Param) (string, concept.Exception)
	seed                      ValueLanguageFunctionSeed
}

func (o *ValueLanguageFunction) SetLanguageOnCallDefaultSeed(defaultSeed func(string, concept.Function, concept.Pool, string, concept.Param) (string, concept.Exception)) {
	o.languageOnCallDefaultSeed = defaultSeed
}

func (o *ValueLanguageFunction) GetLanguageOnCallDefaultSeed() func(string, concept.Function, concept.Pool, string, concept.Param) (string, concept.Exception) {
	return o.languageOnCallDefaultSeed
}

func (o *ValueLanguageFunction) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *ValueLanguageFunction) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (f *ValueLanguageFunction) ToCallLanguage(language string, space concept.Pool, self string, param concept.Param) (string, concept.Exception) {
	seed := f.GetLanguageOnCallSeed(language)
	if seed != nil {
		return seed(f, space, self, param)
	}
	if f.languageOnCallDefaultSeed != nil {
		return f.languageOnCallDefaultSeed(language, f, space, self, param)
	}
	return f.ToLanguage(language, space)
}

func (f *ValueLanguageFunction) ToString(prefix string) string {
	return fmt.Sprintf("value_language_function (%v) %v {}", StringJoin(f.ParamNames(), ", "), StringJoin(f.ReturnNames(), ", "))
}

func (f *ValueLanguageFunction) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	return params
}

func (f *ValueLanguageFunction) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	return params, nil
}

func (s *ValueLanguageFunction) Type() string {
	return s.seed.Type()
}

func (s *ValueLanguageFunction) FunctionType() string {
	return FunctionValueLanguageFunctionType
}

type ValueLanguageFunctionCreatorParam struct {
	NullCreator      func() concept.Null
	ParamCreator     func() concept.Param
	ExceptionCreator func(string, string) concept.Exception
}

type ValueLanguageFunctionCreator struct {
	Seeds map[string]func(concept.Pool, *ValueLanguageFunction) (string, concept.Exception)
	Inits []func(*ValueLanguageFunction)
	param *ValueLanguageFunctionCreatorParam
}

func (s *ValueLanguageFunctionCreator) New(paramNames []concept.String, returnNames []concept.String) *ValueLanguageFunction {
	valueLanguage := &ValueLanguageFunction{
		AdaptorFunction: adaptor.NewAdaptorFunction(&adaptor.AdaptorFunctionParam{
			NullCreator:      s.param.NullCreator,
			ParamCreator:     s.param.ParamCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		seed: s,
	}

	valueLanguage.AddParamName(paramNames...)
	valueLanguage.AddReturnName(returnNames...)
	for _, init := range s.Inits {
		init(valueLanguage)
	}
	return valueLanguage
}

func (s *ValueLanguageFunctionCreator) ToLanguage(language string, space concept.Pool, instance *ValueLanguageFunction) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *ValueLanguageFunctionCreator) Type() string {
	return VariableValueLanguageFunctionType
}

func NewValueLanguageFunctionCreator(param *ValueLanguageFunctionCreatorParam) *ValueLanguageFunctionCreator {
	return &ValueLanguageFunctionCreator{
		Seeds: map[string]func(concept.Pool, *ValueLanguageFunction) (string, concept.Exception){},
		param: param,
	}
}
