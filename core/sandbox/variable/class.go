package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableClassType = "class"
)

type ClassSeed interface {
	ToLanguage(string, *Class) string
	Type() string
	NewEmpty() concept.Null
}

type Class struct {
	name          string
	methods       *concept.Mapping
	fields        *concept.Mapping
	staticMethods *concept.Mapping
	staticFields  *concept.Mapping
	seed          ClassSeed
}

func (f *Class) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (c *Class) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)

	items := []string{}

	c.fields.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vvar %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	c.staticFields.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vstatic var %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	c.methods.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vfunc %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	c.staticMethods.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vstatic func %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})

	return fmt.Sprintf("class %v {\n%v\n%v}", c.name, strings.Join(items, ",\n"), prefix)
}

func (c *Class) Type() string {
	return c.seed.Type()
}

func (c *Class) GetName() string {
	return c.name
}

func (c *Class) SizeMethod() int {
	return c.methods.Size()
}

func (c *Class) SetMethod(specimen concept.String, action concept.Function) {
	c.methods.Set(specimen, action)
}

func (c *Class) GetMethod(specimen concept.String) concept.Function {
	return c.methods.Get(specimen).(concept.Function)
}

func (c *Class) HasMethod(specimen concept.String) bool {
	return c.methods.Has(specimen)
}

func (c *Class) IterateMethods(on func(key concept.String, value concept.Function) bool) bool {
	return c.methods.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Function))
	})
}

func (c *Class) SizeField() int {
	return c.fields.Size()
}

func (c *Class) SetField(specimen concept.String, defaultField concept.Variable) {
	c.fields.Set(specimen, defaultField)
}

func (c *Class) GetField(specimen concept.String) concept.Variable {
	return c.fields.Get(specimen).(concept.Variable)
}

func (c *Class) HasField(specimen concept.String) bool {
	return c.fields.Has(specimen)
}

func (c *Class) IterateFields(on func(key concept.String, value concept.Variable) bool) bool {
	return c.fields.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Variable))
	})
}

func (c *Class) SizeStaticMethod() int {
	return c.staticMethods.Size()
}

func (c *Class) SetStaticMethod(specimen concept.String, action concept.Function) {
	c.staticMethods.Set(specimen, action)
}

func (c *Class) GetStaticMethod(specimen concept.String) concept.Function {
	return c.staticMethods.Get(specimen).(concept.Function)
}

func (c *Class) HasStaticMethod(specimen concept.String) bool {
	return c.staticMethods.Has(specimen)
}

func (c *Class) IterateStaticMethods(on func(key concept.String, value concept.Function) bool) bool {
	return c.staticMethods.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Function))
	})
}

func (c *Class) SizeStaticField() int {
	return c.staticFields.Size()
}

func (c *Class) SetStaticField(specimen concept.String, defaultField concept.Variable) {
	c.staticFields.Set(specimen, defaultField)
}

func (c *Class) GetStaticField(specimen concept.String) concept.Variable {
	return c.staticFields.Get(specimen).(concept.Variable)
}

func (c *Class) HasStaticField(specimen concept.String) bool {
	return c.staticFields.Has(specimen)
}

func (c *Class) IterateStaticFields(on func(key concept.String, value concept.Variable) bool) bool {
	return c.staticFields.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Variable))
	})
}

type ClassCreatorParam struct {
	NullCreator func() concept.Null
}

type ClassCreator struct {
	Seeds map[string]func(string, *Class) string
	param *ClassCreatorParam
}

func (s *ClassCreator) New(name string) *Class {
	return &Class{
		name: name,
		methods: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewEmpty(),
		}),
		fields: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewEmpty(),
		}),
		staticMethods: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewEmpty(),
		}),
		staticFields: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewEmpty(),
		}),
		seed: s,
	}
}

func (s *ClassCreator) ToLanguage(language string, instance *Class) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ClassCreator) Type() string {
	return VariableClassType
}

func (s *ClassCreator) NewEmpty() concept.Null {
	return s.param.NullCreator()
}

func NewClassCreator(param *ClassCreatorParam) *ClassCreator {
	return &ClassCreator{
		Seeds: map[string]func(string, *Class) string{},
		param: param,
	}
}
