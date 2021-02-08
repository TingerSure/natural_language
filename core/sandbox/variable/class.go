package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable/adaptor"
	"strings"
)

const (
	VariableClassType = "class"
)

type ClassSeed interface {
	ToLanguage(string, *Class) string
	Type() string
	NewNull() concept.Null
}

type Class struct {
	*adaptor.AdaptorVariable
	name    string
	provide *concept.Mapping
	require *concept.Mapping
	seed    ClassSeed
}

func (f *Class) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (c *Class) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)

	items := []string{}
	c.require.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vrequire %v", subprefix, key.ToString(subprefix)))
		return false
	})
	c.provide.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vprovide %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
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

func (c *Class) SetProvide(specimen concept.String, action concept.Function) {
	c.provide.Set(specimen, action)
}

func (c *Class) GetProvide(specimen concept.String) concept.Function {
	return c.provide.Get(specimen).(concept.Function)
}

func (c *Class) HasProvide(specimen concept.String) bool {
	return c.provide.Has(specimen)
}

func (c *Class) IterateProvide(on func(key concept.String, value concept.Function) bool) bool {
	return c.provide.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Function))
	})
}

func (c *Class) SetRequire(specimen concept.String) {
	c.require.Set(specimen, nil)
}

func (c *Class) RemoveRequire(specimen concept.String) {
	c.require.Remove(specimen)
}

func (c *Class) HasRequire(specimen concept.String) bool {
	return c.require.Has(specimen)
}

func (c *Class) IterateRequire(on func(key concept.String) bool) bool {
	return c.require.Iterate(func(key concept.String, value interface{}) bool {
		return on(key)
	})
}

type ClassCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type ClassCreator struct {
	Seeds map[string]func(string, *Class) string
	param *ClassCreatorParam
}

func (s *ClassCreator) New(name string) *Class {
	return &Class{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
		name: name,
		provide: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewNull(),
		}),
		require: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: nil,
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

func (s *ClassCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewClassCreator(param *ClassCreatorParam) *ClassCreator {
	return &ClassCreator{
		Seeds: map[string]func(string, *Class) string{},
		param: param,
	}
}
