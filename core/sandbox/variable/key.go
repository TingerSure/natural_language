package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableKeyType = "key"
)

type Key struct {
	mapping map[string]string
}

func (k *Key) Is(specimen concept.KeySpecimen) bool {
	return k.mapping[specimen.GetLanguage()] == specimen.GetName()
}

func (k *Key) Iterate(on func(concept.KeySpecimen) bool) bool {
	for language, name := range k.mapping {
		if on(NewKeySpecimenWithInit(language, name)) {
			return true
		}
	}
	return false
}

func (k *Key) Equal(other concept.Key) bool {
	count := 0

	return !k.Iterate(func(specimen concept.KeySpecimen) bool {
		return other.Iterate(func(otherSpecimen concept.KeySpecimen) bool {
			if specimen.GetLanguage() == otherSpecimen.GetLanguage() {
				if specimen.GetName() == otherSpecimen.GetName() {
					count++
				} else {
					return true
				}
			}
			return false
		})
	}) && count != 0
}

func (k *Key) Get(specimen concept.KeySpecimen) {
	specimen.SetName(k.mapping[specimen.GetLanguage()])
}

func (k *Key) Set(specimen concept.KeySpecimen) {
	k.mapping[specimen.GetLanguage()] = specimen.GetName()
}

func (k *Key) Type() string {
	return VariableKeyType
}

func (k *Key) ToString(prefix string) string {
	paramsToString := make([]string, 0, len(k.mapping))
	for language, name := range k.mapping {
		paramsToString = append(paramsToString, fmt.Sprintf("%v.%v", language, name))
	}
	return strings.Join(paramsToString, ", ")
}

func NewKey() *Key {
	return &Key{
		mapping: make(map[string]string),
	}
}

func NewKeyWhitInit(specimen concept.KeySpecimen) *Key {
	key := &Key{
		mapping: make(map[string]string),
	}
	key.Set(specimen)
	return key
}
