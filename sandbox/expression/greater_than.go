package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type GreaterThan struct {
	*template.BinaryOperatorNumber
}

func NewGreaterThan(left concept.Index, right concept.Index) *GreaterThan {
	return &GreaterThan{
		template.NewBinaryOperatorNumber(">", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() > right.Value()), nil
		}),
	}
}
