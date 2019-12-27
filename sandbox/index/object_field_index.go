package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type ObjectFieldIndex struct {
	key    string
	object concept.Index
}

func (s *ObjectFieldIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ObjectFieldIndex) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", s.object.ToString(prefix), s.key)
}

func (s *ObjectFieldIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return nil, interrupt.NewException("type error", "There is not an effective object When system call the ObjectFieldIndex.Get")
	}
	return object.GetField(s.key)
}

func (s *ObjectFieldIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return interrupt.NewException("type error", "There is not an effective object When system call the ObjectFieldIndex.Set")
	}
	return object.SetField(s.key, value)
}

func NewObjectFieldIndex(object concept.Index, key string) *ObjectFieldIndex {
	return &ObjectFieldIndex{
		key:    key,
		object: object,
	}
}
