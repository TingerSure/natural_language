package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

var (
	ExpressionFamilyInstance *ExpressionFamily = newExpressionFamily()
)

type ExpressionFamily struct {
}

func newExpressionFamily() *ExpressionFamily {
	return &ExpressionFamily{}
}

func (v *ExpressionFamily) IsComponent(value concept.Pipe) (*Component, bool) {
	if value == nil {
		return nil, false
	}
	component, yes := value.(*Component)
	return component, yes
}
