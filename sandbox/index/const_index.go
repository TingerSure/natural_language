package index

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
)

type ConstIndex struct {
	value concept.Variable
}

func (s *ConstIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ConstIndex) ToString(prefix string) string {
	return s.value.ToString(prefix)
}

func (s *ConstIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return s.value, nil
}

func (s *ConstIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return interrupt.NewException("read only", "Constants cannot be changed.")
}

func NewConstIndex(value concept.Variable) *ConstIndex {
	return &ConstIndex{
		value: value,
	}
}
