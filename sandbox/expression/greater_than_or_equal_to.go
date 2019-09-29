package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type GreaterThanOrEqualTo struct {
	*template.BinaryOperatorNumber
}

func NewGreaterThanOrEqualTo(left concept.Index, right concept.Index, result concept.Index) *GreaterThanOrEqualTo {
	return &GreaterThanOrEqualTo{
		template.NewBinaryOperatorNumber(">=", left, right, result, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() >= right.Value()), nil
		}),
	}
}