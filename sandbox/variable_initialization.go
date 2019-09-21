package sandbox

type VariableInitialization struct {
	param string
}

func (a *VariableInitialization) Exec(space *Closure) error {
	space.InitLocal(a.param)
	return nil
}

func NewVariableInitialization(param string) *VariableInitialization {
	return &VariableInitialization{
		param: param,
	}
}
