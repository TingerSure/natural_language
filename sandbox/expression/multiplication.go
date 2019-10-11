package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type Multiplication struct {
	*template.BinaryOperatorNumber
}

func NewMultiplication(left concept.Index, right concept.Index) *Multiplication {
	return &Multiplication{
		template.NewBinaryOperatorNumber("*", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewNumber(left.Value() * right.Value()), nil
		}),
	}
}
