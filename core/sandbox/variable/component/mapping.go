package component

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Mapping struct {
	autoInit           bool
	values             map[concept.Key]interface{}
	keyCreator         func() concept.Key
	keySpecimenCreator func() concept.KeySpecimen
}

func (k *Mapping) Size() int {
	return len(k.values)
}

func (k *Mapping) Set(specimen concept.KeySpecimen, value interface{}) bool {
	exist := k.Iterate(func(key concept.Key, _ interface{}) bool {
		if key.Is(specimen) {
			k.values[key] = value
			return true
		}
		return false
	})

	if !exist && k.autoInit {
		key := k.keyCreator()
		key.Set(specimen)
		k.values[key] = value
	}

	return exist
}

func (k *Mapping) Get(specimen concept.KeySpecimen) interface{} {
	var choosen interface{} = nil
	k.Iterate(func(key concept.Key, value interface{}) bool {
		if key.Is(specimen) {
			choosen = value
			return true
		}
		return false
	})
	return choosen
}

func (k *Mapping) Has(specimen concept.KeySpecimen) bool {
	return k.Iterate(func(key concept.Key, _ interface{}) bool {
		return key.Is(specimen)
	})
}

func (k *Mapping) Iterate(on func(concept.Key, interface{}) bool) bool {
	for key, value := range k.values {
		if on(key, value) {
			return true
		}
	}
	return false
}

func (k *Mapping) IterateLanguage(
	on func(specimen concept.KeySpecimen, value interface{}) bool,
	language string,
) bool {
	return k.Iterate(func(key concept.Key, value interface{}) bool {
		specimen := k.keySpecimenCreator()
		specimen.SetLanguage(language)
		key.Get(specimen)
		return on(specimen, value)
	})
}

type MappingParam struct {
	KeySpecimenCreator func() concept.KeySpecimen
	KeyCreator         func() concept.Key
	AutoInit           bool
}

func NewMapping(param *MappingParam) *Mapping {
	return &Mapping{
		values:             make(map[concept.Key]interface{}),
		autoInit:           param.AutoInit,
		keyCreator:         param.KeyCreator,
		keySpecimenCreator: param.KeySpecimenCreator,
	}
}
