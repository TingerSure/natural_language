package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

const (
	selfIndexKey = "self"
)

type SelfIndex struct {
}

var (
	SelfIndexLanguageSeeds = map[string]func(string, *SelfIndex) string{}
)

func (f *SelfIndex) ToLanguage(language string) string {
	seed := SelfIndexLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (s *SelfIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *SelfIndex) ToString(prefix string) string {
	return selfIndexKey
}

func (s *SelfIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(variable.NewString(selfIndexKey))
}

func (s *SelfIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException(variable.NewString("read only"), variable.NewString("Self cannot be changed."))

}

func NewSelfIndex() *SelfIndex {
	return &SelfIndex{}
}
