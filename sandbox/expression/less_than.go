package expression

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
    "github.com/TingerSure/natural_language/sandbox/variable"
	"github.com/TingerSure/natural_language/sandbox/expression/template"
)

type LessThan struct {
    *template.BinaryOperatorNumber
}

func NewLessThan(left concept.Index, right concept.Index, result concept.Index) *LessThan {
	return &LessThan{
		template.NewBinaryOperatorNumber("<",left,right,result,func(left *variable.Number, right *variable.Number)concept.Variable{
            return variable.NewBool(left.Value() < right.Value())
        }),
	}
}
