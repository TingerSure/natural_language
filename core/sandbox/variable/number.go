package variable

import (
	"fmt"
)

const (
	VariableNumberType = "number"
)

type Number struct {
	value float64
}

var (
	NumberLanguageSeeds = map[string]func(string, *Number) string{}
)

func (f *Number) ToLanguage(language string) string {
	seed := NumberLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
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
