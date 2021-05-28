package concept

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
)

type Mapping struct {
	param  *MappingParam
	values map[string]*MappingItem
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
	_, yes := k.values[specimen.Value()]
	if !yes {
		if nl_interface.IsNil(defaultValue) {
			defaultValue = k.param.EmptyValue
		}
		k.values[specimen.Value()] = &MappingItem{
			key:   specimen,
			value: defaultValue,
		}
	}
	return yes
}

func (k *Mapping) Remove(specimen String) {
	delete(k.values, specimen.Value())
}

func (k *Mapping) Set(specimen String, value interface{}) bool {
	if nl_interface.IsNil(value) {
		value = k.param.EmptyValue
	}
	item, yes := k.values[specimen.Value()]
	if yes {
		item.value = value
		return true
	}
	if k.param.AutoInit {
		k.values[specimen.Value()] = &MappingItem{
			key:   specimen,
			value: value,
		}
	}
	return false
}

func (k *Mapping) Get(specimen String) interface{} {
	item, yes := k.values[specimen.Value()]
	if !yes {
		return k.param.EmptyValue
	}
	return item.value
}

func (k *Mapping) Key(specimen String) String {
	item, yes := k.values[specimen.Value()]
	if !yes {
		return specimen
	}
	return item.key
}

func (k *Mapping) Has(specimen String) bool {
	_, yes := k.values[specimen.Value()]
	return yes
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
	for _, item := range k.values {
		if on(item.key, item.value) {
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
		values: map[string]*MappingItem{},
		param:  param,
	}
}
