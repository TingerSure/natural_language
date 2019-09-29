package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/code_block"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"strings"
)

const (
	VariableFunctionType = "function"
)

type Function struct {
	body       *code_block.CodeBlock
	paramNames []string
	parent     concept.Closure
}

func (f *Function) ToString(prefix string) string {
	return fmt.Sprintf("function (%v) %v", strings.Join(f.paramNames, ", "), f.body.ToString(prefix))
}

func (f *Function) AddParamName(paramName string) {
	f.paramNames = append(f.paramNames, paramName)
}

func (f *Function) AddStep(step concept.Expression) {
	f.body.AddStep(step)
}

func (f *Function) Exec(params *Param) (*Param, *interrupt.Exception) {

	space, suspend := f.body.Exec(f.parent, false, func(space concept.Closure) concept.Interrupt {
		for _, name := range f.paramNames {
			space.InitLocal(name)
			suspend := space.SetLocal(name, params.Get(name))
			if !nl_interface.IsNil(suspend) {
				return suspend
			}
		}
		// for name, value := range f.paramNames {
		// 	suspend := space.SetLocal(name, value)
		// 	if !nl_interface.IsNil(suspend) {
		// 		return suspend
		// 	}
		// }
		return nil
	})
	defer space.Clear()

	if !nl_interface.IsNil(suspend) {
		switch suspend.InterruptType() {
		case interrupt.ExceptionInterruptType:
			exception, yes := interrupt.InterruptFamilyInstance.IsException(suspend)
			if !yes {
				return nil, interrupt.NewException("system panic", fmt.Sprintf("ExceptionInterruptType does not mean an Exception anymore.\n%+v", suspend))
			}
			return nil, exception
		case interrupt.EndInterruptType:
			return NewParamWithInit(space.Return()), nil
		default:
			return nil, interrupt.NewException("system error", fmt.Sprintf("Unknown Interrupt \"%v\".\n%+v", suspend.InterruptType(), suspend))
		}
	}

	return NewParamWithInit(space.Return()), nil
}

func (s *Function) Type() string {
	return VariableFunctionType
}

func NewFunction(parent concept.Closure) *Function {
	return &Function{
		parent: parent,
		body:   code_block.NewCodeBlock(),
	}
}
