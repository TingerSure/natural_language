package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type BubbleIndex struct {
	key string
}

func (s *BubbleIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *BubbleIndex) ToString(prefix string) string {
	return s.key
}

func (s *BubbleIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(s.key)
}

func (s *BubbleIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return space.SetBubble(s.key, value)
}

func NewBubbleIndex(key string) *BubbleIndex {
	return &BubbleIndex{
		key: key,
	}
}
