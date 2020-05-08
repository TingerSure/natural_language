package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type NotEqualTo struct {
	*template.BinaryOperatorNumber
}

var (
	NotEqualToLanguageSeeds = map[string]func(string, *NotEqualTo) string{}
)

func (f *NotEqualTo) ToLanguage(language string) string {
	seed := NotEqualToLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewNotEqualTo(left concept.Index, right concept.Index) *NotEqualTo {
	return &NotEqualTo{
		template.NewBinaryOperatorNumber("!=", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() != right.Value()), nil
		}),
	}
}
