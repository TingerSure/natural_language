package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableStringType = "string"
)

type String struct {
	value   string
	mapping map[string]string
}

var (
	StringLanguageSeeds = map[string]func(string, *String) string{}
)

func (f *String) ToLanguage(language string) string {
	seed := StringLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
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
	return VariableStringType
}

func (s *String) Clone() concept.String {
	instance := NewString(s.value)
	for language, value := range s.mapping {
		instance.mapping[language] = value
	}
	return instance
}

func NewString(value string) *String {
	return &String{
		value:   value,
		mapping: make(map[string]string),
	}
}

type StringSeed struct {
	Seeds map[string]func(string, *String) string
}

func (s *StringSeed) New(value string) *String {
	return &String{
		value:   value,
		mapping: make(map[string]string),
	}
}

func NewStringSeed() *StringSeed {
	return &StringSeed{
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
