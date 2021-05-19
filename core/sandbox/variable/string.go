package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
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
	*adaptor.AdaptorVariable
	value   string
	mapping map[string]string
	seed    StringSeed
}

func (o *String) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
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

	if other.IterateLanguages(func(otherLanguage string, otherValue string) bool {
		if n.value == otherValue {
			return true
		}
		return false
	}) {
		return true
	}

	if n.IterateLanguages(func(language string, value string) bool {
		if other.Value() == value {
			return true
		}
		return false
	}) {
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

type StringCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type StringCreator struct {
	Seeds map[string]func(string, *String) string
	param *StringCreatorParam
}

func (s *StringCreator) New(value string) *String {
	return &String{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
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

func NewStringCreator(param *StringCreatorParam) *StringCreator {
	return &StringCreator{
		Seeds: map[string]func(string, *String) string{},
		param: param,
	}
}

func StringJoin(values []concept.String, separator string) string {
	paramsToString := make([]string, 0, len(values))
	for _, value := range values {
		paramsToString = append(paramsToString, value.Value())
	}
	return strings.Join(paramsToString, separator)
}
