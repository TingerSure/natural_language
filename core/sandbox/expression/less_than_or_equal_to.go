package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type LessThanOrEqualTo struct {
	*template.BinaryOperatorNumber
}

var (
	LessThanOrEqualToLanguageSeeds = map[string]func(string, *LessThanOrEqualTo) string{}
)

func (f *LessThanOrEqualTo) ToLanguage(language string) string {
	seed := LessThanOrEqualToLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewLessThanOrEqualTo(left concept.Index, right concept.Index) *LessThanOrEqualTo {
	return &LessThanOrEqualTo{
		template.NewBinaryOperatorNumber("<=", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() <= right.Value()), nil
		}),
	}
}
