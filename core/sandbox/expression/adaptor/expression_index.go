package adaptor

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type ExpressionIndex struct {
	exec func(concept.Closure) (concept.Variable, concept.Interrupt)
}

var (
	IndexExpressionType = "Expression"
)

func (e *ExpressionIndex) Type() string {
	return IndexExpressionType
}

func (e *ExpressionIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (e *ExpressionIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return e.exec(space)
}

func (e *ExpressionIndex) Set(concept.Closure, concept.Variable) concept.Interrupt {
	return interrupt.NewException("read only", "Expression result does not need to be changed.")
}

func NewExpressionIndex(exec func(concept.Closure) (concept.Variable, concept.Interrupt)) *ExpressionIndex {
	return &ExpressionIndex{
		exec: exec,
	}
}
