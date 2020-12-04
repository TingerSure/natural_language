package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
	"strings"
)

type Road struct {
	sentence  []rune
	left      *IndexList
	right     *IndexList
	rightMap  *IndexMap
	rightType *IndexType
	size      int
	types     *Types
}

func (r *Road) ToString() string {
	back := r.GetSentence() + "\n"
	for index := r.size - 1; index >= 0; index-- {
		sections := r.right.GetAll(index)
		for _, section := range sections {
			back += strings.Repeat(" ", index) + section.ToContent() + "\n"
		}
	}
	return back
}

func NewRoad(sentence string, types *Types) *Road {
	road := &Road{
		sentence: []rune(sentence),
		types:    types,
	}
	road.size = len(road.sentence)
	road.left = NewIndexList(road.size)
	road.right = NewIndexList(road.size)
	road.rightType = NewIndexType(road.size, types)
	road.rightMap = NewIndexMap()
	return road
}

func (r *Road) Iterate(onPhrase func(tree.Phrase) bool) bool {
	return r.left.Iterate(onPhrase)
}

func (r *Road) AddSection(index int, section tree.Phrase) {
	r.left.Add(index-section.ContentSize()+1, section)
	r.right.Add(index, section)
	r.rightType.Add(index, section)
	r.rightMap.Add(index, section)
}

func (r *Road) removeSection(index int, value tree.Phrase) {
	r.right.Remove(index, value)
	r.rightType.Remove(index, value)
	r.rightMap.Remove(value)
	r.left.Remove(index-value.ContentSize()+1, value)
}

func (r *Road) RemoveSection(index int, value tree.Phrase) {
	parents := r.rightMap.Get(value)
	if parents != nil && len(parents) != 0 {
		for parent, parentIndex := range parents {
			r.RemoveSection(parentIndex, parent)
		}
	}
	r.removeSection(index, value)
}

func (r *Road) GetActiveSection() []tree.Phrase {
	return r.GetSectionBySize(r.SentenceSize()-1, r.SentenceSize())
}

func (r *Road) GetSectionBySize(index int, size int) []tree.Phrase {
	return r.right.GetBySize(index, size)
}

func (r *Road) GetSectionByTypesAndSize(index int, types string, size int) []tree.Phrase {
	return r.right.GetByTypesAndSize(index, types, size)
}

func (r *Road) GetSectionByTypes(index int, types string) map[tree.Phrase]bool {
	return r.rightType.Get(index, types)
}

func (r *Road) GetLeftSectionMax(index int) tree.Phrase {
	return r.left.GetMaxBySize(index)
}

func (r *Road) GetSections(index int) []tree.Phrase {
	return r.right.GetAll(index)
}

func (r *Road) HasLeftSection(index int) bool {
	return r.left.Has(index)
}

func (r *Road) HasRightSection(index int) bool {
	return r.right.Has(index)
}

func (r *Road) DependencyCheck(value tree.Phrase) bool {
	for index := 0; index < value.Size(); index++ {
		child := value.GetChild(index)
		if !r.rightMap.Has(child.DependencyCheckValue()) {
			return false
		}
	}
	return true
}

func (r *Road) GetSentence() string {
	return string(r.sentence)
}

func (r *Road) SubSentenceFrom(from int) string {
	return string(r.sentence[from:])
}

func (r *Road) SubSentence(from, to int) string {
	return string(r.sentence[from:to])
}

func (r *Road) SentenceSize() int {
	return r.size
}
