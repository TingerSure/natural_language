package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type EqualTo struct {
	*template.BinaryOperatorNumber
}

var (
	EqualToLanguageSeeds = map[string]func(string, *EqualTo) string{}
)

func (f *EqualTo) ToLanguage(language string) string {
	seed := EqualToLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewEqualTo(left concept.Index, right concept.Index) *EqualTo {
	return &EqualTo{
		template.NewBinaryOperatorNumber("==", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() == right.Value()), nil
		}),
	}
}
