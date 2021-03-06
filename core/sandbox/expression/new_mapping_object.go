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
	ToLanguage(string, concept.Pool, *NewMappingObject) (string, concept.Exception)
	NewException(string, string) concept.Exception
	NewMappingObject(concept.Variable, concept.Class) *variable.MappingObject
}

type NewMappingObject struct {
	*adaptor.ExpressionIndex
	object       concept.Pipe
	class        concept.Pipe
	mapping      []concept.Pipe
	line         concept.Line
	mappingLines []concept.Line
	seed         NewMappingObjectSeed
}

func (f *NewMappingObject) SetLine(line concept.Line) {
	f.line = line
}

func (f *NewMappingObject) SetMapping(mapping []concept.Pipe, mappingLines []concept.Line) {
	f.mapping = mapping
	f.mappingLines = mappingLines
}

func (f *NewMappingObject) SetObject(object concept.Pipe) {
	f.object = object
}

func (f *NewMappingObject) SetClass(class concept.Pipe) {
	f.class = class
}

func (f *NewMappingObject) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
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
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported variable type as class in NewMappingObject: %v", classPre.Type())).AddExceptionLine(a.line)
	}
	mappingObject := a.seed.NewMappingObject(object, class)
	for cursor, keykeyPre := range a.mapping {
		keykey, yes := index.IndexFamilyInstance.IsKeyKeyIndex(keykeyPre)
		if !yes {
			return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported index type in NewMappingObject : %v", keykeyPre.Type())).AddExceptionLine(a.mappingLines[cursor])
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
	Seeds map[string]func(concept.Pool, *NewMappingObject) (string, concept.Exception)
	param *NewMappingObjectCreatorParam
}

func (s *NewMappingObjectCreator) New() *NewMappingObject {
	back := &NewMappingObject{
		mappingLines: []concept.Line{},
		seed:         s,
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

func (s *NewMappingObjectCreator) ToLanguage(language string, space concept.Pool, instance *NewMappingObject) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewMappingObjectCreator(param *NewMappingObjectCreatorParam) *NewMappingObjectCreator {
	return &NewMappingObjectCreator{
		Seeds: map[string]func(concept.Pool, *NewMappingObject) (string, concept.Exception){},
		param: param,
	}
}
