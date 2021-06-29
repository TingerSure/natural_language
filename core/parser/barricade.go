package parser

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
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

func (a *Barricade) Check(left, right tree.Phrase) (*tree.PriorityResult, concept.Exception) {
	for _, rule := range a.rules {
		yes, exception := rule.Match(left, right)
		if !nl_interface.IsNil(exception) {
			return nil, exception
		}
		if yes {
			return rule.Choose(left, right)
		}
	}
	return tree.NewPriorityResult(0), nil
}

func (a *Barricade) AddRule(rule *tree.PriorityRule) {
	if rule == nil {
		return
	}
	a.rules = append(a.rules, rule)
}

func (a *Barricade) RemoveRule(need func(rule *tree.PriorityRule) bool) {
	for index := 0; index < len(a.rules); index++ {
		rule := a.rules[index]
		if need(rule) {
			a.rules = append(a.rules[:index], a.rules[index+1:]...)
		}
	}
}
