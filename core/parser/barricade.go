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

func (a *Barricade) Check(left, right tree.Phrase) int {
	for _, rule := range a.rules {
		if rule.Match(left, right) {
			return rule.Choose(left, right)
		}
	}
	return 0
}

func (a *Barricade) TargetFilter(phrases []tree.Phrase, target tree.Phrase) []tree.Phrase {
	result := []tree.Phrase{}
	obsolete := false
leftLoop:
	for _, left := range phrases {
		switch a.Check(left, target) {
		case 0:
			result = append(result, left)
			continue leftLoop
		case -1:
			result = append(result, left)
			obsolete = true
			continue leftLoop
		case 1:
			continue leftLoop
		}
	}
	if !obsolete {
		result = append(result, target)
	}
	return result
}

func (a *Barricade) Filter(phrases []tree.Phrase) []tree.Phrase {
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
			if nl_interface.IsNil(left) && nl_interface.IsNil(right) {
				continue
			}
			switch a.Check(left, right) {
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

func (a *Barricade) DeepFilter(phrases []tree.Phrase) []tree.Phrase {
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
