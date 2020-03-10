package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	VariableKeySpecimenType = "key_specimen"
)

type KeySpecimen struct {
	language string
	name     string
}

func (k *KeySpecimen) Equal(other concept.KeySpecimen) bool {
	return k.language == other.GetLanguage() && k.name == other.GetName()
}

func (k *KeySpecimen) GetName() string {
	return k.name
}

func (k *KeySpecimen) GetLanguage() string {
	return k.language
}

func (k *KeySpecimen) SetLanguage(language string) {
	k.language = language
}

func (k *KeySpecimen) SetName(name string) {
	k.name = name
}

func (k *KeySpecimen) Type() string {
	return VariableKeySpecimenType
}

func (k *KeySpecimen) ToString(prefix string) string {
	return fmt.Sprintf("%v.%v", k.language, k.name)
}

func NewKeySpecimen() *KeySpecimen {
	return &KeySpecimen{}
}

func NewKeySpecimenWithInit(language string, name string) *KeySpecimen {
	return &KeySpecimen{
		language: language,
		name:     name,
	}
}
