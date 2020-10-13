package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type Lexer struct {
	section *Section
	sources map[string]tree.Source
}

func NewLexer(section *Section) *Lexer {
	return &Lexer{
		sources: map[string]tree.Source{},
		section: section,
	}
}

func (p *Lexer) ParseVocabulary(road *Road) error {
	return p.parseVocabularyStep(road, 0)
}

func (p *Lexer) parseVocabularyStep(road *Road, index int) error {

	if index >= road.SentenceSize() || road.HasLeftSection(index) {
		return nil
	}

	vocabularies := p.getVocabularies(road.SubSentenceFrom(index))
	if len(vocabularies) == 0 {
		end := index + 1
		for ; end < road.SentenceSize(); end++ {
			vocabularies = p.getVocabularies(road.SubSentenceFrom(end))
			if len(vocabularies) != 0 {
				break
			}
		}
		unknown := tree.NewVocabulary(road.SubSentence(index, end), nil)
		rules, err := p.section.Check(unknown)
		if err != nil {
			return err
		}
		for _, rule := range rules {
			phrase := rule.Create(unknown)
			road.AddLeftSection(index, phrase)
		}
		if end == road.SentenceSize() {
			return nil
		}
		index = end
	}

	for _, vocabulary := range vocabularies {
		rules, err := p.section.Check(vocabulary)
		if err != nil {
			return err
		}
		for _, rule := range rules {
			phrase := rule.Create(vocabulary)
			road.AddLeftSection(index, phrase)
			err = p.parseVocabularyStep(road, index+phrase.ContentSize())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Lexer) getVocabularies(sentence string) []*tree.Vocabulary {
	var vocabularies []*tree.Vocabulary
	vocabularies = p.getVocabulariesBySources(sentence, p.sources, vocabularies)
	return vocabularies
}

func (p *Lexer) getVocabulariesBySources(sentence string, sources map[string]tree.Source, vocabularies []*tree.Vocabulary) []*tree.Vocabulary {
	for _, source := range sources {
		var words []*tree.Vocabulary = source.GetWords(sentence)
		if words == nil {
			continue
		}
		vocabularies = append(vocabularies, words...)
	}
	return vocabularies
}

func (p *Lexer) AddSource(source tree.Source) {
	p.sources[source.GetName()] = source
}

func (p *Lexer) GetSource(name string) tree.Source {
	return p.sources[name]
}

func (p *Lexer) RemoveSource(name string) {
	p.sources[name] = nil
}
