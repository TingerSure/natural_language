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
	ToLanguage(string, concept.Pool, *DefineFunction) string
	Type() string
	NewParam() concept.Param
	NewNull() concept.Null
	NewException(string, string) concept.Exception
}

type DefineFunction struct {
	*adaptor.AdaptorFunction
	seed DefineFunctionSeed
}

func (o *DefineFunction) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *DefineFunction) ToLanguage(language string, space concept.Pool) string {
	return f.seed.ToLanguage(language, space, f)
}

func (f *DefineFunction) ToCallLanguage(language string, space concept.Pool, self string, param concept.Param) string {
	return f.ToCallLanguageAdaptor(f, language, space, self, param)
}

func (f *DefineFunction) ToString(prefix string) string {
	return fmt.Sprintf("function (%v) %v", StringJoin(f.ParamNames(), ", "), StringJoin(f.ReturnNames(), ", "))
}

func (f *DefineFunction) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	return f.seed.NewParam()
}

func (f *DefineFunction) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	return nil, f.seed.NewException("runtime err", "define_function cannot be executed directly.")
}

func (s *DefineFunction) Type() string {
	return s.seed.Type()
}

func (s *DefineFunction) FunctionType() string {
	return FunctionDefineFunctionType
}

type DefineFunctionCreatorParam struct {
	NullCreator           func() concept.Null
	ParamCreator          func() concept.Param
	ExceptionCreator      func(string, string) concept.Exception
	StringCreator         func(string) concept.String
	DelayStringCreator    func(string) concept.String
	DelayFunctionCreator  func(func() concept.Function) concept.Function
	ArrayCreator          func() concept.Array
	SystemFunctionCreator func(
		funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
		anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
		paramNames []concept.String,
		returnNames []concept.String,
	) concept.Function
}

type DefineFunctionCreator struct {
	Seeds map[string]func(string, concept.Pool, *DefineFunction) string
	param *DefineFunctionCreatorParam
}

func (s *DefineFunctionCreator) New(paramNames []concept.String, returnNames []concept.String) *DefineFunction {
	define := &DefineFunction{
		AdaptorFunction: adaptor.NewAdaptorFunction(&adaptor.AdaptorFunctionParam{
			NullCreator:           s.param.NullCreator,
			ParamCreator:          s.param.ParamCreator,
			ExceptionCreator:      s.param.ExceptionCreator,
			SystemFunctionCreator: s.param.SystemFunctionCreator,
			ArrayCreator:          s.param.ArrayCreator,
			DelayFunctionCreator:  s.param.DelayFunctionCreator,
			DelayStringCreator:    s.param.DelayStringCreator,
			StringCreator:         s.param.StringCreator,
		}),
		seed: s,
	}

	define.AddParamName(paramNames...)
	define.AddReturnName(returnNames...)

	return define
}

func (s *DefineFunctionCreator) ToLanguage(language string, space concept.Pool, instance *DefineFunction) string {
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
		Seeds: map[string]func(string, concept.Pool, *DefineFunction) string{},
		param: param,
	}
}
