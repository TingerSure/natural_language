package ambiguity

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/lexer"
	"github.com/TingerSure/natural_language/tree"
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

func (a *Ambiguity) Filter(flow *lexer.Flow, phrases []tree.Phrase) tree.Phrase {
	a.Sort(phrases)
	var base tree.Phrase = nil
	for _, now := range phrases {
		if nl_interface.IsNil(base) {
			base = now
			continue
		}
		diffBase, diffNow := a.Diff(base, now)
		if nl_interface.IsNil(diffBase) && nl_interface.IsNil(diffNow) {
			continue
		}
		switch a.Check(diffBase, diffNow) {
		case 0, -1:
		case 1:
			base = now
		}
	}
	return base
}

func NewAmbiguity() *Ambiguity {
	return &Ambiguity{}
}
