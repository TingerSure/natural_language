package sandbox

import (
	"errors"
)

type Subtraction struct {
	left   Index
	right  Index
	result Index
}

func (a *Subtraction) Exec(space *Closure) (bool, error) {
	preLeft, errLeft := a.left.Get(space)
	preRight, errRight := a.right.Get(space)

	if errLeft != nil {
		return false, errLeft
	}
	if errRight != nil {
		return false, errRight
	}

	left, yesLeft := VariableFamilyInstance.IsNumber(preLeft)
	right, yesRight := VariableFamilyInstance.IsNumber(preRight)
	if !yesLeft || !yesRight {
		return false, errors.New("Only numbers can be subtracted.")
	}
	return true, a.result.Set(space, NewNumber(left.Value()-right.Value()))
}

func NewSubtraction(left Index, right Index, result Index) *Subtraction {
	return &Subtraction{
		left:   left,
		right:  right,
		result: result,
	}
}
