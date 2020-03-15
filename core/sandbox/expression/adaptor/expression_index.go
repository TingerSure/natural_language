package adaptor

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ExpressionIndex struct {
	exec func(concept.Closure) (concept.Variable, concept.Interrupt)
}

func (e *ExpressionIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (e *ExpressionIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return e.exec(space)
}

func (e *ExpressionIndex) Set(concept.Closure, concept.Variable) concept.Interrupt {
	return interrupt.NewException(variable.NewString("read only"), variable.NewString("Expression result does not need to be changed."))
}

func NewExpressionIndex(exec func(concept.Closure) (concept.Variable, concept.Interrupt)) *ExpressionIndex {
	return &ExpressionIndex{
		exec: exec,
	}
}
