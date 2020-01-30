package variable

import (
	"fmt"
)

const (
	VariableStringType = "string"
)

type String struct {
	value string
}

func (n *String) ToString(prefix string) string {
	return fmt.Sprintf("\"%v\"", n.value)
}
func (n *String) Value() string {
	return n.value
}

func (s *String) Type() string {
	return VariableStringType
}

func NewString(value string) *String {
	return &String{
		value: value,
	}
}
