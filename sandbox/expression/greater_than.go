package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type GreaterThan struct {
	left   concept.Index
	right  concept.Index
	result concept.Index
}

func (a *GreaterThan) ToString(prefix string) string {
	return fmt.Sprintf("%v%v = %v > %v", prefix, a.result.ToString(prefix), a.left.ToString(prefix), a.right.ToString(prefix))
}

func (a *GreaterThan) Exec(space concept.Closure) concept.Interrupt {
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
		return interrupt.NewException("type error", "Illegal type.")
	}
	return a.result.Set(space, variable.NewBool(left.Value() > right.Value()))
}

func NewGreaterThan(left concept.Index, right concept.Index, result concept.Index) *GreaterThan {
	return &GreaterThan{
		left:   left,
		right:  right,
		result: result,
	}
}
