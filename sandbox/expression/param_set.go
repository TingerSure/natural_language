package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type ParamSet struct {
	*adaptor.ExpressionIndex
	key   string
	value concept.Index
	param concept.Index
}

func (a *ParamSet) ToString(prefix string) string {
	return fmt.Sprintf("%v[%v] = %v", a.param.ToString(prefix), a.key, a.value.ToString(prefix))
}

func (a *ParamSet) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return nil, interrupt.NewException("type error", "Only Param can be set in ParamSet")
	}

	preValue, suspend := a.value.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	param.Set(a.key, preValue)
	return preValue, nil
}

func NewParamSet(param concept.Index, key string, value concept.Index) *ParamSet {
	back := &ParamSet{
		key:   key,
		value: value,
		param: param,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

func NewParamSetWithoutKey(param concept.Index, value concept.Index) *ParamSet {
	return NewParamSet(param, variable.ParamDefaultKey, value)
}
