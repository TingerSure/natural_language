package component

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Mapping struct {
	param  *MappingParam
	values map[concept.String]interface{}
}

func (k *Mapping) Size() int {
	return len(k.values)
}

func (k *Mapping) Init(specimen concept.String, defaultValue interface{}) bool {
	if nl_interface.IsNil(defaultValue) {
		defaultValue = k.param.EmptyValue
	}
	exist := k.Iterate(func(key concept.String, value interface{}) bool {
		if key.EqualLanguage(specimen) {
			if nl_interface.IsNil(value) {
				k.values[specimen.Clone()] = defaultValue
			}
			return true
		}
		return false
	})
	if !exist {
		k.values[specimen.Clone()] = defaultValue
	}
	return exist
}

func (k *Mapping) Set(specimen concept.String, value interface{}) bool {
	if nl_interface.IsNil(value) {
		value = k.param.EmptyValue
	}
	exist := k.Iterate(func(key concept.String, _ interface{}) bool {
		if key.EqualLanguage(specimen) {
			k.values[specimen.Clone()] = value
			return true
		}
		return false
	})

	if !exist && k.param.AutoInit {
		k.values[specimen.Clone()] = value
	}

	return exist
}

func (k *Mapping) Get(specimen concept.String) interface{} {
	var choosen interface{} = k.param.EmptyValue
	k.Iterate(func(key concept.String, value interface{}) bool {
		if key.EqualLanguage(specimen) {
			choosen = value
			return true
		}
		return false
	})
	return choosen
}

func (k *Mapping) Has(specimen concept.String) bool {
	return k.Iterate(func(key concept.String, _ interface{}) bool {
		return key.EqualLanguage(specimen)
	})
}

func (k *Mapping) Iterate(on func(concept.String, interface{}) bool) bool {
	for key, value := range k.values {
		if on(key, value) {
			return true
		}
	}
	return false
}

type MappingParam struct {
	AutoInit   bool
	EmptyValue interface{}
}

func NewMapping(param *MappingParam) *Mapping {
	return &Mapping{
		values: make(map[concept.String]interface{}),
		param:  param,
	}
}
