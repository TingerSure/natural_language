package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type Addition struct {
	*template.BinaryOperatorNumber
}

var (
	AdditionLanguageSeeds = map[string]func(string, *Addition) string{}
)

func (f *Addition) ToLanguage(language string) string {
	seed := AdditionLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewAddition(left concept.Index, right concept.Index) *Addition {
	return &Addition{
		template.NewBinaryOperatorNumber("+", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewNumber(left.Value() + right.Value()), nil
		}),
	}
}
