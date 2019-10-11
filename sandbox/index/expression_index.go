package index

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
)

type ExpressionIndex struct {
	flow concept.Expression
}

func (s *ExpressionIndex) ToString(prefix string) string {
	return s.flow.ToString(prefix)
}

func (s *ExpressionIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.flow.Exec(space)
}

func (s *ExpressionIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException("read only", "Expression result does not need to be changed.")
}

func NewExpressionIndex(flow concept.Expression) *ExpressionIndex {
	return &ExpressionIndex{
		flow: flow,
	}
}
