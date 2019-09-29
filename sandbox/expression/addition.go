package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type Addition struct {
	*template.BinaryOperatorNumber
}

func NewAddition(left concept.Index, right concept.Index, result concept.Index) *Addition {
	return &Addition{
		template.NewBinaryOperatorNumber("+", left, right, result, func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt) {
			return variable.NewNumber(left.Value() + right.Value()), nil
		}),
	}
}
