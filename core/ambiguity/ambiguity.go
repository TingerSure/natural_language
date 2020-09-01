package ambiguity

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/tree"
	"sort"
)

type sortPhrase struct {
	phrases []tree.Phrase
}

func (s *sortPhrase) Len() int {
	return len(s.phrases)
}

func (s *sortPhrase) Swap(left, right int) {
	s.phrases[left], s.phrases[right] = s.phrases[right], s.phrases[left]
}

func (s *sortPhrase) Compare(left, right tree.Phrase) int {
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

func (s *sortPhrase) Less(left, right int) bool {
	return s.Compare(s.phrases[left], s.phrases[right]) < 0
}

type Ambiguity struct {
	rules []*tree.PriorityRule
}

func (a *Ambiguity) Sort(phrases []tree.Phrase) {
	sort.Sort(&sortPhrase{
		phrases: phrases,
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

func (a *Ambiguity) Filter(phrases []tree.Phrase) []tree.Phrase {
	a.Sort(phrases)
	var obsolete []bool = make([]bool, len(phrases), len(phrases))
	result := []tree.Phrase{}
	for leftIndex, left := range phrases {
		if obsolete[leftIndex] {
			continue
		}
	rightLoop:
		for rightIndex := leftIndex + 1; rightIndex < len(phrases); rightIndex++ {
			if obsolete[rightIndex] {
				continue
			}
			right := phrases[rightIndex]
			diffLeft, diffRight := a.Diff(left, right)
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
			result = append(result, phrases[leftIndex])
		}
	}
	return result
}

func NewAmbiguity() *Ambiguity {
	return &Ambiguity{}
}
