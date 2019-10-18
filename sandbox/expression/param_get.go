package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type ParamGet struct {
	*adaptor.ExpressionIndex
	key   string
	param concept.Index
}

func (a *ParamGet) ToString(prefix string) string {
	return fmt.Sprintf("%v[%v]", a.param.ToString(prefix), a.key)
}

func (a *ParamGet) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return nil, interrupt.NewException("type error", "Only Param can be get in ParamGet")
	}

	return param.Get(a.key), nil
}

func NewParamGet(param concept.Index, key string) *ParamGet {
	back := &ParamGet{
		key:   key,
		param: param,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

func NewParamGetWithoutKey(param concept.Index) *ParamGet {
	return NewParamGet(param, variable.ParamDefaultKey)
}
