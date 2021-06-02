package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableSystemFunctionType = "system_function"
	FunctionSystemFunctionType = "system"
)

type SystemFunctionSeed interface {
	ToLanguage(string, concept.Pool, *SystemFunction) string
	Type() string
	NewParam() concept.Param
}

type SystemFunction struct {
	*adaptor.AdaptorFunction
	funcs           func(concept.Param, concept.Variable) (concept.Param, concept.Exception)
	anticipateFuncs func(concept.Param, concept.Variable) concept.Param
	seed            SystemFunctionSeed
}

func (o *SystemFunction) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *SystemFunction) ToLanguage(language string, space concept.Pool) string {
	return f.seed.ToLanguage(language, space, f)
}

func (f *SystemFunction) ToCallLanguage(language string, space concept.Pool, self string, param concept.Param) string {
	return f.ToCallLanguageAdaptor(f, language, space, self, param)
}

func (f *SystemFunction) ToString(prefix string) string {
	return fmt.Sprintf("system_function (%v) %v {}", StringJoin(f.ParamNames(), ", "), StringJoin(f.ReturnNames(), ", "))
}

func (f *SystemFunction) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	if f.anticipateFuncs == nil {
		back, suspend := f.Exec(params, object)
		if !nl_interface.IsNil(suspend) {
			return f.seed.NewParam()
		}
		return back
	}
	return f.anticipateFuncs(f.paramFormat(params), object)
}

func (f *SystemFunction) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	return f.funcs(f.paramFormat(params), object)
}

func (f *SystemFunction) paramFormat(params concept.Param) concept.Param {
	if params.ParamType() == concept.ParamTypeList {
		for index, name := range f.ParamNames() {
			if index < params.SizeIndex() {
				params.Set(name, params.GetIndex(index))
			}
		}
	}
	return params
}

func (s *SystemFunction) Type() string {
	return s.seed.Type()
}

func (s *SystemFunction) FunctionType() string {
	return FunctionSystemFunctionType
}

type SystemFunctionCreatorParam struct {
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

type SystemFunctionCreator struct {
	Seeds map[string]func(string, concept.Pool, *SystemFunction) string
	Inits []func(*SystemFunction)
	param *SystemFunctionCreatorParam
}

func (s *SystemFunctionCreator) New(
	funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
	anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
	paramNames []concept.String,
	returnNames []concept.String,
) *SystemFunction {
	system := &SystemFunction{
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
		funcs:           funcs,
		anticipateFuncs: anticipateFuncs,
		seed:            s,
	}
	system.AddParamName(paramNames...)
	system.AddReturnName(returnNames...)

	for _, init := range s.Inits {
		init(system)
	}

	return system
}

func (s *SystemFunctionCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *SystemFunctionCreator) ToLanguage(language string, space concept.Pool, instance *SystemFunction) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *SystemFunctionCreator) Type() string {
	return VariableSystemFunctionType
}

func NewSystemFunctionCreator(param *SystemFunctionCreatorParam) *SystemFunctionCreator {
	return &SystemFunctionCreator{
		Seeds: map[string]func(string, concept.Pool, *SystemFunction) string{},
		param: param,
	}
}
