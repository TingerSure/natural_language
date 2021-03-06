package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type IndexType struct {
	values []map[string]map[tree.Phrase]bool
	size   int
	types  *Types
}

func NewIndexType(size int, types *Types) *IndexType {
	return &IndexType{
		types:  types,
		size:   size,
		values: make([]map[string]map[tree.Phrase]bool, size),
	}
}

func (t *IndexType) Get(index int, types string) map[tree.Phrase]bool {
	return t.values[index][types]
}

func (t *IndexType) Add(index int, section tree.Phrase) error {

	if t.values[index] == nil {
		t.values[index] = make(map[string]map[tree.Phrase]bool)
	}

	types, err := section.Types()
	if err != nil {
		return err
	}

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
	return nil
}

func (t *IndexType) Remove(index int, section tree.Phrase) error {
	if t.values[index] == nil {
		return nil
	}

	types, err := section.Types()
	if err != nil {
		return err
	}
	t.types.IterateParentsDistinct(types, func(parent string) bool {
		delete(t.values[index][parent], section)
		return false
	})
	delete(t.values[index][types], section)
	return nil
}
