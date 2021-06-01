package nl_interface

import (
	"fmt"
	"strings"
)

type Key interface {
	MapKey() string
}

type Mapping struct {
	param  *MappingParam
	values map[string]*MappingItem
}

type MappingItem struct {
	key   Key
	value interface{}
}

func (m *Mapping) Param() *MappingParam {
	return m.param
}

func (k *Mapping) Size() int {
	return len(k.values)
}

func (k *Mapping) Init(specimen Key, defaultValue interface{}) bool {
	_, yes := k.values[specimen.MapKey()]
	if !yes {
		if IsNil(defaultValue) {
			defaultValue = k.param.EmptyValue
		}
		k.values[specimen.MapKey()] = &MappingItem{
			key:   specimen,
			value: defaultValue,
		}
	}
	return yes
}

func (k *Mapping) Remove(specimen Key) {
	delete(k.values, specimen.MapKey())
}

func (k *Mapping) Set(specimen Key, value interface{}) bool {
	if IsNil(value) {
		value = k.param.EmptyValue
	}
	item, yes := k.values[specimen.MapKey()]
	if yes {
		item.value = value
		return true
	}
	if k.param.AutoInit {
		k.values[specimen.MapKey()] = &MappingItem{
			key:   specimen,
			value: value,
		}
	}
	return false
}

func (k *Mapping) Get(specimen Key) interface{} {
	item, yes := k.values[specimen.MapKey()]
	if !yes {
		return k.param.EmptyValue
	}
	return item.value
}

func (k *Mapping) Key(specimen Key) Key {
	item, yes := k.values[specimen.MapKey()]
	if !yes {
		return specimen
	}
	return item.key
}

func (k *Mapping) Has(specimen Key) bool {
	_, yes := k.values[specimen.MapKey()]
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

func (k *Mapping) Iterate(on func(key Key, value interface{}) bool) bool {
	for _, item := range k.values {
		if on(item.key, item.value) {
			return true
		}
	}
	return false
}

func (k *Mapping) ToString(prefix string) string {
	keys := []string{}
	for key, _ := range k.values {
		keys = append(keys, key)
	}
	return fmt.Sprintf("{{%v}}", strings.Join(keys, ", "))
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
