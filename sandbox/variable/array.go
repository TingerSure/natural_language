package variable

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"strings"
)

const (
	VariableArrayType = "array"
)

type Array struct {
	values []concept.Variable
	length int
}

func (a *Array) ToString(prefix string) string {
	itemPrefix := fmt.Sprintf("%v\t", prefix)
	valuesToStrings := make([]string, 0, len(a.values))
	for _, value := range a.values {
		valuesToStrings = append(valuesToStrings, value.ToString(itemPrefix))
	}
	return fmt.Sprintf("[%v]", strings.Join(valuesToStrings, ", "))
}

func (a *Array) Set(index int, value concept.Variable) error {
	if index < 0 || index >= a.length {
		return errors.New("array index out of bounds error.")
	}
	a.values[index] = value
	return nil
}

func (a *Array) Get(index int) (concept.Variable, error) {
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
		values: make([]concept.Variable, size),
		length: size,
	}
}
