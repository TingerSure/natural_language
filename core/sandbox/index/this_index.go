package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

const (
	thisIndexKey = "this"
)

type ThisIndex struct {
}

func (s *ThisIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ThisIndex) ToString(prefix string) string {
	return thisIndexKey
}

func (s *ThisIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(thisIndexKey)
}

func (s *ThisIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException("read only", "This cannot be changed.")

}

func NewThisIndex() *ThisIndex {
	return &ThisIndex{}
}
