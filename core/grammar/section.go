package grammar

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type Section struct {
	vocabularies []*tree.VocabularyRule
}

func NewSection() *Section {
	return &Section{}
}

func (r *Section) Check(flow *lexer.Flow, onVocabulary func(*tree.VocabularyRule)) error {
	if flow.IsEnd() {
		return nil
	}
	word := flow.Peek()
	count := 0

	source := word.GetSource()
	if !nl_interface.IsNil(source) {
		rules := source.GetVocabularyRules()
		for _, leaf := range rules {
			if leaf.Match(word) {
				onVocabulary(leaf)
				count++
			}
		}
		if count != 0 {
			return nil
		}
	}

	for _, leaf := range r.vocabularies {
		if leaf.Match(word) {
			onVocabulary(leaf)
			count++
		}
	}
	if count == 0 {
		return errors.New(fmt.Sprintf("This vocabulary has no rules to parse! ( %v )", word.ToString()))
	}
	return nil
}

func (g *Section) AddRule(rules []*tree.VocabularyRule) {
	if rules == nil {
		return
	}
	g.vocabularies = append(g.vocabularies, rules...)
}

func (g *Section) RemoveRule(need func(rule *tree.VocabularyRule) bool) {
	for index := 0; index < len(g.vocabularies); index++ {
		rule := g.vocabularies[index]
		if need(rule) {
			g.vocabularies = append(g.vocabularies[:index], g.vocabularies[index+1:]...)
		}
	}
}
