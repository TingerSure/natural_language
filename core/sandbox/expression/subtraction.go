package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type Subtraction struct {
	*template.BinaryOperatorNumber
}

var (
	SubtractionLanguageSeeds = map[string]func(string, *Subtraction) string{}
)

func (f *Subtraction) ToLanguage(language string) string {
	seed := SubtractionLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewSubtraction(left concept.Index, right concept.Index) *Subtraction {
	return &Subtraction{
		template.NewBinaryOperatorNumber("-", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewNumber(left.Value() - right.Value()), nil
		}),
	}
}
