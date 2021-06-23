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
				originTypes, err := origin.Types()
				if err != nil {
					return err
				}
				rules := g.reach.GetRulesByLastType(originTypes)
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
				targetTypes, err := target.Types()
				if err != nil {
					return err
				}
				olds, oldsErr := road.GetSectionByTypesAndSize(index, targetTypes, target.ContentSize())
				if oldsErr != nil {
					return oldsErr
				}
				if len(olds) == 0 {
					err := road.AddSection(index, target)
					if err != nil {
						return err
					}
					activeTargets[target] = true
					continue
				}

				targetNeed := true
				for _, old := range olds {
					results, abandons := g.barricade.Check(old, target)
					err := g.cut(road, index, abandons)
					if err != nil {
						return err
					}
					switch results {
					case 0:
						break
					case -1:
						targetNeed = false
						break
					case 1:
						err := road.RemoveSection(index, old)
						if err != nil {
							return err
						}
						break
					}
				}
				if targetNeed {
					err := road.AddSection(index, target)
					if err != nil {
						return err
					}
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

func (g *Grammar) cut(road *Road, index int, abandons *tree.AbandonGroup) error {
	if abandons == nil {
		return nil
	}
	for _, abandon := range abandons.Values() {
		err := road.RemoveSection(index+abandon.Offset, abandon.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Grammar) match(road *Road, roadIndex int, last tree.Phrase, rule *tree.StructRule, back map[tree.Phrase]bool) error {
	size := rule.Size()
	treasures := make([]tree.Phrase, size, size)
	lastTypes, err := last.Types()
	if err != nil {
		return err
	}
	treasures[size-1] = g.types.Package(rule.Types()[size-1], lastTypes, last)
	if size == 1 {
		section := rule.Create(treasures)
		sectionTypes, err := section.Types()
		if err != nil {
			return err
		}
		if sectionTypes == "" {
			sectionTypes, matchErr := g.diversion.Match(section)
			if matchErr != nil {
				return matchErr
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
		phraseTypes, err := phrase.Types()
		if err != nil {
			return err
		}
		treasures[index] = g.types.Package(rule.Types()[index], phraseTypes, phrase)
		if index == 0 {
			section := rule.Create(treasures)
			sectionTypes, err := section.Types()
			if err != nil {
				return err
			}
			if sectionTypes == "" {
				sectionTypes, matchErr := g.diversion.Match(section)
				if matchErr != nil {
					return matchErr
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
