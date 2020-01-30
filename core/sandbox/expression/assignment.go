package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type Assignment struct {
	*adaptor.ExpressionIndex
	from concept.Index
	to   concept.Index
}

func (a *Assignment) ToString(prefix string) string {
	return fmt.Sprintf("%v = %v", a.to.ToString(prefix), a.from.ToString(prefix))
}

func (a *Assignment) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preFrom, suspend := a.from.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return preFrom, a.to.Set(space, preFrom)
}

func NewAssignment(from concept.Index, to concept.Index) *Assignment {
	back := &Assignment{
		from: from,
		to:   to,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
