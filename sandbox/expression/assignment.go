package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type Assignment struct {
	from concept.Index
	to   concept.Index
}

func (a *Assignment) ToString(prefix string) string {
	return fmt.Sprintf("%v%v = %v", prefix, a.to.ToString(prefix), a.from.ToString(prefix))
}

func (a *Assignment) Exec(space concept.Closure) concept.Interrupt {
	preFrom, suspend := a.from.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	return a.to.Set(space, preFrom)
}

func NewAssignment(from concept.Index, to concept.Index) *Assignment {
	return &Assignment{
		from: from,
		to:   to,
	}
}
