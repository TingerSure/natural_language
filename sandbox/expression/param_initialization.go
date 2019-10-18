package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
)

type ParamInitialization struct {
	*adaptor.ExpressionIndex
	param       string
	defaltValue concept.Index
}

func (a *ParamInitialization) ToString(prefix string) string {
	return fmt.Sprintf("var %v = %v", a.param, a.defaltValue.ToString(prefix))
}

func (a *ParamInitialization) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	space.InitLocal(a.param)
	value, suspend := a.defaltValue.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return value, space.SetLocal(a.param, value)

}

func NewParamInitialization(param string, defaltValue concept.Index) *ParamInitialization {
	back := &ParamInitialization{
		param:       param,
		defaltValue: defaltValue,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
