package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	VariableMappingObjectType = "mapping_object"
)

type MappingObjectSeed interface {
	ToLanguage(string, *MappingObject) string
	Type() string
	NewException(string, string) concept.Exception
}

type MappingObject struct {
	mapping   map[concept.String]concept.String
	alias     string
	class     concept.Class
	className string
	object    concept.Object
	seed      MappingObjectSeed
}

func (f *MappingObject) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
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
		return nil, m.seed.NewException("system error", fmt.Sprintf("No mapping who's class is \"%v\" and alias is \"%v\"", className, alias))
	}
	return m.mapping, nil
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
	return m.seed.Type()
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
	return m.seed.NewException("system error", "Mapping object cannot init.")
}

func (m *MappingObject) HasField(specimen concept.String) bool {
	return !nl_interface.IsNil(m.specimenClassToObject(specimen))
}

func (m *MappingObject) HasMethod(specimen concept.String) bool {
	return m.class.HasMethod(specimen)
}

func (m *MappingObject) SetMethod(specimen concept.String, value concept.Function) concept.Exception {
	return m.seed.NewException("system error", "Mapping object cannot set method.")
}

func (m *MappingObject) GetMethod(specimen concept.String) (concept.Function, concept.Exception) {
	value := m.class.GetMethod(specimen)
	if nl_interface.IsNil(value) {
		return nil, m.seed.NewException("system error", fmt.Sprintf("no method called %v", specimen.ToString("")))
	}
	return value, nil
}

type MappingObjectCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
}

type MappingObjectCreator struct {
	Seeds map[string]func(string, *MappingObject) string
	param *MappingObjectCreatorParam
}

func (s *MappingObjectCreator) ToLanguage(language string, instance *MappingObject) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *MappingObjectCreator) Type() string {
	return VariableObjectType
}

func (s *MappingObjectCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *MappingObjectCreator) New(object concept.Object, className string, alias string) (*MappingObject, concept.Exception) {

	class := object.GetClass(className)
	if nl_interface.IsNil(class) {
		return nil, s.NewException("system error", "Class name does not exist.")
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
		seed:      s,
	}, nil
}

func NewMappingObjectCreator(param *MappingObjectCreatorParam) *MappingObjectCreator {
	return &MappingObjectCreator{
		Seeds: map[string]func(string, *MappingObject) string{},
		param: param,
	}
}
