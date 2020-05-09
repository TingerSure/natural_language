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

func NewAdaptorFunction() *AdaptorFunction {
	return &AdaptorFunction{
		languageOnCallSeeds: map[string]func(concept.Function, *concept.Mapping) string{},
	}
}
