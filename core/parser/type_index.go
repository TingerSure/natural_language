package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type TypeIndex struct {
	values []map[*tree.PhraseType]map[tree.Phrase]bool
	size   int
}

func NewTypeIndex(size int) *TypeIndex {
	return &TypeIndex{
		size:   size,
		values: make([]map[*tree.PhraseType]map[tree.Phrase]bool, size),
	}
}

func (t *TypeIndex) Get(index int, types *tree.PhraseType) map[tree.Phrase]bool {
	return t.values[index][types]
}

func (t *TypeIndex) Add(index int, section tree.Phrase) {

	if t.values[index] == nil {
		t.values[index] = make(map[*tree.PhraseType]map[tree.Phrase]bool)
	}

	types := section.Types()

	types.IterateParentsDistinct(func(parent *tree.PhraseType) bool {
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

	types.IterateParentsDistinct(func(parent *tree.PhraseType) bool {
		delete(t.values[index][parent], section)
		return false
	})
	delete(t.values[index][types], section)
}
