package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type Division struct {
	*template.BinaryOperatorNumber
}

func NewDivision(left concept.Index, right concept.Index, result concept.Index) *Division {
	return &Division{
		template.NewBinaryOperatorNumber("/", left, right, result, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			if right.Value() == 0 {
				return variable.NewNumber(0), interrupt.NewException("param error", "Division right cannot be 0")
			}
			return variable.NewNumber(left.Value() / right.Value()), nil
		}),
	}
}
