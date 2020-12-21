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
	name         string
	methodMoulds *concept.Mapping
	fieldMoulds  *concept.Mapping
	seed         ClassSeed
}

func (f *Class) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (c *Class) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)

	items := []string{}
	c.fieldMoulds.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vvar %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
		return false
	})
	c.methodMoulds.Iterate(func(key concept.String, value interface{}) bool {
		items = append(items, fmt.Sprintf("%vfunc %v = %v", subprefix, key.ToString(subprefix), value.(concept.ToString).ToString(subprefix)))
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

func (c *Class) SizeMethodMould() int {
	return c.methodMoulds.Size()
}

func (c *Class) SetMethodMould(specimen concept.String, action concept.Function) {
	c.methodMoulds.Set(specimen, action)
}

func (c *Class) GetMethodMould(specimen concept.String) concept.Function {
	return c.methodMoulds.Get(specimen).(concept.Function)
}

func (c *Class) HasMethodMould(specimen concept.String) bool {
	return c.methodMoulds.Has(specimen)
}

func (c *Class) IterateMethodMoulds(on func(key concept.String, value concept.Function) bool) bool {
	return c.methodMoulds.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Function))
	})
}

func (c *Class) SizeFieldMould() int {
	return c.fieldMoulds.Size()
}

func (c *Class) SetFieldMould(specimen concept.String, defaultFieldMould concept.Variable) {
	c.fieldMoulds.Set(specimen, defaultFieldMould)
}

func (c *Class) GetFieldMould(specimen concept.String) concept.Variable {
	return c.fieldMoulds.Get(specimen).(concept.Variable)
}

func (c *Class) HasFieldMould(specimen concept.String) bool {
	return c.fieldMoulds.Has(specimen)
}

func (c *Class) IterateFieldMoulds(on func(key concept.String, value concept.Variable) bool) bool {
	return c.fieldMoulds.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Variable))
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
		methodMoulds: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewNull(),
		}),
		fieldMoulds: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewNull(),
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
