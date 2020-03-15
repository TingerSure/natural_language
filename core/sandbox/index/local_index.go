package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type LocalIndex struct {
	key concept.String
}

func (s *LocalIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *LocalIndex) ToString(prefix string) string {
	return s.key.ToString(prefix)
}

func (s *LocalIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetLocal(s.key)
}

func (s *LocalIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return space.SetLocal(s.key, value)
}

func NewLocalIndex(key concept.String) *LocalIndex {
	return &LocalIndex{
		key: key,
	}
}
