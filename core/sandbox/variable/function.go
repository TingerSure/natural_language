package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
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
	ToLanguage(string, *Function) string
	Type() string
	NewString(string) concept.String
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewNull() concept.Null
}

type Function struct {
	*adaptor.AdaptorFunction
	body           *code_block.CodeBlock
	anticipateBody *code_block.CodeBlock
	paramNames     []concept.String
	returnNames    []concept.String
	parent         concept.Closure
	seed           FunctionSeed
}

func (f *Function) ParamFormat(params *concept.Mapping) *concept.Mapping {
	return f.AdaptorFunction.AdaptorParamFormat(f, params)
}

func (f *Function) ReturnFormat(back concept.String) concept.String {
	return f.AdaptorFunction.AdaptorReturnFormat(f, back)
}

func (o *Function) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Function) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *Function) ParamNames() []concept.String {
	return s.paramNames
}

func (s *Function) ReturnNames() []concept.String {
	return s.returnNames
}

func (s *Function) FunctionType() string {
	return FunctionFunctionType
}

func (f *Function) ToString(prefix string) string {
	return fmt.Sprintf("function (%v) %v %v", StringJoin(f.paramNames, ", "), StringJoin(f.returnNames, ", "), f.body.ToString(prefix))
}

func (f *Function) AddParamName(paramNames ...concept.String) {
	f.paramNames = append(f.paramNames, paramNames...)
}

func (f *Function) AddReturnName(returnNames ...concept.String) {
	f.returnNames = append(f.returnNames, returnNames...)
}

func (f *Function) AnticipateBody() *code_block.CodeBlock {
	return f.anticipateBody
}

func (f *Function) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	space, suspend := f.anticipateBody.Exec(f.parent, func(space concept.Closure) concept.Interrupt {
		space.InitLocal(f.seed.NewString(FunctionAutoParamSelf), f)
		space.InitLocal(f.seed.NewString(FunctionAutoParamThis), object)
		if params.ParamType() == concept.ParamTypeList {
			for index, name := range f.paramNames {
				if index < params.SizeIndex() {
					space.InitLocal(name, params.GetIndex(index))
				} else {
					space.InitLocal(name, params.Get(name))
				}
			}
		}
		if params.ParamType() == concept.ParamTypeKeyValue {
			for _, name := range f.paramNames {
				space.InitLocal(name, params.Get(name))
			}
		}
		for _, name := range f.returnNames {
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
	for _, name := range f.returnNames {
		value, _ := space.PeekLocal(name)
		returnParams.Set(name, value)
	}
	return returnParams
}

func (f *Function) Body() *code_block.CodeBlock {
	return f.body
}

func (f *Function) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	space, suspend := f.body.Exec(f.parent, func(space concept.Closure) concept.Interrupt {
		space.InitLocal(f.seed.NewString(FunctionAutoParamSelf), f)
		space.InitLocal(f.seed.NewString(FunctionAutoParamThis), object)
		if params.ParamType() == concept.ParamTypeList {
			for index, name := range f.paramNames {
				if index < params.SizeIndex() {
					space.InitLocal(name, params.GetIndex(index))
				} else {
					space.InitLocal(name, params.Get(name))
				}
			}
		}
		if params.ParamType() == concept.ParamTypeKeyValue {
			for _, name := range f.paramNames {
				space.InitLocal(name, params.Get(name))
			}
		}
		for _, name := range f.returnNames {
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
	for _, name := range f.returnNames {
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
	CodeBlockCreator      func() *code_block.CodeBlock
	StringCreator         func(string) concept.String
	ParamCreator          func() concept.Param
	ExceptionCreator      func(string, string) concept.Exception
	NullCreator           func() concept.Null
	DelayStringCreator    func(string) concept.String
	DelayFunctionCreator  func(func() concept.Function) concept.Function
	ArrayCreator          func() *Array
	SystemFunctionCreator func(
		funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
		anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
		paramNames []concept.String,
		returnNames []concept.String,
	) concept.Function
}

type FunctionCreator struct {
	Seeds map[string]func(string, *Function) string
	param *FunctionCreatorParam
}

func (s *FunctionCreator) New(parent concept.Closure) *Function {
	funcs := &Function{
		AdaptorFunction: adaptor.NewAdaptorFunction(&adaptor.AdaptorFunctionParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		parent:         parent,
		body:           s.param.CodeBlockCreator(),
		anticipateBody: s.param.CodeBlockCreator(),
		seed:           s,
	}

	funcs.SetField(s.param.DelayStringCreator("paramList"), s.param.DelayFunctionCreator(s.FieldParamList(funcs)))
	funcs.SetField(s.param.DelayStringCreator("returnList"), s.param.DelayFunctionCreator(s.FieldReturnList(funcs)))
	return funcs
}

func (s *FunctionCreator) FieldParamList(funcs concept.Function) func() concept.Function {
	return func() concept.Function {
		backList := s.param.StringCreator("list")
		return s.param.SystemFunctionCreator(
			func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				paramNames := s.param.ArrayCreator()
				for _, paramName := range funcs.ParamNames() {
					paramNames.Append(paramName)
				}
				back := s.param.ParamCreator()
				back.Set(backList, paramNames)
				return back, nil
			},
			func(param concept.Param, _ concept.Variable) concept.Param {
				paramNames := s.param.ArrayCreator()
				for _, paramName := range funcs.ParamNames() {
					paramNames.Append(paramName)
				}
				back := s.param.ParamCreator()
				back.Set(backList, paramNames)
				return back
			},
			[]concept.String{},
			[]concept.String{
				backList,
			},
		)
	}
}

func (s *FunctionCreator) FieldReturnList(funcs concept.Function) func() concept.Function {
	return func() concept.Function {
		backList := s.param.StringCreator("list")
		return s.param.SystemFunctionCreator(
			func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				returnNames := s.param.ArrayCreator()
				for _, returnName := range funcs.ReturnNames() {
					returnNames.Append(returnName)
				}
				back := s.param.ParamCreator()
				back.Set(backList, returnNames)
				return back, nil
			},
			func(param concept.Param, _ concept.Variable) concept.Param {
				returnNames := s.param.ArrayCreator()
				for _, returnName := range funcs.ReturnNames() {
					returnNames.Append(returnName)
				}
				back := s.param.ParamCreator()
				back.Set(backList, returnNames)
				return back
			},
			[]concept.String{},
			[]concept.String{
				backList,
			},
		)
	}
}

func (s *FunctionCreator) ToLanguage(language string, instance *Function) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
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
		Seeds: map[string]func(string, *Function) string{},
		param: param,
	}
}
