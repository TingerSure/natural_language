package grammar

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/lexer"
	"github.com/TingerSure/natural_language/tree"
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

func (l *Grammar) Instances(flow *lexer.Flow) ([]*River, error) {
	flow.Reset()
	wait := NewCollector()
	river := NewRiver(wait, flow)
	bases, err := river.Step(l.vocabularies, l.structs)
	if err != nil {
		return nil, err
	}
	bases, err = l.riversfilter(bases)
	if err != nil {
		return nil, err
	}
	return bases, nil
}

func (l *Grammar) riversfilter(inputs []*River) ([]*River, error) {
	if len(inputs) == 0 {
		return nil, errors.New("This is an empty sentence!")
	}
	actives := []*River{}
	var min *River = nil
	for _, input := range inputs {
		if input.IsActive() {
			actives = append(actives, input)
			continue
		}
		if min == nil {
			min = input
		}
		if input.GetWait().Len() < min.GetWait().Len() {
			min = input
		}
	}
	if len(actives) == 0 {
		return nil, errors.New(fmt.Sprintf("There is an unknown rule in this sentence!\n%v", min.ToString()))
	}
	return actives, nil
}

func (l *Grammar) init() *Grammar {
	return l
}

func NewGrammar() *Grammar {
	return (&Grammar{}).init()
}
