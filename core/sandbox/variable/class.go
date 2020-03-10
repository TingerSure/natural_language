package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/component"
	"strings"
)

const (
	VariableClassType = "class"
)

type Class struct {
	name          string
	methods       *component.Mapping
	fields        *component.Mapping
	staticMethods *component.Mapping
	staticFields  *component.Mapping
}

func (c *Class) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)

	items := []string{}

	c.fields.Iterate(func(key concept.Key, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vvar %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	c.staticFields.Iterate(func(key concept.Key, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vstatic var %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	c.methods.Iterate(func(key concept.Key, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vfunc %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	c.staticMethods.Iterate(func(key concept.Key, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vstatic func %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})

	return fmt.Sprintf("class %v {\n%v\n%v}", c.name, strings.Join(items, ",\n"), prefix)
}

func (c *Class) Type() string {
	return VariableClassType
}

func (c *Class) GetName() string {
	return c.name
}

func (c *Class) SizeMethod() int {
	return c.methods.Size()
}

func (c *Class) SetMethod(specimen concept.KeySpecimen, action concept.Function) {
	c.methods.Set(specimen, action)
}

func (c *Class) GetMethod(specimen concept.KeySpecimen) concept.Function {
	return c.methods.Get(specimen).(concept.Function)
}

func (c *Class) HasMethod(specimen concept.KeySpecimen) bool {
	return c.methods.Has(specimen)
}

func (c *Class) IterateMethods(on func(key concept.Key, value concept.Function) bool) bool {
	return c.methods.Iterate(func(key concept.Key, value interface{}) bool {
		return on(key, value.(concept.Function))
	})
}

func (c *Class) SizeField() int {
	return c.fields.Size()
}

func (c *Class) SetField(specimen concept.KeySpecimen, defaultField concept.Variable) {
	c.fields.Set(specimen, defaultField)
}

func (c *Class) GetField(specimen concept.KeySpecimen) concept.Variable {
	return c.fields.Get(specimen).(concept.Variable)
}

func (c *Class) HasField(specimen concept.KeySpecimen) bool {
	return c.fields.Has(specimen)
}

func (c *Class) IterateFields(on func(key concept.Key, value concept.Variable) bool) bool {
	return c.fields.Iterate(func(key concept.Key, value interface{}) bool {
		return on(key, value.(concept.Variable))
	})
}

func (c *Class) SizeStaticMethod() int {
	return c.staticMethods.Size()
}

func (c *Class) SetStaticMethod(specimen concept.KeySpecimen, action concept.Function) {
	c.staticMethods.Set(specimen, action)
}

func (c *Class) GetStaticMethod(specimen concept.KeySpecimen) concept.Function {
	return c.staticMethods.Get(specimen).(concept.Function)
}

func (c *Class) HasStaticMethod(specimen concept.KeySpecimen) bool {
	return c.staticMethods.Has(specimen)
}

func (c *Class) IterateStaticMethods(on func(key concept.Key, value concept.Function) bool) bool {
	return c.staticMethods.Iterate(func(key concept.Key, value interface{}) bool {
		return on(key, value.(concept.Function))
	})
}

func (c *Class) SizeStaticField() int {
	return c.staticFields.Size()
}

func (c *Class) SetStaticField(specimen concept.KeySpecimen, defaultField concept.Variable) {
	c.staticFields.Set(specimen, defaultField)
}

func (c *Class) GetStaticField(specimen concept.KeySpecimen) concept.Variable {
	return c.staticFields.Get(specimen).(concept.Variable)
}

func (c *Class) HasStaticField(specimen concept.KeySpecimen) bool {
	return c.staticFields.Has(specimen)
}

func (c *Class) IterateStaticFields(on func(key concept.Key, value concept.Variable) bool) bool {
	return c.staticFields.Iterate(func(key concept.Key, value interface{}) bool {
		return on(key, value.(concept.Variable))
	})
}

func NewClass(name string) *Class {

	keySpecimenCreator := func() concept.KeySpecimen {
		return NewKeySpecimen()
	}

	keyCreator := func() concept.Key {
		return NewKey()
	}

	return &Class{
		name: name,
		methods: component.NewMapping(&component.MappingParam{
			KeySpecimenCreator: keySpecimenCreator,
			KeyCreator:         keyCreator,
			AutoInit:           true,
		}),
		fields: component.NewMapping(&component.MappingParam{
			KeySpecimenCreator: keySpecimenCreator,
			KeyCreator:         keyCreator,
			AutoInit:           true,
		}),
		staticMethods: component.NewMapping(&component.MappingParam{
			KeySpecimenCreator: keySpecimenCreator,
			KeyCreator:         keyCreator,
			AutoInit:           true,
		}),
		staticFields: component.NewMapping(&component.MappingParam{
			KeySpecimenCreator: keySpecimenCreator,
			KeyCreator:         keyCreator,
			AutoInit:           true,
		}),
	}
}
