package adaptor

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type AdaptorVariableParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type AdaptorVariable struct {
	fields      *concept.Mapping
	methods     *concept.Mapping
	reflections []*concept.Reflection
	param       *AdaptorVariableParam
}

func (o *AdaptorVariable) GetClasses() []string {
	classes := []string{}
	for _, reflection := range o.reflections {
		classes = append(classes, reflection.GetClass().GetName())
	}
	return classes
}

func (o *AdaptorVariable) GetClass(className string) concept.Class {

	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == className {
			return reflection.GetClass()
		}
	}
	return nil
}

func (o *AdaptorVariable) GetAliases(class string) []string {
	aliases := []string{}
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class {
			aliases = append(aliases, reflection.GetAlias())
		}
	}
	return aliases
}

func (o *AdaptorVariable) IsClassAlias(class string, alias string) bool {
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == alias {
			return true
		}
	}
	return false
}

func (o *AdaptorVariable) UpdateAlias(class string, old, new string) bool {
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == old {
			reflection.SetAlias(new)
			return true
		}
	}
	return false
}

func (o *AdaptorVariable) CheckMapping(class concept.Class, mapping map[concept.String]concept.String) bool {
	if len(mapping) != class.SizeFieldMould() {
		return false
	}

	if class.IterateFieldMoulds(func(key concept.String, _ concept.Variable) bool {
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

func (o *AdaptorVariable) GetMapping(class string, alias string) (map[concept.String]concept.String, concept.Exception) {
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == alias {
			return reflection.GetMapping(), nil
		}
	}
	return nil, o.param.ExceptionCreator("system error", fmt.Sprintf("No mapping who's class is \"%v\" and alias is \"%v\"", class, alias))
}

func (o *AdaptorVariable) RemoveClass(class string, alias string) concept.Exception {
	for index, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == alias {
			o.reflections = append(o.reflections[:index], o.reflections[index+1:]...)
			return nil
		}
	}
	return o.param.ExceptionCreator("system error", fmt.Sprintf("No class who's name is \"%v\" and alias is \"%v\"", class, alias))
}

func (o *AdaptorVariable) AddClass(class concept.Class, alias string, mapping map[concept.String]concept.String) concept.Exception {
	for _, old := range o.reflections {
		if old.GetClass().GetName() == class.GetName() && old.GetAlias() == alias {
			return o.param.ExceptionCreator("system error", "Duplicate class reflections are added.")
		}
	}
	o.reflections = append(o.reflections, concept.NewReflectionWithMapping(class, mapping, alias))
	return nil
}

func (o *AdaptorVariable) HasField(specimen concept.String) bool {
	o.initFields()
	return o.fields.Has(specimen)
}

func (o *AdaptorVariable) InitField(specimen concept.String, defaultValue concept.Variable) concept.Exception {
	o.initFields()
	o.fields.Init(specimen, defaultValue)
	return nil
}

func (o *AdaptorVariable) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	o.initFields()
	if !o.fields.Has(specimen) {
		return o.param.ExceptionCreator("system error", fmt.Sprintf("There is no field called \"%v\" to be set here.", specimen.ToString("")))
	}
	o.fields.Set(specimen, value)
	return nil
}

func (o *AdaptorVariable) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	o.initMethods()
	value := o.fields.Get(specimen)
	if nl_interface.IsNil(value) {
		return nil, o.param.ExceptionCreator("system error", fmt.Sprintf("There is no field called \"%v\" to be got here.", specimen.ToString("")))
	}
	return value.(concept.Variable), nil
}

func (o *AdaptorVariable) HasMethod(specimen concept.String) bool {
	o.initMethods()
	return o.methods.Has(specimen)
}

func (o *AdaptorVariable) SetMethod(specimen concept.String, value concept.Function) concept.Exception {
	o.initMethods()
	o.methods.Set(specimen, value)
	return nil
}

func (o *AdaptorVariable) GetMethod(specimen concept.String) (concept.Function, concept.Exception) {
	o.initMethods()
	value := o.methods.Get(specimen)
	if nl_interface.IsNil(value) {
		return nil, o.param.ExceptionCreator("system error", fmt.Sprintf("no method called %v", specimen.ToString("")))
	}
	return value.(concept.Function), nil
}

func (a *AdaptorVariable) ToString(prefix string) string {
	if (a.fields == nil || 0 == a.fields.Size()) &&
		(a.methods == nil || 0 == a.methods.Size()) &&
		0 == len(a.reflections) {
		return "object <> {}"
	}

	subPrefix := fmt.Sprintf("%v\t", prefix)

	paramsToString := make([]string, 0, a.fields.Size())
	if a.fields != nil {
		a.fields.Iterate(func(key concept.String, value interface{}) bool {
			paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.(concept.ToString).ToString(subPrefix)))
			return false
		})
	}
	if a.methods != nil {
		a.methods.Iterate(func(key concept.String, value interface{}) bool {
			paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.(concept.ToString).ToString(subPrefix)))
			return false
		})
	}
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

func (o *AdaptorVariable) initMethods() {
	if o.methods == nil {
		o.methods = concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: o.param.NullCreator(),
		})
	}
}

func (o *AdaptorVariable) initFields() {
	if o.fields == nil {
		o.fields = concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: o.param.NullCreator(),
		})
	}
}

func NewAdaptorVariable(param *AdaptorVariableParam) *AdaptorVariable {
	return &AdaptorVariable{
		reflections: make([]*concept.Reflection, 0),
		param:       param,
	}
}
