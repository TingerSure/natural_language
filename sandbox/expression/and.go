package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type And struct {
	*template.BinaryOperatorBool
}

func NewAnd(left concept.Index, right concept.Index) *And {
	return &And{
		template.NewBinaryOperatorBool("&&", left, right, func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() && right.Value()), nil
		}),
	}
}
