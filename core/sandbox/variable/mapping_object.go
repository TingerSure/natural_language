package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

const (
	VariableMappingObjectType = "mapping_object"
)

type MappingObject struct {
	mapping   map[string]string
	alias     string
	class     concept.Class
	className string
	object    concept.Object
}

func NewMappingObject(object concept.Object, className string, alias string) (*MappingObject, concept.Exception) {

	class := object.GetClass(className)
	if nl_interface.IsNil(class) {
		return nil, interrupt.NewException("system error", "Class name does not exist.")
	}

	mapping, exception := object.GetMapping(className, alias)
	if !nl_interface.IsNil(exception) {
		return nil, exception
	}

	return &MappingObject{
		mapping:   mapping,
		alias:     alias,
		class:     class,
		className: className,
		object:    object,
	}, nil
}
func (m *MappingObject) GetSource() concept.Object {
	return m.object
}

func (m *MappingObject) GetClasses() []string {
	return []string{
		m.className,
	}
}

func (m *MappingObject) GetClass(className string) concept.Class {
	if className == m.className {
		return m.class
	}
	return nil
}

func (m *MappingObject) GetAliases(className string) []string {
	if className == m.className {
		return []string{
			m.alias,
		}
	}
	return []string{}
}

func (m *MappingObject) IsClassAlias(className string, alias string) bool {
	return className == m.className && alias == m.alias
}

func (m *MappingObject) GetMapping(className string, alias string) (map[string]string, concept.Exception) {
	if className != m.className || alias != m.alias {
		return nil, interrupt.NewException("system error", fmt.Sprintf("No mapping who's class is \"%v\" and alias is \"%v\"", className, alias))
	}
	var mapping map[string]string
	for key, _ := range m.mapping {
		mapping[key] = key
	}
	return mapping, nil
}

func (a *MappingObject) ToString(prefix string) string {
	if a.alias == "" {
		return fmt.Sprintf("%v<%v>", a.object.ToString(prefix), a.className)
	}
	return fmt.Sprintf("%v<%v(%v)>", a.object.ToString(prefix), a.className, a.alias)
}

func (m *MappingObject) Type() string {
	return VariableMappingObjectType
}

func (m *MappingObject) SetField(key string, value concept.Variable) concept.Exception {
	return m.object.SetField(m.mapping[key], value)
}

func (m *MappingObject) GetField(key string) (concept.Variable, concept.Exception) {
	return m.object.GetField(m.mapping[key])
}

func (m *MappingObject) InitField(string, concept.Variable) concept.Exception {
	return interrupt.NewException("system error", "Mapping object cannot init.")
}

func (m *MappingObject) HasField(key string) bool {
	return m.mapping[key] != ""
}

func (m *MappingObject) HasMethod(key string) bool {
	return m.class.HasMethod(key)
}

func (m *MappingObject) SetMethod(key string, value concept.Function) concept.Exception {
	return interrupt.NewException("system error", "Mapping object cannot set method.")

}

func (m *MappingObject) GetMethod(key string) (concept.Function, concept.Exception) {
	value := m.class.GetMethod(key)
	if value == nil {
		return nil, interrupt.NewException("system error", fmt.Sprintf("no method called %v", key))
	}
	return value, nil
}
