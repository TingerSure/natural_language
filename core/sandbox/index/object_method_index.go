package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ObjectMethodIndex struct {
	key    concept.String
	object concept.Index
}

const (
	IndexObjectMethodType = "ObjectMethod"
)

func (f *ObjectMethodIndex) Type() string {
	return IndexObjectMethodType
}

var (
	ObjectMethodIndexLanguageSeeds = map[string]func(string, *ObjectMethodIndex) string{}
)

func (f *ObjectMethodIndex) ToLanguage(language string) string {
	seed := ObjectMethodIndexLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (s *ObjectMethodIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ObjectMethodIndex) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", s.object.ToString(prefix), s.key.ToString(prefix))
}

func (s *ObjectMethodIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return nil, interrupt.NewException(variable.NewString("type error"), variable.NewString("There is not an effective object When system call the ObjectMethodIndex.Get"))
	}

	function, suspend := object.GetMethod(s.key)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	if nl_interface.IsNil(function) {
		return nil, interrupt.NewException(variable.NewString("runtime error"), variable.NewString(fmt.Sprintf("Object don't have a method named %s.", s.key)))
	}

	return variable.NewPreObjectFunction(function, object), nil
}

func (s *ObjectMethodIndex) Set(space concept.Closure, preFunction concept.Variable) concept.Interrupt {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return interrupt.NewException(variable.NewString("type error"), variable.NewString("There is not an effective object When system call the ObjectMethodIndex.Set"))
	}

	function, ok := variable.VariableFamilyInstance.IsFunctionHome(preFunction)
	if !ok {
		return interrupt.NewException(variable.NewString("type error"), variable.NewString("There is not an effective function When system call the ObjectMethodIndex.Set"))
	}

	return object.SetMethod(s.key, function)
}

func NewObjectMethodIndex(object concept.Index, key concept.String) *ObjectMethodIndex {
	return &ObjectMethodIndex{
		key:    key,
		object: object,
	}
}
