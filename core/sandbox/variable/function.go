package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

const (
	VariableFunctionType  = "function"
	FunctionFunctionType  = "general"
	FunctionAutoParamSelf = "self"
	FunctionAutoParamThis = "this"
)

type Function struct {
	body       *code_block.CodeBlock
	paramNames []concept.String
	parent     concept.Closure
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

func (f *Function) Body() *code_block.CodeBlock {
	return f.body
}

func (f *Function) Exec(params concept.Param, object concept.Object) (concept.Param, concept.Exception) {

	space, suspend := f.body.Exec(f.parent, false, func(space concept.Closure) concept.Interrupt {
		space.InitLocal(NewString(FunctionAutoParamSelf), f)
		space.InitLocal(NewString(FunctionAutoParamThis), object)
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
				return nil, interrupt.NewException(NewString("system panic"), NewString(fmt.Sprintf("ExceptionInterruptType does not mean an Exception anymore.\n%+v", suspend)))
			}
			return nil, exception
		case interrupt.EndInterruptType:
			return NewParamWithIterate(space.IterateReturn), nil
		default:
			return nil, interrupt.NewException(NewString("system error"), NewString(fmt.Sprintf("Unknown Interrupt \"%v\".\n%+v", suspend.InterruptType(), suspend)))
		}
	}

	return NewParamWithIterate(space.IterateReturn), nil
}

func (s *Function) Type() string {
	return VariableFunctionType
}

func NewFunction(parent concept.Closure) *Function {

	return &Function{
		parent: parent,
		body: code_block.NewCodeBlock(&code_block.CodeBlockParam{
			StringCreator: func(value string) concept.String {
				return NewString(value)
			},
			EmptyCreator: func() concept.Null {
				return NewNull()
			},
		}),
	}
}
