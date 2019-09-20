package sandbox

import (
	"errors"
)

type Addition struct {
	left   Index
	right  Index
	result Index
}

func (a *Addition) Exec(space Closure) error {
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
		return errors.New("Only numbers can be added.")
	}
	return a.result.Set(space, NewNumber(left.Value()+right.Value()))
}

func NewAddition(left Index, right Index, result Index) *Addition {
	return &Addition{
		left:   left,
		right:  right,
		result: result,
	}
}
