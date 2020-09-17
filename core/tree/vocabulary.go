package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_string"
	"unicode/utf8"
)

type Vocabulary struct {
	context string
	source  Source
}

func (w *Vocabulary) GetContext() string {
	return w.context
}

func (l *Vocabulary) GetSource() Source {
	return l.source
}

func (l *Vocabulary) init(context string, source Source) *Vocabulary {
	l.context = context
	l.source = source
	return l
}

func (l *Vocabulary) ToString() string {
	sourceName := "unknown"
	if l.source != nil {
		sourceName = l.source.GetName()

	}
	return fmt.Sprintf("%v (%v)", l.context, sourceName)
}

func (w *Vocabulary) StartFor(sentence string) bool {
	return 0 == nl_string.Index(sentence, w.context)
}

func (w *Vocabulary) StartWith(first string) bool {
	return 0 == nl_string.Index(w.context, first)

}

func (w *Vocabulary) Len() int {
	return utf8.RuneCountInString(w.context)
}

func NewVocabulary(context string, source Source) *Vocabulary {
	return (&Vocabulary{
		context: context,
		source:  source,
	})
}

func VocabularysFilter(words []*Vocabulary, sentence string) []*Vocabulary {
	if len(words) == 1 && words[0].StartFor(sentence) {
		return words
	}

	var targets []*Vocabulary

	for _, word := range words {
		if word.StartFor(sentence) {
			targets = append(targets, word)
		}
	}
	return targets
}
