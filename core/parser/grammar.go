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

		phrases := road.GetRightSection(index, nil)
		targetPhrases := []tree.Phrase{}
		for len(phrases) > 0 {
			for _, phrase := range phrases {
				rules := g.reach.GetRulesByLastType(phrase.Types())
				for _, rule := range rules {
					targetPhrases = g.Match(road, index, phrase, rule, targetPhrases)
				}
			}
			for _, phrase := range targetPhrases {
				old := road.GetRightSectionByTypesAndSize(index, phrase.Types(), phrase.ContentSize())
				if nl_interface.IsNil(old) {
					road.AddRightSection(index, phrase)
					continue
				}
				if priority, ok := old.(*tree.PhrasePriority); ok {
					results := g.barricade.TargetFilter(priority.AllValues(), phrase)
					if 1 == len(results) {
						road.ReplaceRight(index, old, results[0])
					} else {
						priority.SetValues(results)
					}
					continue
				}
				switch g.barricade.Check(old, phrase) {
				case 0:
					road.ReplaceRight(index, old, tree.NewPhrasePriority([]tree.Phrase{
						old,
						phrase,
					}))
					break
				case -1:
					// Do Nothing
					break
				case 1:
					road.ReplaceRight(index, old, phrase)
					break
				}

			}
			phrases = targetPhrases
			targetPhrases = []tree.Phrase{}
		}

	}
	return nil
}

func (g *Grammar) Match(road *Road, roadIndex int, last tree.Phrase, rule *tree.StructRule, back []tree.Phrase) []tree.Phrase {
	size := rule.Size()
	treasures := make([]tree.Phrase, size, size)
	treasures[size-1] = last
	return g.matchStep(road, size-2, roadIndex-last.ContentSize(), treasures, rule, back)
}

func (g *Grammar) matchStep(road *Road, index int, cursor int, treasures []tree.Phrase, rule *tree.StructRule, back []tree.Phrase) []tree.Phrase {
	if cursor < 0 {
		return back
	}
	phrases := road.GetRightSectionByTypes(cursor, rule.Types()[index])
	for phrase, _ := range phrases {
		treasures[index] = phrase
		if index == 0 {
			back = append(back, rule.Create(treasures))
		} else {
			back = g.matchStep(road, index-1, cursor-phrase.ContentSize(), treasures, rule, back)
		}
	}
	return back
}
