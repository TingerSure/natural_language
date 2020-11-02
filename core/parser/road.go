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
	left           []*roadNode
	right          []*roadNode
	size           int
	rightTypeIndex *TypeIndex
	types          *Types
}

func (r *Road) ToString() string {
	back := r.GetSentence() + "\n"
	for index := r.size - 1; index >= 0; index-- {
		right := r.right[index]

		for cursor := right; cursor != nil; cursor = cursor.next {
			back += strings.Repeat(" ", index) + cursor.value.ToContent() + "\n"
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
	road.left = make([]*roadNode, road.size)
	road.right = make([]*roadNode, road.size)
	road.rightTypeIndex = NewTypeIndex(road.size, types)
	return road
}

func (r *Road) ReplaceRight(index int, from, to tree.Phrase) {
	r.RemoveRightSection(index, func(phrase tree.Phrase) bool {
		return from == phrase
	})
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
	for _, node := range r.left {
		for ; node != nil; node = node.next {
			if onPhrase(node.value) {
				return true
			}
		}
	}
	return false
}

func (r *Road) CleanSection() {
	for index := 0; index < r.size-1; index++ {
		r.left[index+1] = nil
		r.right[index] = nil
	}

	removeCondition := func(phrase tree.Phrase) bool {
		return phrase.ContentSize() != r.size
	}

	r.left[0] = r.removeSectionOnly(r.left[0], removeCondition)
	r.right[r.size-1] = r.removeSectionOnly(r.right[r.size-1], removeCondition)
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
	r.left[index] = r.addSectionOnly(r.left[index], section)
}

func (r *Road) addRightSectionOnly(index int, section tree.Phrase) {
	r.right[index] = r.addSectionOnly(r.right[index], section)
	r.rightTypeIndex.Add(index, section)
}

func (r *Road) addSectionOnly(node *roadNode, section tree.Phrase) *roadNode {
	if node == nil || node.value.ContentSize() <= section.ContentSize() {
		return &roadNode{
			value: section,
			next:  node,
		}
	}
	last := node
	cursor := node.next
	for cursor != nil && cursor.value.ContentSize() > section.ContentSize() {
		last = cursor
		cursor = cursor.next
	}
	last.next = &roadNode{
		value: section,
		next:  cursor,
	}
	return node
}

func (r *Road) RemoveLeftSection(index int, condition func(tree.Phrase) bool) {
	r.removeLeftSectionOnly(index, func(phraseLeft tree.Phrase) bool {
		if condition(phraseLeft) {
			r.removeRightSectionOnly(index+phraseLeft.ContentSize()-1, func(phraseRight tree.Phrase) bool {
				return phraseLeft == phraseRight
			})
			return true
		}
		return false
	})
}

func (r *Road) RemoveRightSection(index int, condition func(tree.Phrase) bool) {
	r.removeRightSectionOnly(index, func(phraseRight tree.Phrase) bool {
		if condition(phraseRight) {
			r.removeLeftSectionOnly(index-phraseRight.ContentSize()+1, func(phraseLeft tree.Phrase) bool {
				return phraseLeft == phraseRight
			})
			return true
		}
		return false
	})
}

func (r *Road) removeLeftSectionOnly(index int, condition func(tree.Phrase) bool) {
	r.left[index] = r.removeSectionOnly(r.left[index], condition)
}

func (r *Road) removeRightSectionOnly(index int, condition func(tree.Phrase) bool) {
	r.right[index] = r.removeSectionOnly(r.right[index], func(section tree.Phrase) bool {
		yes := condition(section)
		if yes {
			r.rightTypeIndex.Remove(index, section)
		}
		return yes
	})
}

func (r *Road) removeSectionOnly(root *roadNode, condition func(tree.Phrase) bool) *roadNode {
	cursor := root
	var last *roadNode = nil
	for cursor != nil {
		if condition(cursor.value) {
			if cursor == root {
				root = cursor.next
			} else {
				last.next = cursor.next
			}
		} else {
			last = cursor
		}
		cursor = cursor.next
	}
	return root
}

func (r *Road) GetActiveSection() []tree.Phrase {
	return r.GetLeftSection(0, func(phrase tree.Phrase) bool {
		return phrase.ContentSize() == r.SentenceSize()
	})
}

func (r *Road) GetRightSectionByTypesAndSize(index int, types string, size int) tree.Phrase {
	for cursor := r.right[index]; cursor != nil; cursor = cursor.next {
		if cursor.value.ContentSize() > size {
			continue
		}
		if cursor.value.ContentSize() < size {
			return nil
		}
		if types == cursor.value.Types() {
			return cursor.value
		}
	}
	return nil
}

func (r *Road) GetRightSectionByTypes(index int, types string) map[tree.Phrase]bool {
	return r.rightTypeIndex.Get(index, types)
}

func (r *Road) GetLeftSectionMax(index int) tree.Phrase {
	if r.left[index] == nil {
		return nil
	}
	return r.left[index].value
}

func (r *Road) GetLeftSection(index int, condition func(tree.Phrase) bool) []tree.Phrase {
	return r.getSection(r.left[index], condition)
}

func (r *Road) GetRightSection(index int, condition func(tree.Phrase) bool) []tree.Phrase {
	return r.getSection(r.right[index], condition)
}

func (r *Road) getSection(node *roadNode, condition func(tree.Phrase) bool) []tree.Phrase {
	back := []tree.Phrase{}
	for cursor := node; cursor != nil; cursor = cursor.next {
		if condition == nil || condition(cursor.value) {
			back = append(back, cursor.value)
		}
	}
	return back
}

func (r *Road) HasLeftSection(index int) bool {
	return r.left[index] != nil
}
func (r *Road) HasRightSection(index int) bool {
	return r.right[index] != nil
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
