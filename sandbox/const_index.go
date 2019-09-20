package sandbox

import (
	"errors"
)

type ConstIndex struct {
	value Variable
}

func (s *ConstIndex) Get(space Closure) (Variable, error) {
	return s.value, nil
}

func (s *ConstIndex) Set(space Closure, value Variable) error {
	return errors.New("Constants cannot be changed.")
}

func NewConstIndex(value Variable) *ConstIndex {
	return &ConstIndex{
		value: value,
	}
}
