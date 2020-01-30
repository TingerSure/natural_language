package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableClassType = "class"
)

type Class struct {
	name          string
	methods       map[string]concept.Function
	fields        map[string]concept.Variable
	staticMethods map[string]concept.Function
	staticFields  map[string]concept.Variable
}

func (c *Class) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)

	items := []string{}

	for key, value := range c.fields {
		items = append(items, fmt.Sprintf("%vvar %v = %v", subprefix, key, value.ToString(subprefix)))
	}
	for key, value := range c.staticFields {
		items = append(items, fmt.Sprintf("%vstatic var %v = %v", subprefix, key, value.ToString(subprefix)))
	}
	for key, value := range c.methods {
		items = append(items, fmt.Sprintf("%vfunc %v = %v", subprefix, key, value.ToString(subprefix)))
	}
	for key, value := range c.staticMethods {
		items = append(items, fmt.Sprintf("%vstatic func %v = %v", subprefix, key, value.ToString(subprefix)))
	}

	return fmt.Sprintf("class %v {\n%v\n%v}", c.name, strings.Join(items, ",\n"), prefix)
}

func (c *Class) Type() string {
	return VariableClassType
}

func (c *Class) GetName() string {
	return c.name
}

func (c *Class) SetMethod(key string, action concept.Function) {
	c.methods[key] = action
}

func (c *Class) GetMethod(key string) concept.Function {
	return c.methods[key]
}

func (c *Class) HasMethod(key string) bool {
	return !nl_interface.IsNil(c.methods[key])
}

func (c *Class) AllMethods() map[string]concept.Function {
	return c.methods
}

func (c *Class) SetField(key string, defaultField concept.Variable) {
	c.fields[key] = defaultField
}

func (c *Class) GetField(key string) concept.Variable {
	return c.fields[key]
}

func (c *Class) HasField(key string) bool {
	return !nl_interface.IsNil(c.fields[key])
}

func (c *Class) AllFields() map[string]concept.Variable {
	return c.fields
}

func (c *Class) SetStaticMethod(key string, action concept.Function) {
	c.staticMethods[key] = action
}

func (c *Class) GetStaticMethod(key string) concept.Function {
	return c.staticMethods[key]
}

func (c *Class) HasStaticMethod(key string) bool {
	return !nl_interface.IsNil(c.staticMethods[key])
}

func (c *Class) AllStaticMethods() map[string]concept.Function {
	return c.staticMethods
}

func (c *Class) SetStaticField(key string, defaultField concept.Variable) {
	c.staticFields[key] = defaultField
}

func (c *Class) GetStaticField(key string) concept.Variable {
	return c.staticFields[key]
}

func (c *Class) HasStaticField(key string) bool {
	return !nl_interface.IsNil(c.staticFields[key])
}

func (c *Class) AllStaticFields() map[string]concept.Variable {
	return c.staticFields
}

func NewClass(name string) *Class {
	return &Class{
		name:          name,
		methods:       make(map[string]concept.Function),
		fields:        make(map[string]concept.Variable),
		staticMethods: make(map[string]concept.Function),
		staticFields:  make(map[string]concept.Variable),
	}
}
