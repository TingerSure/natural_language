package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/tree"
	"regexp"
	"strings"
)

type Lexer struct {
	rules      []*tree.VocabularyRule
	wordCache  map[rune]map[string]map[*tree.VocabularyRule]bool
	matchCache map[*tree.VocabularyRule]*regexp.Regexp
}

func NewLexer() *Lexer {
	return &Lexer{
		rules:      []*tree.VocabularyRule{},
		wordCache:  map[rune]map[string]map[*tree.VocabularyRule]bool{},
		matchCache: map[*tree.VocabularyRule]*regexp.Regexp{},
	}
}

type lexerUnknownSection struct {
	left  int
	right int
}

func (p *Lexer) ParseVocabulary(road *Road) error {
	unknowns := []*lexerUnknownSection{}
	unknowns, err := p.parseVocabularyStep(road, unknowns, 0)
	if err != nil {
		return err
	}

	realUnknowns := []*lexerUnknownSection{}
	for cursor := len(unknowns) - 1; cursor >= 0; cursor-- {
		unknown := unknowns[cursor]
		if !p.existNext(road, unknown.right) {
			realUnknowns = append(realUnknowns, unknown)
			continue
		}
		err = p.removeUniqueVocabularyStep(road, unknown.left-1)
		if err != nil {
			return err
		}
	}
	if len(realUnknowns) != 0 {
		max := realUnknowns[0]
		for cursor := 1; cursor < len(realUnknowns); cursor++ {
			unknown := realUnknowns[cursor]
			if unknown.left > max.left {
				max = unknown
			}
		}
		return fmt.Errorf("This vocabulary has no rules to parse! (%v)\n\"%v\":%v", road.SubSentence(max.left, max.right+1), road.GetSentence(), max.left)
	}
	return nil
}

func (p *Lexer) removeUniqueVocabularyStep(road *Road, index int) error {
	if index < 0 {
		return nil
	}
	vocabularies := road.GetSections(index)
	for _, vocabulary := range vocabularies {
		left := index - vocabulary.ContentSize() + 1
		if road.GetLeftSectionSize(left) == 1 {
			err := p.removeUniqueVocabularyStep(road, left-1)
			if err != nil {
				return err
			}
		}
		err := road.RemoveSection(index, vocabulary)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Lexer) existNext(road *Road, index int) bool {
	for cursor := index; cursor < road.SentenceSize(); cursor++ {
		if road.HasRightSection(cursor) {
			return true
		}
	}
	return false
}

func (p *Lexer) parseVocabularyStep(road *Road, unknowns []*lexerUnknownSection, index int) ([]*lexerUnknownSection, error) {
	if index >= road.SentenceSize() || road.HasLeftSection(index) {
		return unknowns, nil
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
		return append(unknowns, &lexerUnknownSection{
			left:  index,
			right: end - 1,
		}), nil
	}

	for _, phrase := range phrases {
		err := road.AddSection(index+phrase.ContentSize()-1, phrase)
		if err != nil {
			return unknowns, err
		}
		unknowns, err = p.parseVocabularyStep(road, unknowns, index+phrase.ContentSize())
		if err != nil {
			return unknowns, err
		}
	}
	return unknowns, nil
}

func (p *Lexer) getPhrases(sentence string) []tree.Phrase {
	var phrases []tree.Phrase
	phrases = p.getPhrasesByWords(sentence, phrases)
	phrases = p.getPhrasesByMatch(sentence, phrases)
	return phrases
}

func (p *Lexer) getPhrasesByWords(sentence string, phrases []tree.Phrase) []tree.Phrase {
	for candidate, rules := range p.wordCache[[]rune(sentence)[0]] {
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
		start := []rune(word)[0]
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
			start := []rune(word)[0]
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
