package grammar

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type Reach struct {
	vocabularies []*tree.VocabularyRule
}

func NewReach() *Reach {
	return &Reach{}
}

func (r *Reach) Check(flow *lexer.Flow, onVocabulary func(*tree.VocabularyRule)) error {
	if flow.IsEnd() {
		return nil
	}
	word := flow.Peek()
	count := 0
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

func (g *Reach) AddRule(rules []*tree.VocabularyRule) {
	if rules == nil {
		return
	}
	g.vocabularies = append(g.vocabularies, rules...)
}

func (g *Reach) RemoveRule(need func(rule *tree.VocabularyRule) bool) {
	for index := 0; index < len(g.vocabularies); index++ {
		rule := g.vocabularies[index]
		if need(rule) {
			g.vocabularies = append(g.vocabularies[:index], g.vocabularies[index+1:]...)
		}
	}
}
