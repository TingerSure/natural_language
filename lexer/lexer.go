package lexer

import (
	"github.com/TingerSure/natural_language/library/nl_string"
	"github.com/TingerSure/natural_language/tree"
)

type Lexer struct {
	sources        map[string]tree.Source
	naturalSources map[string]tree.Source
}

func (l *Lexer) getVocabulariesBySources(character string, sources map[string]tree.Source, vocabularies []*tree.Vocabulary) []*tree.Vocabulary {
	for _, source := range sources {
		var words []*tree.Word = source.GetWords(character)
		if words == nil {
			continue
		}
		for _, word := range words {
			vocabularies = append(vocabularies, tree.NewVocabulary(word, source))
		}
	}
	return vocabularies
}

func (l *Lexer) getVocabulary(character string) []*tree.Vocabulary {
	var vocabularies []*tree.Vocabulary
	vocabularies = l.getVocabulariesBySources(character, l.sources, vocabularies)
	vocabularies = l.getVocabulariesBySources(character, l.naturalSources, vocabularies)
	return vocabularies
}

func (l *Lexer) instanceStep(sentence string, index int, now *Flow, group *FlowGroup) {
	if index >= nl_string.Len(sentence) {
		return
	}
	var indexSentence string = nl_string.SubStringFrom(sentence, index)
	var firstCharacter string = nl_string.SubString(indexSentence, 0, 1)
	var vocabularies []*tree.Vocabulary = l.getVocabulary(firstCharacter)
	var count int = 0
	var base *Flow = now.Copy()

	for _, vocabulary := range vocabularies {
		if !vocabulary.GetWord().StartFor(indexSentence) {
			continue
		}
		if count == 0 {
			now.AddVocabulary(vocabulary)
			l.instanceStep(sentence, index+vocabulary.GetWord().Len(), now, group)
		} else {
			var new *Flow = base.Copy()
			group.AddInstance(new)
			new.AddVocabulary(vocabulary)
			l.instanceStep(sentence, index+vocabulary.GetWord().Len(), new, group)
		}
		count++
	}
	if count == 0 {
		var vocabulary *tree.Vocabulary = tree.NewVocabulary(tree.NewUnknownWord(firstCharacter), nil)
		now.AddVocabulary(vocabulary)
		l.instanceStep(sentence, index+vocabulary.GetWord().Len(), now, group)
	}
}

func (l *Lexer) Instances(sentence string) *FlowGroup {
	var group *FlowGroup = NewFlowGroup()
	var now *Flow = NewFlow()
	now.SetSentence(sentence)
	group.AddInstance(now)
	l.instanceStep(sentence, 0, now, group)
	return group
}

func (l *Lexer) AddNaturalSource(source tree.Source) {
	l.naturalSources[source.GetName()] = source
}

func (l *Lexer) RemoveNaturalSource(name string) {
	l.naturalSources[name] = nil
}

func (l *Lexer) AddSource(source tree.Source) {
	l.sources[source.GetName()] = source
}

func (l *Lexer) RemoveSource(name string) {
	l.sources[name] = nil
}

func (l *Lexer) init() *Lexer {
	l.sources = map[string]tree.Source{}
	l.naturalSources = map[string]tree.Source{}
	return l
}

func NewLexer() *Lexer {
	return (&Lexer{}).init()
}
