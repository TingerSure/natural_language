package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ParamGet struct {
	*adaptor.ExpressionIndex
	key   concept.String
	param concept.Index
}

var (
	ParamGetLanguageSeeds = map[string]func(string, *ParamGet) string{}
)

func (f *ParamGet) ToLanguage(language string) string {
	seed := ParamGetLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (a *ParamGet) ToString(prefix string) string {
	return fmt.Sprintf("%v[%v]", a.param.ToString(prefix), a.key.ToString(prefix))
}

func (a *ParamGet) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return nil, interrupt.NewException(variable.NewString("type error"), variable.NewString("Only Param can be get in ParamGet"))
	}

	return param.Get(a.key), nil
}

func NewParamGet(param concept.Index, key concept.String) *ParamGet {
	back := &ParamGet{
		key:   key,
		param: param,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

func NewParamGetWithoutKey(param concept.Index) *ParamGet {
	return NewParamGet(param, variable.NewString(variable.ParamDefaultKey))
}
