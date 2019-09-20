package sandbox

import (
	"errors"
)

type BubbleIndex struct {
	key string
}

func (s *BubbleIndex) Get(space Closure) (Variable, error) {
	return space.GetBubble(s.key)
}

func (s *BubbleIndex) Set(space Closure, value Variable) error {
	return space.SetBubble(s.key, value)
}

func NewBubbleIndex(key string) *BubbleIndex {
	return &BubbleIndex{
		key: key,
	}
}
