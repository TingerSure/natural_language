package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
)

const (
	VariableFunctionType  = "function"
	FunctionFunctionType  = "general"
	FunctionAutoParamSelf = "self"
	FunctionAutoParamThis = "this"
)

type FunctionSeed interface {
	ToLanguage(string, concept.Pool, *Function) string
	Type() string
	NewString(string) concept.String
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type Function struct {
	*adaptor.AdaptorFunction
	body           concept.CodeBlock
	anticipateBody concept.CodeBlock
	parent         concept.Pool
	seed           FunctionSeed
}

func (o *Function) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Function) ToLanguage(language string, space concept.Pool) string {
	return f.body.ToLanguage(language, f.parent)
}

func (f *Function) ToCallLanguage(language string, space concept.Pool, self string, param concept.Param) string {
	return f.ToCallLanguageAdaptor(f, language, space, self, param)
}

func (s *Function) FunctionType() string {
	return FunctionFunctionType
}

func (f *Function) ToString(prefix string) string {
	return fmt.Sprintf("function (%v) %v %v", StringJoin(f.ParamNames(), ", "), StringJoin(f.ReturnNames(), ", "), f.body.ToString(prefix))
}

func (f *Function) AnticipateBody() concept.CodeBlock {
	return f.anticipateBody
}

func (f *Function) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	space, suspend := f.anticipateBody.ExecWithInit(f.parent, func(space concept.Pool) concept.Interrupt {
		space.InitLocal(f.seed.NewString(FunctionAutoParamSelf), f)
		space.InitLocal(f.seed.NewString(FunctionAutoParamThis), object)
		if params.ParamType() == concept.ParamTypeList {
			for index, name := range f.ParamNames() {
				if index < params.SizeIndex() {
					space.InitLocal(name, params.GetIndex(index))
				} else {
					space.InitLocal(name, params.Get(name))
				}
			}
		}
		if params.ParamType() == concept.ParamTypeKeyValue {
			for _, name := range f.ParamNames() {
				space.InitLocal(name, params.Get(name))
			}
		}
		for _, name := range f.ReturnNames() {
			space.InitLocal(name, f.seed.NewNull())
		}
		return nil
	})
	defer space.Clear()

	if !nl_interface.IsNil(suspend) {
		switch suspend.InterruptType() {
		case ExceptionInterruptType:
			return f.seed.NewParam()
		case interrupt.ReturnInterruptType:
			//Do Nothing
		default:
			return f.seed.NewParam()
		}
	}
	returnParams := f.seed.NewParam()
	for _, name := range f.ReturnNames() {
		value, _ := space.PeekLocal(name)
		returnParams.Set(name, value)
	}
	return returnParams
}

func (f *Function) Body() concept.CodeBlock {
	return f.body
}

func (f *Function) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	space, suspend := f.body.ExecWithInit(f.parent, func(space concept.Pool) concept.Interrupt {
		space.InitLocal(f.seed.NewString(FunctionAutoParamSelf), f)
		space.InitLocal(f.seed.NewString(FunctionAutoParamThis), object)
		if params.ParamType() == concept.ParamTypeList {
			for index, name := range f.ParamNames() {
				if index < params.SizeIndex() {
					space.InitLocal(name, params.GetIndex(index))
				} else {
					space.InitLocal(name, params.Get(name))
				}
			}
		}
		if params.ParamType() == concept.ParamTypeKeyValue {
			for _, name := range f.ParamNames() {
				space.InitLocal(name, params.Get(name))
			}
		}
		for _, name := range f.ReturnNames() {
			space.InitLocal(name, f.seed.NewNull())
		}
		return nil
	})
	defer space.Clear()

	if !nl_interface.IsNil(suspend) {
		switch suspend.InterruptType() {
		case ExceptionInterruptType:
			exception, yes := interrupt.InterruptFamilyInstance.IsException(suspend)
			if !yes {
				return nil, f.seed.NewException("system panic", fmt.Sprintf("ExceptionInterruptType does not mean an Exception anymore.\n%+v", suspend))
			}
			return nil, exception
		case interrupt.ReturnInterruptType:
			// Do Nothing
		default:
			return nil, f.seed.NewException("system error", fmt.Sprintf("Unknown Interrupt \"%v\".\n%+v", suspend.InterruptType(), suspend))
		}
	}
	returnParams := f.seed.NewParam()
	for _, name := range f.ReturnNames() {
		value, returnSuspend := space.PeekLocal(name)
		if !nl_interface.IsNil(returnSuspend) {
			return nil, returnSuspend
		}
		returnParams.Set(name, value)
	}
	return returnParams, nil
}

func (s *Function) Type() string {
	return s.seed.Type()
}

type FunctionCreatorParam struct {
	CodeBlockCreator      func() concept.CodeBlock
	StringCreator         func(string) concept.String
	ParamCreator          func() concept.Param
	ExceptionCreator      func(string, string) concept.Exception
	NullCreator           func() concept.Null
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

type FunctionCreator struct {
	Seeds map[string]func(string, concept.Pool, *Function) string
	Inits []func(*Function)
	param *FunctionCreatorParam
}

func (s *FunctionCreator) New(parent concept.Pool) *Function {
	funcs := &Function{
		AdaptorFunction: adaptor.NewAdaptorFunction(&adaptor.AdaptorFunctionParam{
			NullCreator:           s.param.NullCreator,
			ExceptionCreator:      s.param.ExceptionCreator,
			ParamCreator:          s.param.ParamCreator,
			SystemFunctionCreator: s.param.SystemFunctionCreator,
			ArrayCreator:          s.param.ArrayCreator,
			DelayFunctionCreator:  s.param.DelayFunctionCreator,
			DelayStringCreator:    s.param.DelayStringCreator,
			StringCreator:         s.param.StringCreator,
		}),
		parent:         parent,
		body:           s.param.CodeBlockCreator(),
		anticipateBody: s.param.CodeBlockCreator(),
		seed:           s,
	}

	for _, init := range s.Inits {
		init(funcs)
	}

	return funcs
}

func (s *FunctionCreator) ToLanguage(language string, space concept.Pool, instance *Function) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *FunctionCreator) Type() string {
	return VariableFunctionType
}

func (s *FunctionCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func (s *FunctionCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *FunctionCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *FunctionCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewFunctionCreator(param *FunctionCreatorParam) *FunctionCreator {
	return &FunctionCreator{
		Seeds: map[string]func(string, concept.Pool, *Function) string{},
		param: param,
	}
}
