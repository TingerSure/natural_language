package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type Grammar struct {
	reach     *Reach
	barricade *Barricade
	diversion *Diversion
	types     *Types
}

func NewGrammar(types *Types, reach *Reach, barricade *Barricade, diversion *Diversion) *Grammar {
	return &Grammar{
		types:     types,
		reach:     reach,
		barricade: barricade,
		diversion: diversion,
	}
}

func (g *Grammar) ParseStruct(road *Road) error {
	for index := 0; index < road.SentenceSize(); index++ {
		if !road.HasRightSection(index) {
			continue
		}

		originSources := road.GetSections(index)
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
					err := g.match(road, index, origin, rule, targets)
					if err != nil {
						return err
					}
				}
			}
			for target, _ := range targets {
				if !road.DependencyCheck(target) {
					continue
				}
				olds := road.GetSectionByTypesAndSize(index, target.Types(), target.ContentSize())
				if len(olds) == 0 {
					road.AddSection(index, target)
					activeTargets[target] = true
					continue
				}

				targetNeed := true
				for _, old := range olds {
					results, abandons := g.barricade.Check(old, target)
					g.cut(road, index, abandons)
					switch results {
					case 0:
						break
					case -1:
						targetNeed = false
						break
					case 1:
						road.RemoveSection(index, old)
						break
					}
				}
				if targetNeed {
					road.AddSection(index, target)
					activeTargets[target] = true
				}
			}
			origins = activeTargets
			targets = map[tree.Phrase]bool{}
			activeTargets = map[tree.Phrase]bool{}
		}
	}
	return nil
}

func (g *Grammar) cut(road *Road, index int, abandons *tree.AbandonGroup) {
	if abandons == nil {
		return
	}
	for _, abandon := range abandons.Values() {
		road.RemoveSection(index+abandon.Offset, abandon.Value)
	}
}

func (g *Grammar) match(road *Road, roadIndex int, last tree.Phrase, rule *tree.StructRule, back map[tree.Phrase]bool) error {
	size := rule.Size()
	treasures := make([]tree.Phrase, size, size)
	treasures[size-1] = g.types.Package(rule.Types()[size-1], last.Types(), last)
	if size == 1 {
		section := rule.Create(treasures)
		if section.Types() == "" {
			sectionTypes, err := g.diversion.Match(section)
			if err != nil {
				return err
			}
			section.SetTypes(sectionTypes)
		}
		back[section] = true
		return nil
	}
	return g.matchStep(road, size-2, roadIndex-last.ContentSize(), treasures, rule, back)
}

func (g *Grammar) matchStep(road *Road, index int, cursor int, treasures []tree.Phrase, rule *tree.StructRule, back map[tree.Phrase]bool) error {
	if cursor < 0 {
		return nil
	}
	phrases := road.GetSectionByTypes(cursor, rule.Types()[index])
	for phrase, _ := range phrases {
		treasures[index] = g.types.Package(rule.Types()[index], phrase.Types(), phrase)
		if index == 0 {
			section := rule.Create(treasures)
			if section.Types() == "" {
				sectionTypes, err := g.diversion.Match(section)
				if err != nil {
					return err
				}
				section.SetTypes(sectionTypes)
			}
			back[section] = true
		} else {
			err := g.matchStep(road, index-1, cursor-phrase.ContentSize(), treasures, rule, back)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
