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
	mapping   map[concept.String]concept.String
	alias     string
	class     concept.Class
	className string
	object    concept.Object
}

var (
	MappingObjectLanguageSeeds = map[string]func(string, *MappingObject) string{}
)

func (f *MappingObject) ToLanguage(language string) string {
	seed := MappingObjectLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func NewMappingObject(object concept.Object, className string, alias string) (*MappingObject, concept.Exception) {

	class := object.GetClass(className)
	if nl_interface.IsNil(class) {
		return nil, interrupt.NewException(NewString("system error"), NewString("Class name does not exist."))
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

func (m *MappingObject) GetMapping(className string, alias string) (map[concept.String]concept.String, concept.Exception) {
	if className != m.className || alias != m.alias {
		return nil, interrupt.NewException(NewString("system error"), NewString(fmt.Sprintf("No mapping who's class is \"%v\" and alias is \"%v\"", className, alias)))
	}
	var mapping map[concept.String]concept.String
	for key, _ := range m.mapping {
		mapping[key] = key
	}
	return mapping, nil
}

func (m *MappingObject) CheckMapping(concept.Class, map[concept.String]concept.String) bool {
	return false
	// TODO
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

func (m *MappingObject) specimenClassToObject(specimen concept.String) concept.String {
	var objectSpecimen concept.String = nil
	m.class.IterateFields(func(key concept.String, _ concept.Variable) bool {
		if key.EqualLanguage(specimen) {
			for target, source := range m.mapping {
				if key.EqualLanguage(target) {
					objectSpecimen = source
					break
				}
			}
			return true
		}
		return false
	})
	return objectSpecimen
}

func (m *MappingObject) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	return m.object.SetField(m.specimenClassToObject(specimen), value)
}

func (m *MappingObject) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	return m.object.GetField(m.specimenClassToObject(specimen))
}

func (m *MappingObject) InitField(concept.String, concept.Variable) concept.Exception {
	return interrupt.NewException(NewString("system error"), NewString("Mapping object cannot init."))
}

func (m *MappingObject) HasField(specimen concept.String) bool {
	return !nl_interface.IsNil(m.specimenClassToObject(specimen))
}

func (m *MappingObject) HasMethod(specimen concept.String) bool {
	return m.class.HasMethod(specimen)
}

func (m *MappingObject) SetMethod(specimen concept.String, value concept.Function) concept.Exception {
	return interrupt.NewException(NewString("system error"), NewString("Mapping object cannot set method."))
}

func (m *MappingObject) GetMethod(specimen concept.String) (concept.Function, concept.Exception) {
	value := m.class.GetMethod(specimen)
	if nl_interface.IsNil(value) {
		return nil, interrupt.NewException(NewString("system error"), NewString(fmt.Sprintf("no method called %v", specimen.ToString(""))))
	}
	return value, nil
}
