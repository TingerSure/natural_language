package sandbox

import (
	"fmt"
)

type Subtraction struct {
	left   Index
	right  Index
	result Index
}

func (a *Subtraction) ToString(prefix string) string {
	return fmt.Sprintf("%v%v = %v - %v", prefix, a.result.ToString(prefix), a.left.ToString(prefix), a.right.ToString(prefix))
}

func (a *Subtraction) Exec(space *Closure) Interrupt {
	preLeft, exception := a.left.Get(space)
	if exception != nil {
		return exception
	}
	preRight, exception := a.right.Get(space)
	if exception != nil {
		return exception
	}

	left, yesLeft := VariableFamilyInstance.IsNumber(preLeft)
	right, yesRight := VariableFamilyInstance.IsNumber(preRight)
	if !yesLeft || !yesRight {
		return NewException("type error", "Only numbers can be added.")
	}
	return a.result.Set(space, NewNumber(left.Value()-right.Value()))

}

func NewSubtraction(left Index, right Index, result Index) *Subtraction {
	return &Subtraction{
		left:   left,
		right:  right,
		result: result,
	}
}
