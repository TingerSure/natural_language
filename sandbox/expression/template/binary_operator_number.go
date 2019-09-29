package template

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type BinaryOperatorNumber struct {
	left   concept.Index
	right  concept.Index
	result concept.Index
    sign   string
    exec   func (left *variable.Number, right *variable.Number) concept.Variable
}

func (a *BinaryOperatorNumber) ToString(prefix string) string {
	return fmt.Sprintf("%v%v = %v %v %v", prefix, a.result.ToString(prefix), a.left.ToString(prefix), a.sign , a.right.ToString(prefix))
}

func (a *BinaryOperatorNumber) Exec(space concept.Closure) concept.Interrupt {
	preLeft, suspend := a.left.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	preRight, suspend := a.right.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}

	left, yesLeft := variable.VariableFamilyInstance.IsNumber(preLeft)
	right, yesRight := variable.VariableFamilyInstance.IsNumber(preRight)
	if !yesLeft || !yesRight {
		return interrupt.NewException("type error", "Only numbers can be added.")
	}
	return a.result.Set(space, a.exec(left,right))
}

func NewBinaryOperatorNumber(sign string, left concept.Index, right concept.Index, result concept.Index, exec func (left *variable.Number, right *variable.Number) concept.Variable) *BinaryOperatorNumber {
	return &BinaryOperatorNumber{
        sign:   sign,
		left:   left,
		right:  right,
		result: result,
        exec : exec,
	}
}
