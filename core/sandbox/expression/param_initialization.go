package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type ParamInitialization struct {
	*adaptor.ExpressionIndex
	param       concept.String
	defaltValue concept.Index
}

func (a *ParamInitialization) ToString(prefix string) string {
	return fmt.Sprintf("var %v = %v", a.param.ToString(prefix), a.defaltValue.ToString(prefix))
}

func (a *ParamInitialization) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	value, suspend := a.defaltValue.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	space.InitLocal(a.param, value)
	return value, nil

}

func NewParamInitialization(param concept.String, defaltValue concept.Index) *ParamInitialization {
	back := &ParamInitialization{
		param:       param,
		defaltValue: defaltValue,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
