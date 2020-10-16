package parser

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/tree"
)

type Barricade struct {
	rules []*tree.PriorityRule
}

func NewBarricade() *Barricade {
	return &Barricade{}
}

func (a *Barricade) Diff(left tree.Phrase, right tree.Phrase) (tree.Phrase, tree.Phrase) {
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

func (a *Barricade) Check(left, right tree.Phrase) (int, *tree.AbandonGroup) {
	for _, rule := range a.rules {
		if rule.Match(left, right) {
			return rule.Choose(left, right)
		}
	}
	return 0, nil
}

func (a *Barricade) TargetFilter(phrases []tree.Phrase, target tree.Phrase) ([]tree.Phrase, *tree.AbandonGroup) {
	result := []tree.Phrase{}
	abandons := tree.NewAbandonGroup()
	obsolete := false
leftLoop:
	for _, left := range phrases {
		choose, abandon := a.Check(left, target)
		switch choose {
		case 0:
			result = append(result, left)
			continue leftLoop
		case -1:
			result = append(result, left)
			obsolete = true
			abandons.Merge(abandon)
			continue leftLoop
		case 1:
			abandons.Merge(abandon)
			continue leftLoop
		}
	}
	if !obsolete {
		result = append(result, target)
	}
	return result, abandons
}

func (a *Barricade) AddRule(rules []*tree.PriorityRule) {
	if rules == nil {
		return
	}
	a.rules = append(a.rules, rules...)
}

func (a *Barricade) RemoveRule(need func(rule *tree.PriorityRule) bool) {
	for index := 0; index < len(a.rules); index++ {
		rule := a.rules[index]
		if need(rule) {
			a.rules = append(a.rules[:index], a.rules[index+1:]...)
		}
	}
}
