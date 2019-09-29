package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type ParamInitialization struct {
	param string
}

func (a *ParamInitialization) ToString(prefix string) string {
	return fmt.Sprintf("%vvar %v", prefix, a.param)
}

func (a *ParamInitialization) Exec(space concept.Closure) concept.Interrupt {
	space.InitLocal(a.param)
	return nil
}

func NewParamInitialization(param string) *ParamInitialization {
	return &ParamInitialization{
		param: param,
	}
}
