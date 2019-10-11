package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
)

type FunctionEnd struct {
}

func (a *FunctionEnd) ToString(prefix string) string {
	return fmt.Sprintf("end")
}

func (a *FunctionEnd) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, interrupt.NewEnd()
}

func NewFunctionEnd() *FunctionEnd {
	return &FunctionEnd{}
}
