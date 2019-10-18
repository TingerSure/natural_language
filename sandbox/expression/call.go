package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

var (
	callDefaultParam = variable.NewParam()
)

type Call struct {
	*adaptor.ExpressionIndex
	funcs concept.Index
	param concept.Index
}

func (a *Call) ToString(prefix string) string {
	return fmt.Sprintf("%v(%v)", a.funcs.ToString(prefix), a.param.ToString(prefix))
}

func (a *Call) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preFuncs, suspend := a.funcs.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	funcs, yesFuncs := variable.VariableFamilyInstance.IsFunction(preFuncs)
	if !yesFuncs {
		return nil, interrupt.NewException("type error", "Only Function can be Called.")
	}

	param := callDefaultParam
	if !nl_interface.IsNil(a.param) {
		preParam, suspend := a.param.Get(space)
		if !nl_interface.IsNil(suspend) {
			return nil, suspend
		}
		yesParam := false
		param, yesParam = variable.VariableFamilyInstance.IsParam(preParam)
		if !yesParam {
			return nil, interrupt.NewException("type error", "Only Param can are passed to a Function")
		}
	}

	return funcs.Exec(param)
}

func NewCall(funcs concept.Index, param concept.Index) *Call {
	back := &Call{
		funcs: funcs,
		param: param,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
