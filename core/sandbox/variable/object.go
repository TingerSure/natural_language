package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableObjectType = "object"
)

type ObjectSeed interface {
	ToLanguage(string, *Object) string
	Type() string
	NewEmpty() concept.Null
	NewException(string, string) concept.Exception
}

type Object struct {
	fields      *concept.Mapping
	methods     *concept.Mapping
	reflections []*concept.ClassReflection
	seed        ObjectSeed
}

func (f *Object) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (o *Object) GetClasses() []string {
	classes := []string{}
	for _, reflection := range o.reflections {
		classes = append(classes, reflection.GetClass().GetName())
	}
	return classes
}

func (o *Object) GetClass(className string) concept.Class {

	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == className {
			return reflection.GetClass()
		}
	}
	return nil
}

func (o *Object) GetAliases(class string) []string {
	aliases := []string{}
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class {
			aliases = append(aliases, reflection.GetAlias())
		}
	}
	return aliases
}

func (o *Object) IsClassAlias(class string, alias string) bool {
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == alias {
			return true
		}
	}
	return false
}

func (o *Object) UpdateAlias(class string, old, new string) bool {
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == old {
			reflection.SetAlias(new)
			return true
		}
	}
	return false
}
func (o *Object) CheckMapping(class concept.Class, mapping map[concept.String]concept.String) bool {
	if len(mapping) != class.SizeField() {
		return false
	}

	if class.IterateFields(func(key concept.String, _ concept.Variable) bool {
		for specimen := range mapping {
			if key.EqualLanguage(specimen) {
				return false
			}
		}
		return true
	}) {
		return false
	}

	for _, specimen := range mapping {
		if !o.HasField(specimen) {
			return false
		}
	}
	return true
}
func (o *Object) GetMapping(class string, alias string) (map[concept.String]concept.String, concept.Exception) {
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == alias {
			return reflection.GetMapping(), nil
		}
	}
	return nil, o.seed.NewException("system error", fmt.Sprintf("No mapping who's class is \"%v\" and alias is \"%v\"", class, alias))
}

func (o *Object) RemoveClass(class string, alias string) concept.Exception {
	for index, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == alias {
			o.reflections = append(o.reflections[:index], o.reflections[index+1:]...)
			return nil
		}
	}
	return o.seed.NewException("system error", fmt.Sprintf("No class who's name is \"%v\" and alias is \"%v\"", class, alias))

}

func (o *Object) AddClass(class concept.Class, alias string, mapping map[concept.String]concept.String) concept.Exception {
	for _, old := range o.reflections {
		if old.GetClass().GetName() == class.GetName() && old.GetAlias() == alias {
			return o.seed.NewException("system error", "Duplicate class reflections are added.")
		}
	}
	o.reflections = append(o.reflections, concept.NewClassReflectionWithMapping(class, mapping, alias))
	return nil
}

func (o *Object) HasField(specimen concept.String) bool {
	return o.fields.Has(specimen)
}

func (o *Object) InitField(specimen concept.String, defaultValue concept.Variable) concept.Exception {
	o.fields.Init(specimen, defaultValue)
	return nil
}

func (o *Object) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	if !o.fields.Has(specimen) {
		return o.seed.NewException("system error", fmt.Sprintf("There is no field called \"%v\" to be set here.", specimen.ToString("")))
	}
	o.fields.Set(specimen, value)
	return nil
}

func (o *Object) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	value := o.fields.Get(specimen)
	if nl_interface.IsNil(value) {
		return nil, o.seed.NewException("system error", fmt.Sprintf("There is no field called \"%v\" to be got here.", specimen.ToString("")))
	}
	return value.(concept.Variable), nil
}

func (o *Object) HasMethod(specimen concept.String) bool {
	return o.methods.Has(specimen)
}

func (o *Object) SetMethod(specimen concept.String, value concept.Function) concept.Exception {
	o.methods.Set(specimen, value)
	return nil
}

func (o *Object) GetMethod(specimen concept.String) (concept.Function, concept.Exception) {
	value := o.methods.Get(specimen)
	if nl_interface.IsNil(value) {
		return nil, o.seed.NewException("system error", fmt.Sprintf("no method called %v", specimen.ToString("")))
	}
	return value.(concept.Function), nil
}

func (a *Object) ToString(prefix string) string {
	if 0 == a.fields.Size() && 0 == len(a.reflections) {
		return "object <> {}"
	}

	subPrefix := fmt.Sprintf("%v\t", prefix)

	paramsToString := make([]string, 0, a.fields.Size())

	a.fields.Iterate(func(key concept.String, value interface{}) bool {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.(concept.ToString).ToString(subPrefix)))
		return false
	})

	a.methods.Iterate(func(key concept.String, value interface{}) bool {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.(concept.ToString).ToString(subPrefix)))
		return false
	})

	reflectionToString := make([]string, 0, len(a.reflections))
	for _, reflection := range a.reflections {
		aliasToString := ""
		if reflection.GetAlias() != "" {
			aliasToString = fmt.Sprintf("(%v)", reflection.GetAlias())
		}
		reflectionToString = append(reflectionToString, fmt.Sprintf("%v%v", reflection.GetClass().GetName(), aliasToString))
	}

	return fmt.Sprintf("object <%v> {\n%v\n%v}", strings.Join(reflectionToString, ", "), strings.Join(paramsToString, ",\n"), prefix)
}

func (o *Object) Type() string {
	return o.seed.Type()
}

type ObjectCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type ObjectCreator struct {
	Seeds map[string]func(string, *Object) string
	param *ObjectCreatorParam
}

func (s *ObjectCreator) New() *Object {
	return &Object{
		fields: concept.NewMapping(&concept.MappingParam{
			AutoInit:   false,
			EmptyValue: s.NewEmpty(),
		}),
		methods: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.NewEmpty(),
		}),
		reflections: make([]*concept.ClassReflection, 0),
		seed:        s,
	}
}

func (s *ObjectCreator) ToLanguage(language string, instance *Object) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ObjectCreator) Type() string {
	return VariableObjectType
}

func (s *ObjectCreator) NewEmpty() concept.Null {
	return s.param.NullCreator()
}
func (s *ObjectCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewObjectCreator(param *ObjectCreatorParam) *ObjectCreator {
	return &ObjectCreator{
		Seeds: map[string]func(string, *Object) string{},
		param: param,
	}
}
