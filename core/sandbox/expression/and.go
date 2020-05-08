package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type And struct {
	*template.BinaryOperatorBool
}

var (
	AndLanguageSeeds = map[string]func(string, *And) string{}
)

func (f *And) ToLanguage(language string) string {
	seed := AndLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewAnd(left concept.Index, right concept.Index) *And {
	return &And{
		template.NewBinaryOperatorBool("&&", left, right, func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() && right.Value()), nil
		}),
	}
}
