package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type TypeIndex struct {
	values []map[string]map[tree.Phrase]bool
	size   int
	types  *Types
}

func NewTypeIndex(size int, types *Types) *TypeIndex {
	return &TypeIndex{
		types:  types,
		size:   size,
		values: make([]map[string]map[tree.Phrase]bool, size),
	}
}

func (t *TypeIndex) Get(index int, types string) map[tree.Phrase]bool {
	return t.values[index][types]
}

func (t *TypeIndex) Add(index int, section tree.Phrase) {

	if t.values[index] == nil {
		t.values[index] = make(map[string]map[tree.Phrase]bool)
	}

	types := section.Types()

	t.types.IterateParentsDistinct(types, func(parent string) bool {
		if t.values[index][parent] == nil {
			t.values[index][parent] = make(map[tree.Phrase]bool)
		}
		t.values[index][parent][section] = true
		return false
	})
	if t.values[index][types] == nil {
		t.values[index][types] = make(map[tree.Phrase]bool)
	}
	t.values[index][types][section] = true
}

func (t *TypeIndex) Remove(index int, section tree.Phrase) {
	if t.values[index] == nil {
		return
	}

	types := section.Types()

	t.types.IterateParentsDistinct(types, func(parent string) bool {
		delete(t.values[index][parent], section)
		return false
	})
	delete(t.values[index][types], section)
}
