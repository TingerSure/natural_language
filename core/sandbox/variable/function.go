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
}

type Function struct {
	*adaptor.AdaptorFunction
	name           concept.String
	body           *code_block.CodeBlock
	anticipateBody *code_block.CodeBlock
	paramNames     []concept.String
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
	names := []concept.String{}
	s.body.Iterate(func(index concept.Index) bool {
		end, ok := index.(concept.Return)
		if ok {
			names = append(names, end.Key())
		}
		return false
	})
	return names
}

func (s *Function) FunctionType() string {
	return FunctionFunctionType
}

func (f *Function) ToString(prefix string) string {
	return fmt.Sprintf("function (%v) %v", StringJoin(f.paramNames, ", "), f.body.ToString(prefix))
}

func (f *Function) AddParamName(paramName concept.String) {
	f.paramNames = append(f.paramNames, paramName)
}

func (f *Function) AnticipateBody() *code_block.CodeBlock {
	return f.anticipateBody
}

func (f *Function) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	space, suspend := f.anticipateBody.Exec(f.parent, false, func(space concept.Closure) concept.Interrupt {
		space.InitLocal(f.seed.NewString(FunctionAutoParamSelf), f)
		space.InitLocal(f.seed.NewString(FunctionAutoParamThis), object)
		for _, name := range f.paramNames {
			space.InitLocal(name, params.Get(name))
		}
		return nil
	})
	defer space.Clear()

	if !nl_interface.IsNil(suspend) {
		switch suspend.InterruptType() {
		case interrupt.ExceptionInterruptType:
			return f.seed.NewParam()
		case interrupt.EndInterruptType:
			return f.seed.NewParam().Init(space.IterateReturn)
		default:
			return f.seed.NewParam()
		}
	}

	return f.seed.NewParam().Init(space.IterateReturn)
}

func (f *Function) Body() *code_block.CodeBlock {
	return f.body
}

func (f *Function) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {

	space, suspend := f.body.Exec(f.parent, false, func(space concept.Closure) concept.Interrupt {
		space.InitLocal(f.seed.NewString(FunctionAutoParamSelf), f)
		space.InitLocal(f.seed.NewString(FunctionAutoParamThis), object)
		for _, name := range f.paramNames {
			space.InitLocal(name, params.Get(name))
		}
		return nil
	})
	defer space.Clear()

	if !nl_interface.IsNil(suspend) {
		switch suspend.InterruptType() {
		case interrupt.ExceptionInterruptType:
			exception, yes := interrupt.InterruptFamilyInstance.IsException(suspend)
			if !yes {
				return nil, f.seed.NewException("system panic", fmt.Sprintf("ExceptionInterruptType does not mean an Exception anymore.\n%+v", suspend))
			}
			return nil, exception
		case interrupt.EndInterruptType:
			return f.seed.NewParam().Init(space.IterateReturn), nil
		default:
			return nil, f.seed.NewException("system error", fmt.Sprintf("Unknown Interrupt \"%v\".\n%+v", suspend.InterruptType(), suspend))
		}
	}

	return f.seed.NewParam().Init(space.IterateReturn), nil
}

func (s *Function) Type() string {
	return s.seed.Type()
}

func (s *Function) Name() concept.String {
	return s.name
}

type FunctionCreatorParam struct {
	CodeBlockCreator func() *code_block.CodeBlock
	StringCreator    func(string) concept.String
	ParamCreator     func() concept.Param
	ExceptionCreator func(string, string) concept.Exception
	NullCreator      func() concept.Null
}

type FunctionCreator struct {
	Seeds map[string]func(string, *Function) string
	param *FunctionCreatorParam
}

func (s *FunctionCreator) New(name concept.String, parent concept.Closure) *Function {
	return &Function{
		AdaptorFunction: adaptor.NewAdaptorFunction(&adaptor.AdaptorFunctionParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		name:           name,
		parent:         parent,
		body:           s.param.CodeBlockCreator(),
		anticipateBody: s.param.CodeBlockCreator(),
		seed:           s,
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
