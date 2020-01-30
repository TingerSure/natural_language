package expression

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/template"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type Subtraction struct {
	*template.BinaryOperatorNumber
}

func NewSubtraction(left concept.Index, right concept.Index) *Subtraction {
	return &Subtraction{
		template.NewBinaryOperatorNumber("-", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewNumber(left.Value() - right.Value()), nil
		}),
	}
}