package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable/component"
	"strings"
)

const (
	VariableObjectType = "object"
)

type Object struct {
	fields      *component.Mapping
	methods     *component.Mapping
	reflections []*component.ClassReflection
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
func (o *Object) CheckMapping(class concept.Class, mapping map[concept.KeySpecimen]concept.KeySpecimen) bool {
	if len(mapping) != class.SizeField() {
		return false
	}

	if class.IterateFields(func(key concept.Key, _ concept.Variable) bool {
		for specimen, _ := range mapping {
			if key.Is(specimen) {
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
func (o *Object) GetMapping(class string, alias string) (map[concept.KeySpecimen]concept.KeySpecimen, concept.Exception) {
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == alias {
			return reflection.GetMapping(), nil
		}
	}
	return nil, interrupt.NewException("system error", fmt.Sprintf("No mapping who's class is \"%v\" and alias is \"%v\"", class, alias))
}

func (o *Object) RemoveClass(class string, alias string) concept.Exception {
	for index, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == alias {
			o.reflections = append(o.reflections[:index], o.reflections[index+1:]...)
			return nil
		}
	}
	return interrupt.NewException("system error", fmt.Sprintf("No class who's name is \"%v\" and alias is \"%v\"", class, alias))

}

func (o *Object) AddClass(class concept.Class, alias string, mapping map[concept.KeySpecimen]concept.KeySpecimen) concept.Exception {
	for _, old := range o.reflections {
		if old.GetClass().GetName() == class.GetName() && old.GetAlias() == alias {
			return interrupt.NewException("system error", "Duplicate class reflections are added.")
		}
	}
	o.reflections = append(o.reflections, component.NewClassReflectionWithMapping(class, mapping, alias))
	return nil
}

func (o *Object) HasField(specimen concept.KeySpecimen) bool {
	return o.fields.Has(specimen)
}

func (o *Object) InitField(specimen concept.KeySpecimen, defaultValue concept.Variable) concept.Exception {
	if !o.fields.Has(specimen) {
		o.fields.Set(specimen, defaultValue)
	}
	return nil
}

func (o *Object) SetField(specimen concept.KeySpecimen, value concept.Variable) concept.Exception {
	if !o.fields.Has(specimen) {
		return interrupt.NewException("system error", fmt.Sprintf("There is no field called \"%v\" to be set here.", specimen.ToString("")))
	}
	o.fields.Set(specimen, value)
	return nil
}

func (o *Object) GetField(specimen concept.KeySpecimen) (concept.Variable, concept.Exception) {
	value := o.fields.Get(specimen)
	if nl_interface.IsNil(value) {
		return nil, interrupt.NewException("system error", fmt.Sprintf("There is no field called \"%v\" to be got here.", specimen.ToString("")))
	}
	return value.(concept.Variable), nil
}

func (o *Object) HasMethod(specimen concept.KeySpecimen) bool {
	return o.methods.Has(specimen)
}

func (o *Object) SetMethod(specimen concept.KeySpecimen, value concept.Function) concept.Exception {
	o.methods.Set(specimen, value)
	return nil
}

func (o *Object) GetMethod(specimen concept.KeySpecimen) (concept.Function, concept.Exception) {
	value := o.methods.Get(specimen)
	if nl_interface.IsNil(value) {
		return nil, interrupt.NewException("system error", fmt.Sprintf("no method called %v", specimen.ToString("")))
	}
	return value.(concept.Function), nil
}

func (a *Object) ToString(prefix string) string {
	if 0 == a.fields.Size() && 0 == len(a.reflections) {
		return "object <> {}"
	}

	subPrefix := fmt.Sprintf("%v\t", prefix)

	paramsToString := make([]string, 0, a.fields.Size())

	a.fields.Iterate(func(key concept.Key, value interface{}) bool {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.(concept.ToString).ToString(subPrefix)))
		return false
	})

	a.methods.Iterate(func(key concept.Key, value interface{}) bool {
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
	return VariableObjectType
}

func NewObject() *Object {
	keySpecimenCreator := func() concept.KeySpecimen {
		return NewKeySpecimen()
	}

	keyCreator := func() concept.Key {
		return NewKey()
	}
	return &Object{
		fields: component.NewMapping(&component.MappingParam{
			KeySpecimenCreator: keySpecimenCreator,
			KeyCreator:         keyCreator,
			AutoInit:           false,
		}),
		methods: component.NewMapping(&component.MappingParam{
			KeySpecimenCreator: keySpecimenCreator,
			KeyCreator:         keyCreator,
			AutoInit:           true,
		}),
		reflections: make([]*component.ClassReflection, 0),
	}
}
