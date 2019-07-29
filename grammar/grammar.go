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

func (g *Grammar) AddStructRule(rule ...*tree.StructRule) {
	g.structs = append(g.structs, rule...)
}

func (g *Grammar) RemoveStructRule(need func(rule *tree.StructRule) bool) {
	for index := 0; index < len(g.structs); index++ {
		rule := g.structs[index]
		if need(rule) {
			g.structs = append(g.structs[:index], g.structs[index+1:]...)
		}
	}
}

func (g *Grammar) AddVocabularyRule(rule ...*tree.VocabularyRule) {
	g.vocabularies = append(g.vocabularies, rule...)
}

func (g *Grammar) RemoveVocabularyRule(need func(rule *tree.VocabularyRule) bool) {
	for index := 0; index < len(g.vocabularies); index++ {
		rule := g.vocabularies[index]
		if need(rule) {
			g.vocabularies = append(g.vocabularies[:index], g.vocabularies[index+1:]...)
		}
	}
}

func (l *Grammar) testStruct(wait *Collector) bool {

	if wait.IsEmpty() {
		return false
	}
	for _, rule := range l.structs {
		if wait.Len() < rule.Size() {
			continue
		}
		phrase := rule.Logic(wait.PeekMultiple(rule.Size()))
		if phrase != nil {
			wait.PopMultiple(rule.Size())
			wait.Push(phrase)
			return true
		}
	}
	return false
}
func (l *Grammar) testVocabulary(wait *Collector, vocabulary *tree.Vocabulary) bool {
	for _, rule := range l.vocabularies {
		phrase := rule.Logic(vocabulary)
		if phrase != nil {
			wait.Push(phrase)
			return true
		}
	}
	return false
}
func (l *Grammar) instanceStep(flow *lexer.Flow, wait *Collector) (bool, error) {
	if !wait.IsSingle() {
		for l.testStruct(wait) {
			//Do nothing
		}
	}
	if flow.IsEnd() {
		if wait.IsEmpty() {
			return false, errors.New("This is empty sentence!")
		}
		if !wait.IsSingle() {
			return false, errors.New("There is a syntax error in this sentence!")
		}
		return true, nil
	}

	word := flow.Next()

	if !l.testVocabulary(wait, word) {
		return false, errors.New(fmt.Sprintf("This vocabulary has no rules to parse! ( %v )", word.ToString()))
	}

	return false, nil
}

func (l *Grammar) Instances(flow *lexer.Flow) (tree.Phrase, error) {
	flow.Reset()
	wait := NewCollector()
	for {
		end, err := l.instanceStep(flow, wait)
		if err != nil {
			return nil, err
		}
		if end {
			break
		}
	}
	return wait.Peek(), nil
}

func (l *Grammar) init() *Grammar {
	return l
}

func NewGrammar() *Grammar {
	return (&Grammar{}).init()
}
