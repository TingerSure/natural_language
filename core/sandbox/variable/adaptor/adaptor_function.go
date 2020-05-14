package adaptor

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type AdaptorFunction struct {
	languageOnCallSeeds map[string]func(concept.Function, *concept.Mapping) string
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
			if target.EqualLanguage(src) {
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
		if name.EqualLanguage(back) {
			return name
		}
	}
	return back
}

func NewAdaptorFunction() *AdaptorFunction {
	return &AdaptorFunction{
		languageOnCallSeeds: map[string]func(concept.Function, *concept.Mapping) string{},
	}
}