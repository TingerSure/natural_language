package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/tree"
)

type roadNode struct {
	next  *roadNode
	value tree.Phrase
}

type Road struct {
	sentence []rune
	left     []*roadNode
	right    []*roadNode
}

func NewRoad(sentence string) *Road {
	road := &Road{
		sentence: []rune(sentence),
	}
	road.left = make([]*roadNode, road.SentenceSize())
	road.right = make([]*roadNode, road.SentenceSize())
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
	for index := 0; index < r.SentenceSize()-1; index++ {
		r.left[index+1] = nil
		r.right[index] = nil
	}

	removeCondition := func(phrase tree.Phrase) bool {
		return phrase.ContentSize() != r.SentenceSize()
	}

	r.left[0] = r.removeSectionOnly(r.left[0], removeCondition)
	r.right[r.SentenceSize()-1] = r.removeSectionOnly(r.right[r.SentenceSize()-1], removeCondition)
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
	r.right[index] = r.removeSectionOnly(r.right[index], condition)
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

func (r *Road) GetRightSectionByTypesAndSize(index int, types *tree.PhraseType, size int) tree.Phrase {
	results := r.GetRightSection(index, func(phrase tree.Phrase) bool {
		return size == phrase.ContentSize() && types.Equal(phrase.Types())
	})
	if len(results) == 0 {
		return nil
	}

	if len(results) > 1 {
		panic(fmt.Sprintf("Too much right section when Types = %v and Size = %v", types.Name(), size))
	}

	return results[0]
}

func (r *Road) GetRightSectionByTypes(index int, types *tree.PhraseType) []tree.Phrase {
	return r.GetRightSection(index, func(phrase tree.Phrase) bool {
		return types.Match(phrase.Types())
	})
	//TODO reImplement
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
	return len(r.sentence)
}
