package variable

import (
	"fmt"
)

const (
	VariableBoolType = "bool"
)

type Bool struct {
	value bool
}

func (a *Bool) ToString(prefix string) string {
	return fmt.Sprintf("%v", a.value)
}
func (n *Bool) Value() bool {
	return n.value
}

func (n *Bool) Type() string {
	return VariableBoolType
}

func NewBool(value bool) *Bool {
	return &Bool{
		value: value,
	}
}
