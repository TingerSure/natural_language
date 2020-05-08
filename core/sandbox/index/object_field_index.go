package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ObjectFieldIndex struct {
	key    concept.String
	object concept.Index
}

var (
	ObjectFieldIndexLanguageSeeds = map[string]func(string, *ObjectFieldIndex) string{}
)

func (f *ObjectFieldIndex) ToLanguage(language string) string {
	seed := ObjectFieldIndexLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (s *ObjectFieldIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ObjectFieldIndex) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", s.object.ToString(prefix), s.key.ToString(prefix))
}

func (s *ObjectFieldIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preObject, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return nil, interrupt.NewException(variable.NewString("type error"), variable.NewString("There is not an effective object When system call the ObjectFieldIndex.Get"))
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
		return interrupt.NewException(variable.NewString("type error"), variable.NewString("There is not an effective object When system call the ObjectFieldIndex.Set"))
	}
	return object.SetField(s.key, value)
}

func NewObjectFieldIndex(object concept.Index, key concept.String) *ObjectFieldIndex {
	return &ObjectFieldIndex{
		key:    key,
		object: object,
	}
}
