package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type CacheIndex struct {
	index int
}

func (s *CacheIndex) ToString(prefix string) string {
	return fmt.Sprintf("cache[%v]", s.index)
}

func (s *CacheIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetCache(s.index), nil
}

func (s *CacheIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	space.SetCache(s.index, value)
	return nil
}

func NewCacheIndex(index int) *CacheIndex {
	return &CacheIndex{
		index: index,
	}
}
