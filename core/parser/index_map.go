package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type IndexMap struct {
	values map[tree.Phrase]map[tree.Phrase]int
}

func NewIndexMap() *IndexMap {
	return &IndexMap{
		values: make(map[tree.Phrase]map[tree.Phrase]int),
	}
}

func (m *IndexMap) Add(parentIndex int, parent tree.Phrase) {
	m.values[parent] = make(map[tree.Phrase]int)
	for index := 0; index < parent.Size(); index++ {
		child := parent.GetChild(index)
		if m.values[child] == nil {
			m.values[child] = make(map[tree.Phrase]int)
		}
		m.values[child][parent] = parentIndex
	}
}

func (m *IndexMap) Remove(parent tree.Phrase) {
	delete(m.values, parent)
	for index := 0; index < parent.Size(); index++ {
		child := parent.GetChild(index)
		if m.values[child] != nil {
			delete(m.values[child], parent)
		}
	}
}

func (m *IndexMap) Has(value tree.Phrase) bool {
	_, ok := m.values[value]
	return ok
}

func (m *IndexMap) IsChildParent(child tree.Phrase, parent tree.Phrase) bool {
	_, ok := m.values[child][parent]
	return ok
}

func (m *IndexMap) Get(child tree.Phrase) map[tree.Phrase]int {
	return m.values[child]
}
