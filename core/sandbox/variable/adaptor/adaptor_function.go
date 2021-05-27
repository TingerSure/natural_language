package adaptor

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type AdaptorFunctionParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
}

type AdaptorFunction struct {
	*AdaptorVariable
	param               *AdaptorFunctionParam
	languageOnCallSeeds map[string]func(concept.Function, concept.Param) string
}

func (o *AdaptorFunction) IsFunction() bool {
	return true
}

func (a *AdaptorFunction) GetLanguageOnCallSeed(language string) func(concept.Function, concept.Param) string {
	return a.languageOnCallSeeds[language]
}

func (a *AdaptorFunction) SetLanguageOnCallSeed(language string, seed func(concept.Function, concept.Param) string) {
	a.languageOnCallSeeds[language] = seed
}

func (a *AdaptorFunction) AdaptorParamFormat(f concept.Function, params concept.Param) concept.Param {
	keys := f.ParamNames()
	instance := a.param.ParamCreator()
	params.Iterate(func(target concept.String, value concept.Variable) bool {
		for _, src := range keys {
			if target.Equal(src) {
				instance.Set(src, value)
				return false
			}
		}
		instance.Set(target, value)
		return false
	})
	return instance
}

func (*AdaptorFunction) AdaptorReturnFormat(f concept.Function, back concept.String) concept.String {
	for _, name := range f.ReturnNames() {
		if name.Equal(back) {
			return name
		}
	}
	return back
}

func NewAdaptorFunction(param *AdaptorFunctionParam) *AdaptorFunction {
	return &AdaptorFunction{
		AdaptorVariable: NewAdaptorVariable(&AdaptorVariableParam{
			NullCreator:      param.NullCreator,
			ExceptionCreator: param.ExceptionCreator,
		}),
		param:               param,
		languageOnCallSeeds: map[string]func(concept.Function, concept.Param) string{},
	}
}
