package sandbox

type Assignment struct {
	from Index
	to   Index
}

func (a *Assignment) Exec(space *Closure) error {
	preFrom, errFrom := a.from.Get(space)
	if errFrom != nil {
		return errFrom
	}
	return a.to.Set(space, preFrom)
}

func NewAssignment(from Index, to Index) *Assignment {
	return &Assignment{
		from: from,
		to:   to,
	}
}
