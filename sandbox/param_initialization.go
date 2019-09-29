package sandbox

import (
	"fmt"
)

type ParamInitialization struct {
	param string
}

func (a *ParamInitialization) ToString(prefix string) string {
	return fmt.Sprintf("%vvar %v", prefix, a.param)
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
