package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type Or struct {
	*template.BinaryOperatorBool
}

func NewOr(left concept.Index, right concept.Index) *Or {
	return &Or{
		template.NewBinaryOperatorBool("||", left, right, func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() || right.Value()), nil
		}),
	}
}
