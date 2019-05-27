package lexer

import (
	"fmt"
	"github.com/TingerSure/natural_language/source"
	"github.com/TingerSure/natural_language/tree"
	"strings"
)

type Lexer struct {
	sources        map[string]source.Source
	naturalSources map[string]source.Source
}

func (l *Lexer) getVocabulariesBySources(character string, sources map[string]source.Source, vocabularies []*tree.Vocabulary) []*tree.Vocabulary {
	for index := range sources {
		var words []string = sources[index].GetWords(character)
		for i := 0; i < len(words); i++ {
			vocabularies = append(vocabularies, tree.NewVocabulary(words[i], sources[index]))
		}
	}
	return vocabularies
}

func (l *Lexer) getVocabulary(character string) []*tree.Vocabulary {
	var vocabularies []*tree.Vocabulary
	vocabularies = l.getVocabulariesBySources(character, l.sources, vocabularies)
	vocabularies = l.getVocabulariesBySources(character, l.naturalSources, vocabularies)
    fmt.Printf("== %s vocabularies : %s \n"  ,character,vocabularies)

	if len(vocabularies) == 0 {
        // fmt.Printf("== now, vocabularies nil : %s \n" ,character)
		vocabularies = append(vocabularies, tree.NewVocabulary(character, nil))
	}
	return vocabularies
}

func (l *Lexer) instanceStep(sentence string, index int, now *LexerInstance, group *LexerInstanceGroup) {
	if index >= len(sentence) {
		return
	}
	var vocabularies []*tree.Vocabulary = l.getVocabulary(sentence[index : index+1])
    for _, v := range vocabularies {
        fmt.Printf("++ %s  %s\n",v.GetWord(), v.GetSource())
    }

    var count int = 0

    var base *LexerInstance = now.Copy()

	for _, vocabulary := range vocabularies {
		if strings.Index(sentence[index:], vocabulary.GetWord()) != 0 {
			continue
		}
		if count == 0 {
			now.AddVocabulary(vocabulary)
			l.instanceStep(sentence, index+len(vocabulary.GetWord()), now, group)
		} else {
            var new *LexerInstance = base.Copy()
			group.AddInstance(new)
			new.AddVocabulary(vocabulary)
			l.instanceStep(sentence, index+len(vocabulary.GetWord()), new, group)
		}
		count++
	}
}

func (l *Lexer) Instances(sentence string) *LexerInstanceGroup {
	var group *LexerInstanceGroup= NewLexerInstanceGroup()
    var now *LexerInstance = NewLexerInstance()
    group.AddInstance(now)
	l.instanceStep(sentence, 0, now , group)
    return group
}

func (l *Lexer) AddNaturalSource(name string, source source.Source) {
	l.naturalSources[name] = source
}

func (l *Lexer) RemoveNaturalSource(name string) {
	l.naturalSources[name] = nil
}

func (l *Lexer) AddSource(name string, source source.Source) {
	l.sources[name] = source
}

func (l *Lexer) RemoveSource(name string) {
	l.sources[name] = nil
}

func (l *Lexer) init() *Lexer {
	l.sources = map[string]source.Source{}
	l.naturalSources = map[string]source.Source{}
	return l
}

func NewLexer() *Lexer {
	return (&Lexer{}).init()
}
