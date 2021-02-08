package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariablePreObjectFunctionType = "pre_object_function"
	FunctionPreObjectFunctionType = "pre_object"
)

type PreObjectFunctionSeed interface {
	ToLanguage(string, *PreObjectFunction) string
	Type() string
}

type PreObjectFunction struct {
	*adaptor.AdaptorFunction
	function concept.Function
	object   concept.Variable
	seed     PreObjectFunctionSeed
}

func (f *PreObjectFunction) ParamFormat(params *concept.Mapping) *concept.Mapping {
	return f.AdaptorFunction.AdaptorParamFormat(f, params)
}

func (f *PreObjectFunction) ReturnFormat(back concept.String) concept.String {
	return f.AdaptorFunction.AdaptorReturnFormat(f, back)
}

func (f *PreObjectFunction) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *PreObjectFunction) ParamNames() []concept.String {
	return s.function.ParamNames()
}

func (s *PreObjectFunction) ReturnNames() []concept.String {
	return s.function.ReturnNames()
}

func (f *PreObjectFunction) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", f.object.ToString(prefix), f.function.ToString(prefix))
}

func (f *PreObjectFunction) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	if nl_interface.IsNil(object) {
		object = f.object
	}
	return f.function.Anticipate(params, object)
}

func (f *PreObjectFunction) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	if nl_interface.IsNil(object) {
		object = f.object
	}
	return f.function.Exec(params, object)
}

func (s *PreObjectFunction) Type() string {
	return s.seed.Type()
}

func (s *PreObjectFunction) FunctionType() string {
	return FunctionPreObjectFunctionType
}

func (s *PreObjectFunction) Name() concept.String {
	return s.function.Name()
}

type PreObjectFunctionCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type PreObjectFunctionCreator struct {
	Seeds map[string]func(string, *PreObjectFunction) string
	param *PreObjectFunctionCreatorParam
}

func (s *PreObjectFunctionCreator) New(
	function concept.Function,
	object concept.Variable,
) *PreObjectFunction {
	return &PreObjectFunction{
		AdaptorFunction: adaptor.NewAdaptorFunction(&adaptor.AdaptorFunctionParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		function: function,
		object:   object,
		seed:     s,
	}
}

func (s *PreObjectFunctionCreator) ToLanguage(language string, instance *PreObjectFunction) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *PreObjectFunctionCreator) Type() string {
	return VariablePreObjectFunctionType
}

func NewPreObjectFunctionCreator(param *PreObjectFunctionCreatorParam) *PreObjectFunctionCreator {
	return &PreObjectFunctionCreator{
		Seeds: map[string]func(string, *PreObjectFunction) string{},
		param: param,
	}
}
