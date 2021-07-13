package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"strings"
)

type NewObjectSeed interface {
	ToLanguage(string, concept.Pool, *NewObject) (string, concept.Exception)
	NewException(string, string) concept.Exception
	NewObject() concept.Object
}

type NewObject struct {
	*adaptor.ExpressionIndex
	fields []*index.KeyValueIndex
	lines  []concept.Line
	seed   NewObjectSeed
}

func (f *NewObject) SetKeyValue(keyValues []concept.Pipe, lines []concept.Line) error {
	f.lines = lines
	fieldMap := map[string]bool{}
	f.fields = []*index.KeyValueIndex{}
	for cursor, keyValuePre := range keyValues {
		keyValue, yes := index.IndexFamilyInstance.IsKeyValueIndex(keyValuePre)
		if !yes {
			return fmt.Errorf("Unsupported index type in NewObject.SetKeyValue : %v\n%v", keyValuePre.Type(), lines[cursor].ToString())
		}
		if fieldMap[keyValue.Key().Value()] {
			return fmt.Errorf("Duplicate field: '%v'\n%v", keyValue.Key().Value(), lines[cursor].ToString())
		}
		fieldMap[keyValue.Key().Value()] = true
		f.fields = append(f.fields, keyValue)
	}
	return nil
}

func (f *NewObject) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewObject) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)

	if len(a.fields) == 0 {
		return "{}"
	}
	paramsToString := make([]string, 0, len(a.fields))
	for _, field := range a.fields {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, field.Key().Value(), field.Value().ToString(subPrefix)))
	}
	return fmt.Sprintf("{%v\n%v\n%v}", prefix, strings.Join(paramsToString, ",\n"), prefix)

}

func (a *NewObject) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	object := a.seed.NewObject()
	var suspend concept.Interrupt = nil
	var value concept.Variable = nil
	for cursor, field := range a.fields {
		value, suspend = field.Value().Get(space)
		if !nl_interface.IsNil(suspend) {
			return nil, suspend
		}
		suspend = object.SetField(field.Key(), value)
		if !nl_interface.IsNil(suspend) {
			return nil, suspend.AddLine(a.lines[cursor])
		}
	}
	return object, nil
}

type NewObjectCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	ExceptionCreator       func(string, string) concept.Exception
	ObjectCreator          func() concept.Object
	NullCreator            func() concept.Null
}

type NewObjectCreator struct {
	Seeds map[string]func(concept.Pool, *NewObject) (string, concept.Exception)
	param *NewObjectCreatorParam
}

func (s *NewObjectCreator) New() *NewObject {
	back := &NewObject{
		fields: []*index.KeyValueIndex{},
		lines:  []concept.Line{},
		seed:   s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewObjectCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *NewObjectCreator) NewObject() concept.Object {
	return s.param.ObjectCreator()
}

func (s *NewObjectCreator) ToLanguage(language string, space concept.Pool, instance *NewObject) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewObjectCreator(param *NewObjectCreatorParam) *NewObjectCreator {
	return &NewObjectCreator{
		Seeds: map[string]func(concept.Pool, *NewObject) (string, concept.Exception){},
		param: param,
	}
}
