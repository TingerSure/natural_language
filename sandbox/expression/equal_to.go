package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type EqualTo struct {
	*template.BinaryOperatorNumber
}

func NewEqualTo(left concept.Index, right concept.Index) *EqualTo {
	return &EqualTo{
		template.NewBinaryOperatorNumber("==", left, right, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewBool(left.Value() == right.Value()), nil
		}),
	}
}
