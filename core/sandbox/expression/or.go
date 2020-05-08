package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type Or struct {
	*template.BinaryOperatorBool
}

var (
	OrLanguageSeeds = map[string]func(string, *Or) string{}
)

func (f *Or) ToLanguage(language string) string {
	seed := OrLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewOr(left concept.Index, right concept.Index) *Or {
	return &Or{
		template.NewBinaryOperatorBool("||", left, right, func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() || right.Value()), nil
		}),
	}
}
