package sandbox

import (
	"fmt"
)

type FunctionEnd struct {
}

func (a *FunctionEnd) ToString(prefix string) string {
	return fmt.Sprintf("%vend", prefix)
}

func (a *FunctionEnd) Exec(space *Closure) Interrupt {
	return NewEnd()
}

func NewFunctionEnd() *FunctionEnd {
	return &FunctionEnd{}
}
