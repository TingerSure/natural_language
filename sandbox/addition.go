package sandbox

import (
	"fmt"
)

type Addition struct {
	left   Index
	right  Index
	result Index
}

func (a *Addition) ToString(prefix string) string {
	return fmt.Sprintf("%v%v = %v + %v", prefix, a.result.ToString(prefix), a.left.ToString(prefix), a.right.ToString(prefix))
}

func (a *Addition) Exec(space *Closure) Interrupt {
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
	return a.result.Set(space, NewNumber(left.Value()+right.Value()))
}

func NewAddition(left Index, right Index, result Index) *Addition {
	return &Addition{
		left:   left,
		right:  right,
		result: result,
	}
}
