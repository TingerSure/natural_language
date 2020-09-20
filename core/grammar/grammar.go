package grammar

import (
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type Grammar struct {
	reach   *Reach
	section *Section
	dam     *Dam
}

func (g *Grammar) GetDam() *Dam {
	return g.dam
}

func (g *Grammar) AddPriorityRule(rules []*tree.PriorityRule) {
	g.dam.AddRule(rules)
}

func (g *Grammar) AddStructRule(rules []*tree.StructRule) {
	g.reach.AddRule(rules)
}

func (g *Grammar) RemoveStructRule(need func(rule *tree.StructRule) bool) {
	g.reach.RemoveRule(need)
}

func (g *Grammar) AddVocabularyRule(rules []*tree.VocabularyRule) {
	g.section.AddRule(rules)
}

func (g *Grammar) RemoveVocabularyRule(need func(rule *tree.VocabularyRule) bool) {
	g.section.RemoveRule(need)
}

func (l *Grammar) Instances(flow *lexer.Flow) (*Valley, error) {
	flow.Reset()
	valley := NewValley()
	err := valley.Step(flow, l.section, l.reach, l.dam)
	if err != nil {
		return nil, err
	}

	return valley, nil
}

func (l *Grammar) init() *Grammar {
	return l
}

func NewGrammar() *Grammar {
	return (&Grammar{
		dam:     NewDam(),
		section: NewSection(),
		reach:   NewReach(),
	}).init()
}
