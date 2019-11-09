package template

import (
	"fmt"
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type BinaryOperatorNumber struct {
	*adaptor.ExpressionIndex
	left  concept.Index
	right concept.Index
	sign  string
	exec  func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt)
}

func (a *BinaryOperatorNumber) ToString(prefix string) string {
	return fmt.Sprintf("%v %v %v", a.left.ToString(prefix), a.sign, a.right.ToString(prefix))
}

func (a *BinaryOperatorNumber) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preLeft, suspend := a.left.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	preRight, suspend := a.right.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	left, yesLeft := variable.VariableFamilyInstance.IsNumber(preLeft)
	right, yesRight := variable.VariableFamilyInstance.IsNumber(preRight)
	if !yesLeft || !yesRight {
		return nil, interrupt.NewException("type error", "Only numbers can be added.")
	}
	return a.exec(left, right)
}

func NewBinaryOperatorNumber(sign string, left concept.Index, right concept.Index, exec func(left *variable.Number, right *variable.Number) (concept.Variable, concept.Interrupt)) *BinaryOperatorNumber {
	back := &BinaryOperatorNumber{
		sign:  sign,
		left:  left,
		right: right,
		exec:  exec,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
