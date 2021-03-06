package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableMappingObjectType = "mapping_object"
)

type MappingObjectSeed interface {
	ToLanguage(string, concept.Pool, *MappingObject) (string, concept.Exception)
	Type() string
	NewNull() concept.Null
	NewException(string, string) concept.Exception
}

type MappingObject struct {
	mapping *nl_interface.Mapping
	class   concept.Class
	object  concept.Variable
	seed    MappingObjectSeed
}

func (o *MappingObject) IsFunction() bool {
	return false
}

func (o *MappingObject) IsNull() bool {
	return false
}

func (o *MappingObject) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	if o.class.HasRequire(specimen) {
		return o.callRequire(specimen, param)
	}
	if o.class.HasProvide(specimen) {
		return o.callProvide(specimen, param)
	}
	return nil, o.seed.NewException("runtime error", fmt.Sprintf("There is no function called '%v' to be called here.", specimen.Value()))
}

func (o *MappingObject) callRequire(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	mappingSpecimen := o.mapping.Get(specimen).(concept.Variable)
	if mappingSpecimen.IsNull() {
		return o.object.Call(specimen, param)
	}
	return o.object.Call(mappingSpecimen.(concept.String), param)
}

func (o *MappingObject) callProvide(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.class.GetProvide(specimen).Exec(param, o)
}

func (f *MappingObject) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (m *MappingObject) GetSource() concept.Variable {
	return m.object
}

func (m *MappingObject) GetClass() concept.Class {
	return m.class
}

func (m *MappingObject) GetMapping() *nl_interface.Mapping {
	return m.mapping
}

func (m *MappingObject) SetMapping(from, to concept.String) {
	m.mapping.Set(from, to)
}

func (a *MappingObject) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)
	if a.mapping.Size() == 0 {
		return fmt.Sprintf("%v -> %v", a.object.ToString(prefix), a.class.ToString(prefix))
	}
	keykey := []string{}
	a.mapping.Iterate(func(key nl_interface.Key, value interface{}) bool {
		keykey = append(keykey, fmt.Sprintf("%v%v : %v", subprefix, key.(concept.String).Value(), value.(concept.String).Value()))
		return false
	})
	return fmt.Sprintf("%v -> %v {\n%v\n%v}",
		a.object.ToString(prefix),
		a.class.ToString(prefix),
		strings.Join(keykey, ",\n"),
		prefix,
	)
}

func (m *MappingObject) Type() string {
	return m.seed.Type()
}

func (m *MappingObject) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	return m.seed.NewException("system error", "Mapping object cannot set field.")
}

func (m *MappingObject) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	return nil, m.seed.NewException("system error", "Mapping object cannot get field.")
}

func (o *MappingObject) KeyField(specimen concept.String) concept.String {
	if o.class.HasRequire(specimen) || o.class.HasProvide(specimen) {
		if !o.mapping.Has(specimen) {
			o.mapping.Set(specimen, specimen)
		}
		return o.mapping.Key(specimen).(concept.String)
	}
	return specimen
}

func (o *MappingObject) SizeField() int {
	return 0
}

func (o *MappingObject) Iterate(on func(concept.String, concept.Variable) bool) bool {
	return false
}

func (m *MappingObject) HasField(specimen concept.String) bool {
	return false
}

type MappingObjectCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type MappingObjectCreator struct {
	Seeds map[string]func(concept.Pool, *MappingObject) (string, concept.Exception)
	param *MappingObjectCreatorParam
}

func (s *MappingObjectCreator) ToLanguage(language string, space concept.Pool, instance *MappingObject) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *MappingObjectCreator) Type() string {
	return VariableMappingObjectType
}

func (s *MappingObjectCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *MappingObjectCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *MappingObjectCreator) New(object concept.Variable, classInstance concept.Class) *MappingObject {
	return &MappingObject{
		mapping: nl_interface.NewMapping(&nl_interface.MappingParam{
			AutoInit:   true,
			EmptyValue: s.param.NullCreator(),
		}),
		class:  classInstance,
		object: object,
		seed:   s,
	}
}

func NewMappingObjectCreator(param *MappingObjectCreatorParam) *MappingObjectCreator {
	return &MappingObjectCreator{
		Seeds: map[string]func(concept.Pool, *MappingObject) (string, concept.Exception){},
		param: param,
	}
}
