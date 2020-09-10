package grammar

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/tree"
	"sort"
)

type sortRiver struct {
	rivers []*River
}

func (s *sortRiver) Len() int {
	return len(s.rivers)
}

func (s *sortRiver) Swap(left, right int) {
	s.rivers[left], s.rivers[right] = s.rivers[right], s.rivers[left]
}

func (s *sortRiver) Compare(left, right tree.Phrase) int {
	if left.From() < right.From() {
		return -1
	}
	if left.From() > right.From() {
		return 1
	}
	for index := 0; index < left.Size(); index++ {
		subLeft := left.GetChild(index)
		subRight := right.GetChild(index)

		result := s.Compare(subLeft, subRight)
		if result != 0 {
			return result
		}
	}
	return 0
}

func (s *sortRiver) Less(left, right int) bool {
	return s.Compare(s.rivers[left].GetWait().Peek(), s.rivers[right].GetWait().Peek()) < 0
}

type Ambiguity struct {
	rules []*tree.PriorityRule
}

func (a *Ambiguity) Sort(rivers []*River) {
	sort.Sort(&sortRiver{
		rivers: rivers,
	})
}

func (a *Ambiguity) Diff(left tree.Phrase, right tree.Phrase) (tree.Phrase, tree.Phrase) {
	if left.From() != right.From() {
		return left, right
	}

	var diffLeft, diffRight tree.Phrase = nil, nil
	var count int = 0
	for index := 0; index < left.Size(); index++ {
		subLeft, subRight := left.GetChild(index), right.GetChild(index)
		diffSubLeft, diffSubRight := a.Diff(subLeft, subRight)
		if !nl_interface.IsNil(diffSubLeft) && !nl_interface.IsNil(diffSubRight) {
			if count != 0 {
				return left, right
			}
			count++
			diffLeft, diffRight = diffSubLeft, diffSubRight
		}
	}
	return diffLeft, diffRight
}

func (a *Ambiguity) AddRule(rules []*tree.PriorityRule) {
	if rules == nil {
		return
	}
	a.rules = append(a.rules, rules...)
}

func (a *Ambiguity) Check(left, right tree.Phrase) int {
	for _, rule := range a.rules {
		if rule.Match(left, right) {
			return rule.Choose(left, right)
		}
	}
	return 0
}

func (a *Ambiguity) Filter(rivers []*River) []*River {
	a.Sort(rivers)
	var obsolete []bool = make([]bool, len(rivers), len(rivers))
	result := []*River{}
	for leftIndex, left := range rivers {
		if obsolete[leftIndex] {
			continue
		}
	rightLoop:
		for rightIndex := leftIndex + 1; rightIndex < len(rivers); rightIndex++ {
			if obsolete[rightIndex] {
				continue
			}
			right := rivers[rightIndex]
			diffLeft, diffRight := a.Diff(left.GetWait().Peek(), right.GetWait().Peek())
			if nl_interface.IsNil(diffLeft) && nl_interface.IsNil(diffRight) {
				continue
			}
			switch a.Check(diffLeft, diffRight) {
			case 0:
				continue rightLoop
			case -1:
				obsolete[rightIndex] = true
				continue rightLoop
			case 1:
				obsolete[leftIndex] = true
				break rightLoop
			}
		}
		if !obsolete[leftIndex] {
			result = append(result, rivers[leftIndex])
		}
	}
	return result
}

func NewAmbiguity() *Ambiguity {
	return &Ambiguity{}
}

func NewAmbiguityWithRules(rules []*tree.PriorityRule) *Ambiguity {
	return &Ambiguity{
		rules: rules,
	}
}
