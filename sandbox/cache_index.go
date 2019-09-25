package sandbox

type CacheIndex struct {
	index int
}

func (s *CacheIndex) Get(space *Closure) (Variable, *Exception) {
	return space.GetCache(s.index), nil
}

func (s *CacheIndex) Set(space *Closure, value Variable) *Exception {
	space.SetCache(s.index, value)
	return nil
}

func NewCacheIndex(index int) *CacheIndex {
	return &CacheIndex{
		index: index,
	}
}
