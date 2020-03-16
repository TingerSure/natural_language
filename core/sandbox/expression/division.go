package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type Division struct {
	*template.BinaryOperatorNumber
}

func NewDivision(left concept.Index, right concept.Index) *Division {
	return &Division{
		template.NewBinaryOperatorNumber("/", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			if right.Value() == 0 {
				return variable.NewNumber(0), interrupt.NewException(variable.NewString("param error"), variable.NewString("Division right cannot be 0"))
			}
			return variable.NewNumber(left.Value() / right.Value()), nil
		}),
	}
}
