package sandbox

import (
	"fmt"
)

type Assignment struct {
	from Index
	to   Index
}

func (a *Assignment) ToString(prefix string) string {
	return fmt.Sprintf("%v%v = %v", prefix, a.to.ToString(prefix), a.from.ToString(prefix))
}

func (a *Assignment) Exec(space *Closure) Interrupt {
	preFrom, exception := a.from.Get(space)
	if exception != nil {
		return exception
	}
	return a.to.Set(space, preFrom)
}

func NewAssignment(from Index, to Index) *Assignment {
	return &Assignment{
		from: from,
		to:   to,
	}
}
