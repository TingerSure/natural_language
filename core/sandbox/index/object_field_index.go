package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ObjectFieldIndexSeed interface {
	ToLanguage(string, *ObjectFieldIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewNull() concept.Null
}

type ObjectFieldIndex struct {
	key    concept.String
	object concept.Index
	seed   ObjectFieldIndexSeed
}

const (
	IndexObjectFieldType = "ObjectField"
)

func (f *ObjectFieldIndex) Type() string {
	return f.seed.Type()
}

func (f *ObjectFieldIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *ObjectFieldIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ObjectFieldIndex) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", s.object.ToString(prefix), s.key.ToString(prefix))
}

func (s *ObjectFieldIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	object, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return s.seed.NewNull(), suspend
	}
	return object.GetField(s.key)
}

func (s *ObjectFieldIndex) Anticipate(space concept.Closure) concept.Variable {
	object := s.object.Anticipate(space)
	value, _ := object.GetField(s.key)
	return value
}

func (s *ObjectFieldIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	object, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	return object.SetField(s.key, value)
}

type ObjectFieldIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	NullCreator      func() concept.Null
}

type ObjectFieldIndexCreator struct {
	Seeds map[string]func(string, *ObjectFieldIndex) string
	param *ObjectFieldIndexCreatorParam
}

func (s *ObjectFieldIndexCreator) New(object concept.Index, key concept.String) *ObjectFieldIndex {
	return &ObjectFieldIndex{
		key:    key,
		object: object,
		seed:   s,
	}
}
func (s *ObjectFieldIndexCreator) ToLanguage(language string, instance *ObjectFieldIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ObjectFieldIndexCreator) Type() string {
	return IndexObjectFieldType
}

func (s *ObjectFieldIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *ObjectFieldIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewObjectFieldIndexCreator(param *ObjectFieldIndexCreatorParam) *ObjectFieldIndexCreator {
	return &ObjectFieldIndexCreator{
		Seeds: map[string]func(string, *ObjectFieldIndex) string{},
		param: param,
	}
}
