package tree

import (
	"github.com/TingerSure/natural_language/source"
)

type Vocabulary struct {
	word   string
	source source.Source
}

func (l *Vocabulary) GetWord() string {
	return l.word
}

func (l *Vocabulary) GetSource() source.Source {
	return l.source
}

func (l *Vocabulary) init(word string, source source.Source) *Vocabulary {
	l.word = word
	l.source = source
	return l
}

func NewVocabulary(word string, source source.Source) *Vocabulary {
	return (&Vocabulary{}).init(word, source)
}
