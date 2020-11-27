package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
	"strings"
)

type roadNode struct {
	next  *roadNode
	value tree.Phrase
}

type Road struct {
	sentence       []rune
	left           *IndexList
	right          *IndexList
	size           int
	rightIndexType *IndexType
	types          *Types
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
	road.rightIndexType = NewIndexType(road.size, types)
	return road
}

func (r *Road) ReplaceRight(index int, from, to tree.Phrase) {
	r.RemoveRightSection(index, from)
	r.AddRightSection(index, to)

	replace := func(phrase tree.Phrase) {
		if phrase == nil {
			return
		}
		for index := 0; index < phrase.Size(); index++ {
			if phrase.GetChild(index) == from {
				phrase.SetChild(index, to)
			}
		}
	}

	r.Iterate(func(phrase tree.Phrase) bool {
		if priority, ok := phrase.(*tree.PhrasePriority); ok {
			for index := 0; index < priority.ValueSize(); index++ {
				replace(priority)
			}
			return false
		}
		replace(phrase)
		return false
	})
}

func (r *Road) Iterate(onPhrase func(tree.Phrase) bool) bool {
	return r.left.Iterate(onPhrase)
}

func (r *Road) AddLeftSection(index int, section tree.Phrase) {
	r.addLeftSectionOnly(index, section)
	r.addRightSectionOnly(index+section.ContentSize()-1, section)
}

func (r *Road) AddRightSection(index int, section tree.Phrase) {
	r.addLeftSectionOnly(index-section.ContentSize()+1, section)
	r.addRightSectionOnly(index, section)
}

func (r *Road) addLeftSectionOnly(index int, section tree.Phrase) {
	r.left.Add(index, section)
}

func (r *Road) addRightSectionOnly(index int, section tree.Phrase) {
	r.right.Add(index, section)
	r.rightIndexType.Add(index, section)
}

func (r *Road) RemoveLeftSection(index int, value tree.Phrase) {
	r.removeLeftSectionOnly(index, value)
	r.removeRightSectionOnly(index+value.ContentSize()-1, value)
}

func (r *Road) RemoveRightSection(index int, value tree.Phrase) {
	r.removeRightSectionOnly(index, value)
	r.removeLeftSectionOnly(index-value.ContentSize()+1, value)
}

func (r *Road) removeLeftSectionOnly(index int, value tree.Phrase) {
	r.left.Remove(index, value)
}

func (r *Road) removeRightSectionOnly(index int, value tree.Phrase) {
	r.right.Remove(index, value)
	r.rightIndexType.Remove(index, value)
}

func (r *Road) GetActiveSection() []tree.Phrase {
	return r.GetRightSectionBySize(r.SentenceSize()-1, r.SentenceSize())
}

func (r *Road) GetRightSectionBySize(index int, size int) []tree.Phrase {
	return r.right.GetBySize(index, size)
}

func (r *Road) GetRightSectionByTypesAndSize(index int, types string, size int) tree.Phrase {
	return r.right.GetByTypesAndSize(index, types, size)
}

func (r *Road) GetRightSectionByTypes(index int, types string) map[tree.Phrase]bool {
	return r.rightIndexType.Get(index, types)
}

func (r *Road) GetLeftSectionMax(index int) tree.Phrase {
	return r.left.GetMaxBySize(index)
}

func (r *Road) GetRightSections(index int) []tree.Phrase {
	return r.right.GetAll(index)
}

func (r *Road) HasLeftSection(index int) bool {
	return r.left.Has(index)
}
func (r *Road) HasRightSection(index int) bool {
	return r.right.Has(index)
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
