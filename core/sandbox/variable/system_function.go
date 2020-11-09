package variable

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableSystemFunctionType = "system_function"
	FunctionSystemFunctionType = "system"
)

type SystemFunctionSeed interface {
	ToLanguage(string, *SystemFunction) string
	Type() string
}

type SystemFunction struct {
	*adaptor.AdaptorFunction
	name            concept.String
	paramNames      []concept.String
	returnNames     []concept.String
	funcs           func(concept.Param, concept.Object) (concept.Param, concept.Exception)
	anticipateFuncs func(concept.Param, concept.Object) concept.Param
	seed            SystemFunctionSeed
}

func (f *SystemFunction) ParamFormat(params *concept.Mapping) *concept.Mapping {
	return f.AdaptorFunction.AdaptorParamFormat(f, params)
}

func (f *SystemFunction) ReturnFormat(back concept.String) concept.String {
	return f.AdaptorFunction.AdaptorReturnFormat(f, back)
}

func (f *SystemFunction) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *SystemFunction) Name() concept.String {
	return s.name
}

func (s *SystemFunction) ParamNames() []concept.String {
	return s.paramNames
}

func (s *SystemFunction) ReturnNames() []concept.String {
	return s.returnNames
}

func (f *SystemFunction) ToString(prefix string) string {
	return f.name.ToString(prefix)
}

func (f *SystemFunction) Anticipate(params concept.Param, object concept.Object) concept.Param {
	return f.anticipateFuncs(params, object)
}

func (f *SystemFunction) Exec(params concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	return f.funcs(params, object)
}

func (s *SystemFunction) Type() string {
	return s.seed.Type()
}

func (s *SystemFunction) FunctionType() string {
	return FunctionSystemFunctionType
}

type SystemFunctionCreatorParam struct {
}

type SystemFunctionCreator struct {
	Seeds map[string]func(string, *SystemFunction) string
	param *SystemFunctionCreatorParam
}

func (s *SystemFunctionCreator) New(
	name concept.String,
	funcs func(concept.Param, concept.Object) (concept.Param, concept.Exception),
	anticipateFuncs func(concept.Param, concept.Object) concept.Param,
	paramNames []concept.String,
	returnNames []concept.String,
) *SystemFunction {
	return &SystemFunction{
		AdaptorFunction: adaptor.NewAdaptorFunction(),
		name:            name,
		funcs:           funcs,
		anticipateFuncs: anticipateFuncs,
		paramNames:      paramNames,
		returnNames:     returnNames,
		seed:            s,
	}
}

func (s *SystemFunctionCreator) ToLanguage(language string, instance *SystemFunction) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *SystemFunctionCreator) Type() string {
	return VariableSystemFunctionType
}

func NewSystemFunctionCreator(param *SystemFunctionCreatorParam) *SystemFunctionCreator {
	return &SystemFunctionCreator{
		Seeds: map[string]func(string, *SystemFunction) string{},
		param: param,
	}
}
