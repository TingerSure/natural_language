package sandbox

type LocalIndex struct {
	key string
}

func (s *LocalIndex) Get(space *Closure) (Variable, error) {
	return space.GetLocal(s.key)
}

func (s *LocalIndex) Set(space *Closure, value Variable) error {
	return space.SetLocal(s.key, value)
}

func NewLocalIndex(key string) *LocalIndex {
	return &LocalIndex{
		key: key,
	}
}
