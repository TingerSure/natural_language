package sandbox

import (
	"errors"
)

type CacheIndex struct {
	index int
}

func (s *CacheIndex) Get(space Closure) (Variable, error) {
	return space.GetCache(s.index)
}

func (s *CacheIndex) Set(space Closure, value Variable) error {
	return space.SetCache(s.index, value)
}

func NewCacheIndex(index int) *CacheIndex {
	return &CacheIndex{
		index: index,
	}
}
