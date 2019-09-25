package sandbox

type ParamInitialization struct {
	param string
}

func (a *ParamInitialization) Exec(space *Closure) Interrupt {
	space.InitLocal(a.param)
	return nil
}

func NewParamInitialization(param string) *ParamInitialization {
	return &ParamInitialization{
		param: param,
	}
}
