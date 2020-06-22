package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ObjectMethodIndexSeed interface {
	ToLanguage(string, *ObjectMethodIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewPreObjectFunction(concept.Function, concept.Object) *variable.PreObjectFunction
	NewNull() concept.Null
}

type ObjectMethodIndex struct {
	key    concept.String
	object concept.Index
	seed   ObjectMethodIndexSeed
}

const (
	IndexObjectMethodType = "ObjectMethod"
)

func (f *ObjectMethodIndex) Type() string {
	return f.seed.Type()
}

func (f *ObjectMethodIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *ObjectMethodIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ObjectMethodIndex) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", s.object.ToString(prefix), s.key.ToString(prefix))
}

func (s *ObjectMethodIndex) Anticipate(space concept.Closure) concept.Variable {
	preObject := s.object.Anticipate(space)

	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return s.seed.NewNull()
	}

	function, suspend := object.GetMethod(s.key)
	if !nl_interface.IsNil(suspend) {
		return s.seed.NewNull()
	}
	if nl_interface.IsNil(function) {
		return s.seed.NewNull()
	}

	return s.seed.NewPreObjectFunction(function, object)
}

func (s *ObjectMethodIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return s.seed.NewNull(), suspend
	}

	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return s.seed.NewNull(), s.seed.NewException("type error", "There is not an effective object When system call the ObjectMethodIndex.Get")
	}

	function, suspend := object.GetMethod(s.key)
	if !nl_interface.IsNil(suspend) {
		return s.seed.NewNull(), suspend
	}
	if nl_interface.IsNil(function) {
		return s.seed.NewNull(), s.seed.NewException("runtime error", fmt.Sprintf("Object don't have a method named %s.", s.key))
	}

	return s.seed.NewPreObjectFunction(function, object), nil
}

func (s *ObjectMethodIndex) Set(space concept.Closure, preFunction concept.Variable) concept.Interrupt {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return s.seed.NewException("type error", "There is not an effective object When system call the ObjectMethodIndex.Set")
	}

	function, ok := variable.VariableFamilyInstance.IsFunctionHome(preFunction)
	if !ok {
		return s.seed.NewException("type error", "There is not an effective function When system call the ObjectMethodIndex.Set")
	}

	return object.SetMethod(s.key, function)
}

type ObjectMethodIndexCreatorParam struct {
	ExceptionCreator         func(string, string) concept.Exception
	PreObjectFunctionCreator func(concept.Function, concept.Object) *variable.PreObjectFunction
	NullCreator              func() concept.Null
}

type ObjectMethodIndexCreator struct {
	Seeds map[string]func(string, *ObjectMethodIndex) string
	param *ObjectMethodIndexCreatorParam
}

func (s *ObjectMethodIndexCreator) New(object concept.Index, key concept.String) *ObjectMethodIndex {
	return &ObjectMethodIndex{
		key:    key,
		object: object,
		seed:   s,
	}
}
func (s *ObjectMethodIndexCreator) ToLanguage(language string, instance *ObjectMethodIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ObjectMethodIndexCreator) Type() string {
	return IndexObjectMethodType
}

func (s *ObjectMethodIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *ObjectMethodIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ObjectMethodIndexCreator) NewPreObjectFunction(function concept.Function, object concept.Object) *variable.PreObjectFunction {
	return s.param.PreObjectFunctionCreator(function, object)
}

func NewObjectMethodIndexCreator(param *ObjectMethodIndexCreatorParam) *ObjectMethodIndexCreator {
	return &ObjectMethodIndexCreator{
		Seeds: map[string]func(string, *ObjectMethodIndex) string{},
		param: param,
	}
}
