package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ConstIndex struct {
	value concept.Variable
}

const (
	IndexConstType = "Const"
)

func (f *ConstIndex) Type() string {
	return IndexConstType
}

var (
	ConstIndexLanguageSeeds = map[string]func(string, *ConstIndex) string{}
)

func (f *ConstIndex) ToLanguage(language string) string {
	seed := ConstIndexLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (s *ConstIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ConstIndex) ToString(prefix string) string {
	return s.value.ToString(prefix)
}

func (s *ConstIndex) Value() concept.Variable {
	return s.value
}

func (s *ConstIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.value, nil
}

func (s *ConstIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException(variable.NewString("read only"), variable.NewString("Constants cannot be changed."))
}

func NewConstIndex(value concept.Variable) *ConstIndex {
	return &ConstIndex{
		value: value,
	}
}
