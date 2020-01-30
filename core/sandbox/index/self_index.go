package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

const (
	selfIndexKey = "self"
)

type SelfIndex struct {
}

func (s *SelfIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *SelfIndex) ToString(prefix string) string {
	return selfIndexKey
}

func (s *SelfIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(selfIndexKey)
}

func (s *SelfIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException("read only", "Self cannot be changed.")

}

func NewSelfIndex() *SelfIndex {
	return &SelfIndex{}
}
