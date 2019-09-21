package sandbox

import (
	"errors"
)

type Subtraction struct {
	left   Index
	right  Index
	result Index
}

func (a *Subtraction) Exec(space *Closure) error {
	preLeft, errLeft := a.left.Get(space)
	preRight, errRight := a.right.Get(space)

	if errLeft != nil {
		return errLeft
	}
	if errRight != nil {
		return errRight
	}

	left, yesLeft := VariableFamilyInstance.IsNumber(preLeft)
	right, yesRight := VariableFamilyInstance.IsNumber(preRight)
	if !yesLeft || !yesRight {
		return errors.New("Only numbers can be subtracted.")
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
