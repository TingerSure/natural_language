package sandbox

import (
	"fmt"
)

const (
	VariableNumberType = "number"
)

type Number struct {
	value float64
}

func (a *Number) ToString(prefix string) string {
	return fmt.Sprintf("%v", a.value)
}

func (n *Number) Value() float64 {
	return n.value
}

func (n *Number) Type() string {
	return VariableNumberType
}

func NewNumber(value float64) *Number {
	return &Number{
		value: value,
	}
}
