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
	reflections []*component.ClassReflection
}

func (o *Object) GetReflection(class concept.Class) []*component.ClassReflection {
	reflections := []*component.ClassReflection{}
	for _, reflection := range o.reflections {
		if reflection.GetClass().GetName() == class.GetName() {
			reflections = append(reflections, reflection)
		}
	}
	return reflections
}

func (o *Object) AddReflection(reflection *component.ClassReflection) concept.Exception {
	for _, old := range o.reflections {
		if old.GetClass().GetName() == reflection.GetClass().GetName() && old.GetAlias() == reflection.GetAlias() {
			return interrupt.NewException("system error", "Duplicate class reflections are added.")
		}
	}
	o.reflections = append(o.reflections, reflection)
	return nil
}

func (o *Object) HasField(key string) bool {
	return o.fields[key] != nil
}

func (o *Object) InitField(key string, defaultValue concept.Variable) {
	if o.fields[key] == nil {
		o.fields[key] = defaultValue
	}
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

func (a *Object) ToString(prefix string) string {
	if 0 == len(a.fields) && 0 == len(a.reflections) {
		return "object <> {}"
	}

	subPrefix := fmt.Sprintf("%v\t", prefix)

	paramsToString := make([]string, 0, len(a.fields))
	for key, value := range a.fields {
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
		reflections: make([]*component.ClassReflection, 0),
	}
}
