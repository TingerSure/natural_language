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
	ToLanguage(string, *NewObject) string
	NewObject() concept.Object
}

type NewObject struct {
	*adaptor.ExpressionIndex
	fields *concept.Mapping
	seed   NewObjectSeed
}

func (f *NewObject) SetKeyValue(keyValues []concept.Index) {
	for _, keyValuePre := range keyValues {
		keyValue, yes := index.IndexFamilyInstance.IsKeyValueIndex(keyValuePre)
		if !yes {
			panic(fmt.Sprintf("Unsupported index type in NewObject.SetKeyValue : %v", keyValuePre.Type()))
		}
		f.fields.Set(keyValue.Key(), keyValue.Value())
	}
}

func (f *NewObject) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewObject) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)

	if a.fields.Size() == 0 {
		return "{}"
	}
	paramsToString := make([]string, 0, a.fields.Size())
	a.fields.Iterate(func(key concept.String, value interface{}) bool {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.Value(), value.(concept.ToString).ToString(subPrefix)))
		return false
	})
	return fmt.Sprintf("{%v\n%v\n%v}", prefix, strings.Join(paramsToString, ",\n"), prefix)

}

func (a *NewObject) Anticipate(space concept.Closure) concept.Variable {
	object := a.seed.NewObject()
	a.fields.Iterate(func(key concept.String, value interface{}) bool {
		return !nl_interface.IsNil(object.SetField(key, value.(concept.Index).Anticipate(space)))
	})
	return object
}

func (a *NewObject) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	object := a.seed.NewObject()
	var suspend concept.Interrupt = nil
	var value concept.Variable = nil
	if a.fields.Iterate(func(key concept.String, item interface{}) bool {
		value, suspend = item.(concept.Index).Get(space)
		if !nl_interface.IsNil(suspend) {
			return true
		}
		suspend = object.SetField(key, value)
		return !nl_interface.IsNil(suspend)

	}) {
		return nil, suspend
	}
	return object, nil
}

type NewObjectCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	ObjectCreator          func() concept.Object
	NullCreator            func() concept.Null
}

type NewObjectCreator struct {
	Seeds map[string]func(string, *NewObject) string
	param *NewObjectCreatorParam
}

func (s *NewObjectCreator) New() *NewObject {
	back := &NewObject{
		seed: s,
		fields: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.param.NullCreator(),
		}),
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewObjectCreator) NewObject() concept.Object {
	return s.param.ObjectCreator()
}

func (s *NewObjectCreator) ToLanguage(language string, instance *NewObject) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewObjectCreator(param *NewObjectCreatorParam) *NewObjectCreator {
	return &NewObjectCreator{
		Seeds: map[string]func(string, *NewObject) string{},
		param: param,
	}
}
