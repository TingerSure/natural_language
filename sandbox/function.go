package sandbox

import (
	"fmt"
	"strings"
)

const (
	VariableFunctionType = "function"
)

type Function struct {
	body       *CodeBlock
	paramNames []string
	parent     *Closure
}

func (f *Function) ToString(prefix string) string {
	return fmt.Sprintf("function (%v) %v", strings.Join(f.paramNames, ", "), f.body.ToString(prefix))
}

func (f *Function) AddParamName(paramName string) {
	f.paramNames = append(f.paramNames, paramName)
}

func (f *Function) AddStep(step Expression) {
	f.body.AddStep(step)
}

func (f *Function) Exec(params map[string]Variable) (map[string]Variable, *Exception) {

	space, suspend := f.body.Exec(f.parent, false, func(space *Closure) Interrupt {
		for _, name := range f.paramNames {
			space.InitLocal(name)
		}
		for name, value := range params {
			suspend := space.SetLocal(name, value)
			if suspend != nil {
				return suspend
			}
		}
		return nil
	})
	defer space.Clear()

	if suspend != nil {
		switch suspend.InterruptType() {
		case ExceptionInterruptType:
			exception, yes := InterruptFamilyInstance.IsException(suspend)
			if !yes {
				return nil, NewException("system panic", fmt.Sprintf("ExceptionInterruptType does not mean an Exception anymore.\n%+v", suspend))
			}
			return nil, exception
		case EndInterruptType:
			return space.Return(), nil
		default:
			return nil, NewException("system error", fmt.Sprintf("Unknown Interrupt \"%v\".\n%+v", suspend.InterruptType(), suspend))
		}
	}

	return space.Return(), nil
}

func (s *Function) Type() string {
	return VariableFunctionType
}

func NewFunction(parent *Closure) *Function {
	return &Function{
		parent: parent,
		body:   NewCodeBlock(),
	}
}
