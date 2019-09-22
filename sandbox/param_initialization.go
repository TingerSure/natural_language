package sandbox

type ParamInitialization struct {
	param string
}

func (a *ParamInitialization) Exec(space *Closure) (bool, error) {
	space.InitLocal(a.param)
	return true, nil
}

func NewParamInitialization(param string) *ParamInitialization {
	return &ParamInitialization{
		param: param,
	}
}
