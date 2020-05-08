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

var (
	BoolLanguageSeeds = map[string]func(string, *Bool) string{}
)

func (f *Bool) ToLanguage(language string) string {
	seed := BoolLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
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
