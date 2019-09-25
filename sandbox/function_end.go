package sandbox

type FunctionEnd struct {
}

func (a *FunctionEnd) Exec(space *Closure) Interrupt {
	return NewEnd()
}

func NewFunctionEnd() *FunctionEnd {
	return &FunctionEnd{}
}
