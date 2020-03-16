package template

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type BinaryOperatorBool struct {
	*adaptor.ExpressionIndex
	left  concept.Index
	right concept.Index
	sign  string
	exec  func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Interrupt)
}

func (a *BinaryOperatorBool) ToString(prefix string) string {
	return fmt.Sprintf("%v %v %v", a.left.ToString(prefix), a.sign, a.right.ToString(prefix))
}

func (a *BinaryOperatorBool) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preLeft, suspend := a.left.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	preRight, suspend := a.right.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	left, yesLeft := variable.VariableFamilyInstance.IsBool(preLeft)
	right, yesRight := variable.VariableFamilyInstance.IsBool(preRight)
	if !yesLeft || !yesRight {
		return nil, interrupt.NewException(variable.NewString("type error"), variable.NewString("Only numbers can be added."))
	}
	return a.exec(left, right)
}

func NewBinaryOperatorBool(sign string, left concept.Index, right concept.Index, exec func(left *variable.Bool, right *variable.Bool) (concept.Variable, concept.Interrupt)) *BinaryOperatorBool {
	back := &BinaryOperatorBool{
		sign:  sign,
		left:  left,
		right: right,
		exec:  exec,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
