package variable

import (
	"fmt"
)

const (
	VariableNumberType = "number"
)

type NumberSeed interface {
	ToLanguage(string, *Number) string
	Type() string
}

type Number struct {
	value float64
	seed  NumberSeed
}

func (f *Number) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Number) ToString(prefix string) string {
	return fmt.Sprintf("%v", a.value)
}

func (n *Number) Value() float64 {
	return n.value
}

func (n *Number) Type() string {
	return n.seed.Type()
}

type NumberCreator struct {
	Seeds map[string]func(string, *Number) string
}

func (s *NumberCreator) New(value float64) *Number {
	return &Number{
		value: value,
		seed:  s,
	}
}

func (s *NumberCreator) ToLanguage(language string, instance *Number) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *NumberCreator) Type() string {
	return VariableNumberType
}

func NewNumberCreator() *NumberCreator {
	return &NumberCreator{
		Seeds: map[string]func(string, *Number) string{},
	}
}
