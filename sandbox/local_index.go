package sandbox

type LocalIndex struct {
	key string
}

func (s *LocalIndex) ToString(prefix string) string {
	return s.key
}

func (s *LocalIndex) Get(space *Closure) (Variable, *Exception) {
	return space.GetLocal(s.key)
}

func (s *LocalIndex) Set(space *Closure, value Variable) *Exception {
	return space.SetLocal(s.key, value)
}

func NewLocalIndex(key string) *LocalIndex {
	return &LocalIndex{
		key: key,
	}
}
