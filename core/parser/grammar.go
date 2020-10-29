package parser

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/tree"
)

type Grammar struct {
	reach     *Reach
	barricade *Barricade
}

func NewGrammar(reach *Reach, barricade *Barricade) *Grammar {
	return &Grammar{
		reach:     reach,
		barricade: barricade,
	}
}

func (g *Grammar) ParseStruct(road *Road) error {
	for index := 0; index < road.SentenceSize(); index++ {
		if !road.HasRightSection(index) {
			continue
		}

		originSources := road.GetRightSection(index, nil)
		origins := map[tree.Phrase]bool{}
		for _, origin := range originSources {
			origins[origin] = true
		}
		targets := map[tree.Phrase]bool{}
		activeTargets := map[tree.Phrase]bool{}
		for len(origins) > 0 {
			for origin, _ := range origins {
				rules := g.reach.GetRulesByLastType(origin.Types())
				for _, rule := range rules {
					g.match(road, index, origin, rule, targets)
				}
			}
			for target, _ := range targets {
				old := road.GetRightSectionByTypesAndSize(index, target.Types(), target.ContentSize())
				if nl_interface.IsNil(old) {
					road.AddRightSection(index, target)
					activeTargets[target] = true
					continue
				}
				if priority, ok := old.(*tree.PhrasePriority); ok {
					results, abandons := g.barricade.TargetFilter(priority.AllValues(), target)
					g.cut(road, index, abandons)
					if 1 == len(results) {
						g.replace(targets, old, results[0])
						road.ReplaceRight(index, old, results[0])
					} else {
						priority.SetValues(results)
					}
					continue
				}
				results, abandons := g.barricade.Check(old, target)
				g.cut(road, index, abandons)
				switch results {
				case 0:
					multiple := tree.NewPhrasePriority([]tree.Phrase{
						old,
						target,
					})
					g.replace(targets, old, multiple)
					road.ReplaceRight(index, old, multiple)
					break
				case -1:
					// Do Nothing
					break
				case 1:
					g.replace(targets, old, target)
					road.ReplaceRight(index, old, target)
					break
				}
			}
			origins = activeTargets
			targets = map[tree.Phrase]bool{}
			activeTargets = map[tree.Phrase]bool{}
		}
	}
	return nil
}

func (g *Grammar) replace(targets map[tree.Phrase]bool, from, to tree.Phrase) {
	replace := func(target tree.Phrase) {
		if target == nil {
			return
		}
		for index := 0; index < target.Size(); index++ {
			if target.GetChild(index) == from {
				target.SetChild(index, to)
			}
		}
	}
	for target, _ := range targets {
		if priority, ok := target.(*tree.PhrasePriority); ok {
			for index := 0; index < priority.ValueSize(); index++ {
				replace(priority)
			}
		}
		replace(target)
	}
}

func (g *Grammar) cut(road *Road, index int, abandons *tree.AbandonGroup) {
	if abandons == nil {
		return
	}
	for _, abandon := range abandons.Values() {
		road.RemoveRightSection(index+abandon.Offset, func(phrase tree.Phrase) bool {
			return phrase == abandon.Value
		})
	}
}

func (g *Grammar) match(road *Road, roadIndex int, last tree.Phrase, rule *tree.StructRule, back map[tree.Phrase]bool) {
	size := rule.Size()
	treasures := make([]tree.Phrase, size, size)
	treasures[size-1] = last
	if size == 1 {
		back[rule.Create(treasures)] = true
		return
	}
	g.matchStep(road, size-2, roadIndex-last.ContentSize(), treasures, rule, back)
}

func (g *Grammar) matchStep(road *Road, index int, cursor int, treasures []tree.Phrase, rule *tree.StructRule, back map[tree.Phrase]bool) {
	if cursor < 0 {
		return
	}
	phrases := road.GetRightSectionByTypes(cursor, rule.Types()[index])
	for phrase, _ := range phrases {
		treasures[index] = phrase
		if index == 0 {
			back[rule.Create(treasures)] = true
		} else {
			g.matchStep(road, index-1, cursor-phrase.ContentSize(), treasures, rule, back)
		}
	}
}
