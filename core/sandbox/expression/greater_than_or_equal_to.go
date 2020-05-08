package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type GreaterThanOrEqualTo struct {
	*template.BinaryOperatorNumber
}

var (
	GreaterThanOrEqualToLanguageSeeds = map[string]func(string, *GreaterThanOrEqualTo) string{}
)

func (f *GreaterThanOrEqualTo) ToLanguage(language string) string {
	seed := GreaterThanOrEqualToLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewGreaterThanOrEqualTo(left concept.Index, right concept.Index) *GreaterThanOrEqualTo {
	return &GreaterThanOrEqualTo{
		template.NewBinaryOperatorNumber(">=", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() >= right.Value()), nil
		}),
	}
}
