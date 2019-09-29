package sandbox

type BubbleIndex struct {
	key string
}

func (s *BubbleIndex) ToString(prefix string) string {
	return s.key
}

func (s *BubbleIndex) Get(space *Closure) (Variable, *Exception) {
	return space.GetBubble(s.key)
}

func (s *BubbleIndex) Set(space *Closure, value Variable) *Exception {
	return space.SetBubble(s.key, value)
}

func NewBubbleIndex(key string) *BubbleIndex {
	return &BubbleIndex{
		key: key,
	}
}
