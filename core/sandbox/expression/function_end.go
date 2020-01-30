package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type FunctionEnd struct {
	*adaptor.ExpressionIndex
}

func (a *FunctionEnd) ToString(prefix string) string {
	return fmt.Sprintf("end")
}

func (a *FunctionEnd) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, interrupt.NewEnd()
}

func NewFunctionEnd() *FunctionEnd {
	back := &FunctionEnd{}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
