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
	ToLanguage(string, concept.Pool, *Class) string
	Type() string
	NewNull() concept.Null
}

type Class struct {
	*adaptor.AdaptorVariable
	provide *concept.Mapping
	require *concept.Mapping
	seed    ClassSeed
}

func (o *Class) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.CallAdaptor(specimen, param, o)
}

func (f *Class) ToLanguage(language string, space concept.Pool) string {
	return f.seed.ToLanguage(language, space, f)
}

func (c *Class) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)

	items := []string{}
	c.require.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vrequire %v = %v", subprefix, key.Value(), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	c.provide.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vprovide %v = %v", subprefix, key.Value(), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	return fmt.Sprintf("class {\n%v\n%v}", strings.Join(items, "\n"), prefix)
}

func (c *Class) Type() string {
	return c.seed.Type()
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

func (c *Class) SetRequire(specimen concept.String, define concept.Function) {
	c.require.Set(specimen, define)
}

func (c *Class) GetRequire(specimen concept.String) concept.Function {
	return c.require.Get(specimen).(concept.Function)
}

func (c *Class) HasRequire(specimen concept.String) bool {
	return c.require.Has(specimen)
}

func (c *Class) IterateRequire(on func(key concept.String, value concept.Function) bool) bool {
	return c.require.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Function))
	})
}

type ClassCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type ClassCreator struct {
	Seeds map[string]func(string, concept.Pool, *Class) string
	param *ClassCreatorParam
}

func (s *ClassCreator) New() *Class {
	return &Class{
		AdaptorVariable: adaptor.NewAdaptorVariable(&adaptor.AdaptorVariableParam{
			NullCreator:      s.param.NullCreator,
			ExceptionCreator: s.param.ExceptionCreator,
		}),
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

func (s *ClassCreator) ToLanguage(language string, space concept.Pool, instance *Class) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *ClassCreator) Type() string {
	return VariableClassType
}

func (s *ClassCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func NewClassCreator(param *ClassCreatorParam) *ClassCreator {
	return &ClassCreator{
		Seeds: map[string]func(string, concept.Pool, *Class) string{},
		param: param,
	}
}
