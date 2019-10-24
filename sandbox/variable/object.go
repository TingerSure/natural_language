package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable/component"
	"strings"
)

const (
	VariableObjectType = "object"
)

type Object struct {
	fields      map[string]concept.Variable
	methods     map[string]concept.Function
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

func (o *Object) UpdateAlias(class string, old, new string) bool {
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class && reflection.GetAlias() == old {
			reflection.SetAlias(new)
			return true
		}
	}
	return false
}

func (o *Object) GetMapping(class string, alias string) (map[string]string, concept.Exception) {
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

func (o *Object) AddClass(class concept.Class, alias string, mapping map[string]string) concept.Exception {
	for _, old := range o.reflections {
		if old.GetClass().GetName() == class.GetName() && old.GetAlias() == alias {
			return interrupt.NewException("system error", "Duplicate class reflections are added.")
		}
	}
	o.reflections = append(o.reflections, component.NewClassReflectionWithMapping(class, mapping, alias))
	return nil
}

func (o *Object) HasField(key string) bool {
	return o.fields[key] != nil
}

func (o *Object) InitField(key string, defaultValue concept.Variable) concept.Exception {
	if o.fields[key] == nil {
		o.fields[key] = defaultValue
	}
	return nil
}

func (o *Object) SetField(key string, value concept.Variable) concept.Exception {
	if o.fields[key] == nil {
		return interrupt.NewException("system error", fmt.Sprintf("There is no field called \"%v\" to be set here.", key))
	}
	o.fields[key] = value
	return nil
}

func (o *Object) GetField(key string) (concept.Variable, concept.Exception) {
	if o.fields[key] == nil {
		return nil, interrupt.NewException("system error", fmt.Sprintf("There is no field called \"%v\" to be got here.", key))
	}
	return o.fields[key], nil
}

func (o *Object) HasMethod(key string) bool {
	return o.methods[key] == nil
}

func (o *Object) SetMethod(key string, value concept.Function) concept.Exception {
	o.methods[key] = value
	return nil
}

func (o *Object) GetMethod(key string) (concept.Function, concept.Exception) {
	value := o.methods[key]
	if value == nil {
		return nil, interrupt.NewException("system error", fmt.Sprintf("no method called %v", key))
	}
	return o.methods[key], nil
}

func (a *Object) ToString(prefix string) string {
	if 0 == len(a.fields) && 0 == len(a.reflections) {
		return "object <> {}"
	}

	subPrefix := fmt.Sprintf("%v\t", prefix)

	paramsToString := make([]string, 0, len(a.fields))
	for key, value := range a.fields {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key, value.ToString(subPrefix)))
	}
	for key, value := range a.methods {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key, value.ToString(subPrefix)))
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

func (o *Object) Type() string {
	return VariableObjectType
}

func NewObject() *Object {
	return &Object{
		fields:      make(map[string]concept.Variable),
		methods:     make(map[string]concept.Function),
		reflections: make([]*component.ClassReflection, 0),
	}
}
