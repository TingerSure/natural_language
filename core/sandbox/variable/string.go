package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableStringType = "string"
)

type StringSeed interface {
	ToLanguage(string, *String) string
	Type() string
	New(string) *String
}

type String struct {
	value   string
	mapping map[string]string
	seed    StringSeed
}

func (f *String) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (n *String) GetSystem() string {
	return n.value
}

func (n *String) SetSystem(value string) {
	n.value = value
}

func (n *String) GetLanguage(language string) string {
	value := n.mapping[language]
	if value == "" {
		return n.value
	}
	return value
}

func (n *String) SetLanguage(language string, value string) {
	n.mapping[language] = value
}

func (n *String) HasLanguage(language string) bool {
	return n.mapping[language] != ""
}

func (n *String) IsLanguage(language string, value string) bool {
	return n.mapping[language] == value
}

func (n *String) EqualLanguage(other concept.String) bool {
	if n.Equal(other) {
		return true
	}
	hit := false
	return !n.IterateLanguages(func(language string, value string) bool {
		return other.IterateLanguages(func(otherLanguage string, otherValue string) bool {
			if language == otherLanguage {
				if value != otherValue {
					return true
				}
				hit = true
			}
			return false
		})
	}) && hit
}

func (n *String) Equal(other concept.String) bool {
	return n.value == other.Value()
}

func (n *String) IterateLanguages(on func(string, string) bool) bool {
	for language, value := range n.mapping {
		if on(language, value) {
			return true
		}
	}
	return false
}

func (n *String) ToString(prefix string) string {
	return fmt.Sprintf("\"%v\"", n.value)
}

func (n *String) Value() string {
	return n.value
}

func (s *String) Type() string {
	return s.seed.Type()
}

func (s *String) Clone() concept.String {
	instance := s.seed.New(s.value)
	for language, value := range s.mapping {
		instance.mapping[language] = value
	}
	return instance
}

func (s *String) CloneTo(instance concept.String) {
	for language, value := range s.mapping {
		instance.SetLanguage(language, value)
	}
}

type StringCreator struct {
	Seeds map[string]func(string, *String) string
}

func (s *StringCreator) New(value string) *String {
	return &String{
		value:   value,
		mapping: make(map[string]string),
		seed:    s,
	}
}

func (s *StringCreator) ToLanguage(language string, instance *String) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *StringCreator) Type() string {
	return VariableStringType
}

func NewStringCreator() *StringCreator {
	return &StringCreator{
		Seeds: map[string]func(string, *String) string{},
	}
}

func StringJoin(values []concept.String, separator string) string {
	paramsToString := make([]string, 0, len(values))
	for _, value := range values {
		paramsToString = append(paramsToString, value.ToString(""))
	}
	return strings.Join(paramsToString, separator)
}
