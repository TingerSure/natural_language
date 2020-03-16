package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	callDefaultParam = index.NewConstIndex(variable.NewParam())
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
	funcs, yesFuncs := variable.VariableFamilyInstance.IsFunctionHome(preFuncs)
	if !yesFuncs {
		return nil, interrupt.NewException(variable.NewString("type error"), variable.NewString("Only Function can be Called."))
	}

	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	yesParam := false
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return nil, interrupt.NewException(variable.NewString("type error"), variable.NewString("Only Param can are passed to a Function"))
	}

	return funcs.Exec(param, nil)
}

func NewCall(funcs concept.Index, param concept.Index) *Call {
	if nl_interface.IsNil(param) {
		param = callDefaultParam
	}
	back := &Call{
		funcs: funcs,
		param: param,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
