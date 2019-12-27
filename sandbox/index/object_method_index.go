package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type ObjectMethodIndex struct {
	key    string
	object concept.Index
}

func (s *ObjectMethodIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ObjectMethodIndex) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", s.object.ToString(prefix), s.key)
}

func (s *ObjectMethodIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return nil, interrupt.NewException("type error", "There is not an effective object When system call the ObjectMethodIndex.Get")
	}
	return object.GetMethod(s.key)
}

func (s *ObjectMethodIndex) Set(space concept.Closure, preFunction concept.Variable) concept.Interrupt {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return interrupt.NewException("type error", "There is not an effective object When system call the ObjectMethodIndex.Set")
	}

	function, ok := variable.VariableFamilyInstance.IsFunctionHome(preFunction)
	if !ok {
		return interrupt.NewException("type error", "There is not an effective function When system call the ObjectMethodIndex.Set")
	}

	return object.SetMethod(s.key, function)
}

func NewObjectMethodIndex(object concept.Index, key string) *ObjectMethodIndex {
	return &ObjectMethodIndex{
		key:    key,
		object: object,
	}
}
