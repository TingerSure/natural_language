package concept

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
)

type Mapping struct {
	param  *MappingParam
	values []*MappingItem
}

type MappingItem struct {
	key   String
	value interface{}
}

func (m *Mapping) Param() *MappingParam {
	return m.param
}

func (k *Mapping) Size() int {
	return len(k.values)
}

func (k *Mapping) Init(specimen String, defaultValue interface{}) bool {
	if nl_interface.IsNil(defaultValue) {
		defaultValue = k.param.EmptyValue
	}
	exist := k.iterate(func(item *MappingItem) bool {
		if item.key.EqualLanguage(specimen) {
			if nl_interface.IsNil(item.value) {
				item.value = defaultValue
			}
			return true
		}
		return false
	})
	if !exist {
		k.values = append(k.values, &MappingItem{
			key:   specimen,
			value: defaultValue,
		})
	}
	return exist
}

func (k *Mapping) Set(specimen String, value interface{}) bool {
	if nl_interface.IsNil(value) {
		value = k.param.EmptyValue
	}
	exist := k.iterate(func(item *MappingItem) bool {
		if item.key.EqualLanguage(specimen) {
			item.value = value
			return true
		}
		return false
	})

	if !exist && k.param.AutoInit {
		k.values = append(k.values, &MappingItem{
			key:   specimen,
			value: value,
		})
	}

	return exist
}

func (k *Mapping) Get(specimen String) interface{} {
	var choosen interface{} = k.param.EmptyValue
	k.iterate(func(item *MappingItem) bool {
		if item.key.EqualLanguage(specimen) {
			choosen = item.value
			return true
		}
		return false
	})
	return choosen
}

func (k *Mapping) Has(specimen String) bool {
	return k.iterate(func(item *MappingItem) bool {
		return item.key.EqualLanguage(specimen)
	})
}

func (k *Mapping) iterate(on func(item *MappingItem) bool) bool {
	for _, item := range k.values {
		if on(item) {
			return true
		}
	}
	return false
}

func (k *Mapping) Iterate(on func(key String, value interface{}) bool) bool {
	return k.iterate(func(item *MappingItem) bool {
		return on(item.key, item.value)
	})
}

type MappingParam struct {
	AutoInit   bool
	EmptyValue interface{}
}

func NewMapping(param *MappingParam) *Mapping {
	return &Mapping{
		values: make([]*MappingItem, 0),
		param:  param,
	}
}
