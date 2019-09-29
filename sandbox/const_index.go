package sandbox

type ConstIndex struct {
	value Variable
}

func (s *ConstIndex) ToString(prefix string) string {
	return s.value.ToString(prefix)
}

func (s *ConstIndex) Get(space *Closure) (Variable, *Exception) {
	return s.value, nil
}

func (s *ConstIndex) Set(space *Closure, value Variable) *Exception {
	return NewException("read only", "Constants cannot be changed.")
}

func NewConstIndex(value Variable) *ConstIndex {
	return &ConstIndex{
		value: value,
	}
}
