package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type Multiplication struct {
	*template.BinaryOperatorNumber
}

var (
	MultiplicationLanguageSeeds = map[string]func(string, *Multiplication) string{}
)

func (f *Multiplication) ToLanguage(language string) string {
	seed := MultiplicationLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewMultiplication(left concept.Index, right concept.Index) *Multiplication {
	return &Multiplication{
		template.NewBinaryOperatorNumber("*", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewNumber(left.Value() * right.Value()), nil
		}),
	}
}
