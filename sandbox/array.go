package sandbox

import (
	"errors"
	"fmt"
	"strings"
)

const (
	VariableArrayType = "array"
)

type Array struct {
	values []Variable
	length int
}

func (a *Array) ToString(prefix string) string {
	itemPrefix := fmt.Sprintf("%v\t", prefix)
	valuesToStrings := make([]string, len(a.values))
	for _, value := range a.values {
		valuesToStrings = append(valuesToStrings, value.ToString(itemPrefix))
	}
	return fmt.Sprintf("[%v]", strings.Join(valuesToStrings, ", "))
}

func (a *Array) Set(index int, value Variable) error {
	if index < 0 || index >= a.length {
		return errors.New("array index out of bounds error.")
	}
	a.values[index] = value
	return nil
}

func (a *Array) Get(index int) (Variable, error) {
	if index < 0 || index >= a.length {
		return nil, errors.New("array index out of bounds error.")
	}
	return a.values[index], nil
}

func (a *Array) Length() int {
	return a.length
}

func (a *Array) Type() string {
	return VariableArrayType
}

func NewArray(size int) *Array {
	return &Array{
		values: make([]Variable, size),
		length: size,
	}
}
