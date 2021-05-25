package adaptor

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type AdaptorFunctionParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type AdaptorFunction struct {
	*AdaptorVariable
	languageOnCallSeeds map[string]func(concept.Function, *concept.Mapping) string
}

func (o *AdaptorFunction) IsFunction() bool {
	return true
}

func (a *AdaptorFunction) GetLanguageOnCallSeed(language string) func(concept.Function, *concept.Mapping) string {
	return a.languageOnCallSeeds[language]
}

func (a *AdaptorFunction) SetLanguageOnCallSeed(language string, seed func(concept.Function, *concept.Mapping) string) {
	a.languageOnCallSeeds[language] = seed
}

func (*AdaptorFunction) AdaptorParamFormat(f concept.Function, params *concept.Mapping) *concept.Mapping {
	keys := f.ParamNames()
	instance := concept.NewMapping(params.Param())
	params.Iterate(func(target concept.String, value interface{}) bool {
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
		languageOnCallSeeds: map[string]func(concept.Function, *concept.Mapping) string{},
	}
}
