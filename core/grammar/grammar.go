package grammar

import (
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type Grammar struct {
	structs      []*tree.StructRule
	vocabularies []*tree.VocabularyRule
}

func (g *Grammar) AddStructRule(rules []*tree.StructRule) {
	if rules == nil {
		return
	}
	g.structs = append(g.structs, rules...)
}

func (g *Grammar) RemoveStructRule(need func(rule *tree.StructRule) bool) {
	for index := 0; index < len(g.structs); index++ {
		rule := g.structs[index]
		if need(rule) {
			g.structs = append(g.structs[:index], g.structs[index+1:]...)
		}
	}
}

func (g *Grammar) AddVocabularyRule(rules []*tree.VocabularyRule) {
	if rules == nil {
		return
	}
	g.vocabularies = append(g.vocabularies, rules...)
}

func (g *Grammar) RemoveVocabularyRule(need func(rule *tree.VocabularyRule) bool) {
	for index := 0; index < len(g.vocabularies); index++ {
		rule := g.vocabularies[index]
		if need(rule) {
			g.vocabularies = append(g.vocabularies[:index], g.vocabularies[index+1:]...)
		}
	}
}

func (l *Grammar) Instances(flow *lexer.Flow) (*Valley, error) {
	flow.Reset()
	valley := NewValley()
	err := valley.Step(flow, l.vocabularies, l.structs)
	if err != nil {
		return nil, err
	}
	valley, err = valley.Filter()
	if err != nil {
		return nil, err
	}
	return valley, nil
}

func (l *Grammar) init() *Grammar {
	return l
}

func NewGrammar() *Grammar {
	return (&Grammar{}).init()
}
