package variable

import (
	"fmt"
)

const (
	VariableBoolType = "bool"
)

type BoolSeed interface {
	ToLanguage(string, *Bool) string
	Type() string
}

type Bool struct {
	value bool
	seed  BoolSeed
}

func (f *Bool) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Bool) ToString(prefix string) string {
	return fmt.Sprintf("%v", a.value)
}

func (n *Bool) Value() bool {
	return n.value
}

func (n *Bool) Type() string {
	return n.seed.Type()
}

type BoolCreator struct {
	Seeds map[string]func(string, *Bool) string
}

func (s *BoolCreator) New(value bool) *Bool {
	return &Bool{
		value: value,
		seed:  s,
	}
}

func (s *BoolCreator) ToLanguage(language string, instance *Bool) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *BoolCreator) Type() string {
	return VariableBoolType
}

func NewBoolCreator() *BoolCreator {
	return &BoolCreator{
		Seeds: map[string]func(string, *Bool) string{},
	}
}
