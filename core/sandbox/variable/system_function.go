package variable

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableSystemFunctionType = "system_function"
	FunctionSystemFunctionType = "system"
)

type SystemFunction struct {
	*adaptor.AdaptorFunction
	name        concept.String
	paramNames  []concept.String
	returnNames []concept.String
	funcs       func(concept.Param, concept.Object) (concept.Param, concept.Exception)
}

var (
	SystemFunctionLanguageSeeds = map[string]func(string, *SystemFunction) string{}
)

func (f *SystemFunction) ToLanguage(language string) string {
	seed := SystemFunctionLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
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

func (f *SystemFunction) Exec(params concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	return f.funcs(params, object)
}

func (s *SystemFunction) Type() string {
	return VariableSystemFunctionType
}

func (s *SystemFunction) FunctionType() string {
	return FunctionSystemFunctionType
}

func NewSystemFunction(
	name concept.String,
	funcs func(concept.Param, concept.Object) (concept.Param, concept.Exception),
	paramNames []concept.String,
	returnNames []concept.String,
) *SystemFunction {
	return &SystemFunction{
		AdaptorFunction: adaptor.NewAdaptorFunction(),
		name:            name,
		funcs:           funcs,
		paramNames:      paramNames,
		returnNames:     returnNames,
	}
}
