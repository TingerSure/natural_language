package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableDefineFunctionType = "define_function"
	FunctionDefineFunctionType = "define"
)

type DefineFunctionSeed interface {
	ToLanguage(string, concept.Closure, *DefineFunction) string
	Type() string
	NewParam() concept.Param
	NewNull() concept.Null
	NewException(string, string) concept.Exception
}

type DefineFunction struct {
	*adaptor.AdaptorFunction
	paramNames  []concept.String
	returnNames []concept.String
	seed        DefineFunctionSeed
}

func (f *DefineFunction) ParamFormat(params concept.Param) concept.Param {
	return f.AdaptorFunction.AdaptorParamFormat(f, params)
}

func (f *DefineFunction) ReturnFormat(back concept.String) concept.String {
	return f.AdaptorFunction.AdaptorReturnFormat(f, back)
}

func (o *DefineFunction) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *DefineFunction) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (f *DefineFunction) ToCallLanguage(language string, space concept.Closure, self string, param concept.Param) string {
	return f.ToCallLanguageAdaptor(f, language, space, self, param)
}

func (f *DefineFunction) AddParamName(paramNames ...concept.String) {
	f.paramNames = append(f.paramNames, paramNames...)
}

func (f *DefineFunction) AddReturnName(returnNames ...concept.String) {
	f.returnNames = append(f.returnNames, returnNames...)
}

func (s *DefineFunction) ParamNames() []concept.String {
	return s.paramNames
}

func (s *DefineFunction) ReturnNames() []concept.String {
	return s.returnNames
}

func (f *DefineFunction) ToString(prefix string) string {
	return fmt.Sprintf("function (%v) %v", StringJoin(f.paramNames, ", "), StringJoin(f.returnNames, ", "))
}

func (f *DefineFunction) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	return f.seed.NewParam()
}

func (f *DefineFunction) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	return nil, f.seed.NewException("runtime err", "define_function cannot be executed directly.")
}

func (f *DefineFunction) paramFormat(params concept.Param) concept.Param {
	if params.ParamType() == concept.ParamTypeList {
		for index, name := range f.paramNames {
			if index < params.SizeIndex() {
				params.Set(name, params.GetIndex(index))
			}
		}
	}
	return params
}

func (s *DefineFunction) Type() string {
	return s.seed.Type()
}

func (s *DefineFunction) FunctionType() string {
	return FunctionDefineFunctionType
}

type DefineFunctionCreatorParam struct {
	NullCreator      func() concept.Null
	ParamCreator     func() concept.Param
	ExceptionCreator func(string, string) concept.Exception
}

type DefineFunctionCreator struct {
	Seeds map[string]func(string, concept.Closure, *DefineFunction) string
	param *DefineFunctionCreatorParam
}

func (s *DefineFunctionCreator) New() *DefineFunction {
	return &DefineFunction{
		AdaptorFunction: adaptor.NewAdaptorFunction(&adaptor.AdaptorFunctionParam{
			NullCreator:      s.param.NullCreator,
			ParamCreator:     s.param.ParamCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		seed: s,
	}
}

func (s *DefineFunctionCreator) ToLanguage(language string, space concept.Closure, instance *DefineFunction) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *DefineFunctionCreator) Type() string {
	return VariableDefineFunctionType
}

func (s *DefineFunctionCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *DefineFunctionCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *DefineFunctionCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewDefineFunctionCreator(param *DefineFunctionCreatorParam) *DefineFunctionCreator {
	return &DefineFunctionCreator{
		Seeds: map[string]func(string, concept.Closure, *DefineFunction) string{},
		param: param,
	}
}
