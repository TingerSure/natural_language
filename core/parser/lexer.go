package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/tree"
	"regexp"
	"strings"
)

type Lexer struct {
	rules      []*tree.VocabularyRule
	wordCache  map[byte]map[string]map[*tree.VocabularyRule]bool
	matchCache map[*tree.VocabularyRule]*regexp.Regexp
}

func NewLexer() *Lexer {
	return &Lexer{
		rules:      []*tree.VocabularyRule{},
		wordCache:  map[byte]map[string]map[*tree.VocabularyRule]bool{},
		matchCache: map[*tree.VocabularyRule]*regexp.Regexp{},
	}
}

func (p *Lexer) ParseVocabulary(road *Road) error {
	return p.parseVocabularyStep(road, 0)
}

func (p *Lexer) parseVocabularyStep(road *Road, index int) error {
	if index >= road.SentenceSize() || road.HasLeftSection(index) {
		return nil
	}

	phrases := p.getPhrases(road.SubSentenceFrom(index))
	if len(phrases) == 0 {
		end := index + 1
		for ; end < road.SentenceSize(); end++ {
			phrases = p.getPhrases(road.SubSentenceFrom(end))
			if len(phrases) != 0 {
				break
			}
		}
		return fmt.Errorf("This vocabulary has no rules to parse! ( %v )", road.SubSentence(index, end))
	}

	for _, phrase := range phrases {
		road.AddSection(index+phrase.ContentSize()-1, phrase)
		err := p.parseVocabularyStep(road, index+phrase.ContentSize())
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Lexer) getPhrases(sentence string) []tree.Phrase {
	var phrases []tree.Phrase
	phrases = p.getPhrasesByWords(sentence, phrases)
	phrases = p.getPhrasesByMatch(sentence, phrases)
	return phrases
}

func (p *Lexer) getPhrasesByWords(sentence string, phrases []tree.Phrase) []tree.Phrase {
	for candidate, rules := range p.wordCache[sentence[0]] {
		if p.startFor(candidate, sentence) {
			for rule, _ := range rules {
				phrases = append(phrases, rule.Create(candidate))
			}
		}
	}
	return phrases
}

func (p *Lexer) getPhrasesByMatch(sentence string, phrases []tree.Phrase) []tree.Phrase {
	for rule, match := range p.matchCache {
		value := match.FindString(sentence)
		if value != "" {
			phrases = append(phrases, rule.Create(value))
		}
	}
	return phrases
}

func (p *Lexer) startFor(vocabulary, sentence string) bool {
	return 0 == strings.Index(sentence, vocabulary)
}

func (p *Lexer) AddRule(rule *tree.VocabularyRule) {
	if rule == nil {
		return
	}
	p.rules = append(p.rules, rule)
	for _, word := range rule.Words() {
		if word == "" {
			continue
		}
		start := word[0]
		if p.wordCache[start] == nil {
			p.wordCache[start] = map[string]map[*tree.VocabularyRule]bool{}
		}
		if p.wordCache[start][word] == nil {
			p.wordCache[start][word] = map[*tree.VocabularyRule]bool{}
		}
		p.wordCache[start][word][rule] = true
	}
	template := rule.Match()
	if template != "" {
		match := regexp.MustCompile("^" + template)
		p.matchCache[rule] = match
	}
}

func (p *Lexer) RemoveRule(need func(*tree.VocabularyRule) bool) {
	for index := 0; index < len(p.rules); index++ {
		rule := p.rules[index]
		if !need(rule) {
			continue
		}
		p.rules = append(p.rules[:index], p.rules[index+1:]...)
		index--
		for _, word := range rule.Words() {
			if word == "" {
				continue
			}
			start := word[0]
			delete(p.wordCache[start][word], rule)
			if len(p.wordCache[start][word]) == 0 {
				delete(p.wordCache[start], word)
			}
			if len(p.wordCache[start]) == 0 {
				delete(p.wordCache, start)
			}
		}
		delete(p.matchCache, rule)
	}
}
