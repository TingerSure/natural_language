package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type NewMappingObjectSeed interface {
	ToLanguage(string, concept.Pool, *NewMappingObject) string
	NewException(string, string) concept.Exception
	NewMappingObject(concept.Variable, concept.Class) *variable.MappingObject
}

type NewMappingObject struct {
	*adaptor.ExpressionIndex
	object  concept.Pipe
	class   concept.Pipe
	mapping []concept.Pipe
	seed    NewMappingObjectSeed
}

func (f *NewMappingObject) SetMapping(mapping []concept.Pipe) {
	f.mapping = mapping
}

func (f *NewMappingObject) SetObject(object concept.Pipe) {
	f.object = object
}

func (f *NewMappingObject) SetClass(class concept.Pipe) {
	f.class = class
}

func (f *NewMappingObject) ToLanguage(language string, space concept.Pool) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewMappingObject) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	if len(a.mapping) == 0 {
		return fmt.Sprintf("%v -> %v", a.object.ToString(prefix), a.class.ToString(prefix))
	}
	mapping := []string{}
	for _, keykey := range a.mapping {
		mapping = append(mapping, fmt.Sprintf("%v%v", subPrefix, keykey.ToString(subPrefix)))
	}
	return fmt.Sprintf("%v -> %v {\n%v\n%v}",
		a.object.ToString(prefix),
		a.class.ToString(prefix),
		strings.Join(mapping, ",\n"),
		prefix,
	)
}

func (a *NewMappingObject) Anticipate(space concept.Pool) concept.Variable {
	object := a.object.Anticipate(space)
	class, _ := variable.VariableFamilyInstance.IsClass(a.class.Anticipate(space))
	mappingObject := a.seed.NewMappingObject(object, class)
	for _, keykeyPre := range a.mapping {
		keykey, _ := index.IndexFamilyInstance.IsKeyKeyIndex(keykeyPre)
		mappingObject.SetMapping(keykey.From(), keykey.To())
	}
	return mappingObject
}

func (a *NewMappingObject) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	classPre, suspend := a.class.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	class, yes := variable.VariableFamilyInstance.IsClass(classPre)
	if !yes {
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported variable type as class in NewMappingObject: %v", classPre.Type()))
	}
	mappingObject := a.seed.NewMappingObject(object, class)
	for _, keykeyPre := range a.mapping {
		keykey, yes := index.IndexFamilyInstance.IsKeyKeyIndex(keykeyPre)
		if !yes {
			return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported index type in NewMappingObject : %v", keykeyPre.Type()))
		}
		mappingObject.SetMapping(keykey.From(), keykey.To())
	}
	return mappingObject, nil
}

type NewMappingObjectCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	ExceptionCreator       func(string, string) concept.Exception
	MappingObjectCreator   func(concept.Variable, concept.Class) *variable.MappingObject
}

type NewMappingObjectCreator struct {
	Seeds map[string]func(string, concept.Pool, *NewMappingObject) string
	param *NewMappingObjectCreatorParam
}

func (s *NewMappingObjectCreator) New() *NewMappingObject {
	back := &NewMappingObject{
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewMappingObjectCreator) NewMappingObject(object concept.Variable, classInstance concept.Class) *variable.MappingObject {
	return s.param.MappingObjectCreator(object, classInstance)
}

func (s *NewMappingObjectCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *NewMappingObjectCreator) ToLanguage(language string, space concept.Pool, instance *NewMappingObject) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewNewMappingObjectCreator(param *NewMappingObjectCreatorParam) *NewMappingObjectCreator {
	return &NewMappingObjectCreator{
		Seeds: map[string]func(string, concept.Pool, *NewMappingObject) string{},
		param: param,
	}
}
