package variable

import (
	"fmt"
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
	mapping *concept.Mapping
	class   concept.Class
	object  concept.Variable
	seed    MappingObjectSeed
}

func (o *MappingObject) IsFunction() bool {
	return false
}

func (o *MappingObject) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	if o.class.HasRequire(specimen) {
		return o.callRequire(specimen, param)
	}
	if o.class.HasProvide(specimen) {
		return o.callProvide(specimen, param)
	}
	return nil, o.seed.NewException("runtime error", fmt.Sprintf("There is no function called \"%v\" to be called here.", specimen.ToString("")))
}

func (o *MappingObject) callRequire(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	mappingSpecimenInterface := o.mapping.Get(specimen)
	mappingSpecimen, yes := VariableFamilyInstance.IsString(mappingSpecimenInterface.(concept.Variable))
	if !yes {
		return nil, o.seed.NewException("runtime error", fmt.Sprintf("There is no require function called \"%v\" in mapping.", specimen.ToString("")))
	}
	return o.object.Call(mappingSpecimen, param)
}

func (o *MappingObject) callProvide(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	return o.class.GetProvide(specimen).Exec(param, o)
}

func (f *MappingObject) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (m *MappingObject) GetSource() concept.Variable {
	return m.object
}

func (m *MappingObject) GetClass() concept.Class {
	return m.class
}

func (m *MappingObject) GetMapping() *concept.Mapping {
	return m.mapping
}

func (a *MappingObject) ToString(prefix string) string {
	return "TODO MappingObject.ToString()"
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
	return VariableMappingObjectType
}

func (s *MappingObjectCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *MappingObjectCreator) New(object concept.Variable, classInstance concept.Class, mapping *concept.Mapping) *MappingObject {
	return &MappingObject{
		mapping: mapping,
		class:   classInstance,
		object:  object,
		seed:    s,
	}
}

func NewMappingObjectCreator(param *MappingObjectCreatorParam) *MappingObjectCreator {
	return &MappingObjectCreator{
		Seeds: map[string]func(string, *MappingObject) string{},
		param: param,
	}
}
