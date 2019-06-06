package tree

import (
	"github.com/TingerSure/natural_language/source"
	"github.com/TingerSure/natural_language/word"
)

type Vocabulary struct {
	word   *word.Word
	source source.Source
}

func (l *Vocabulary) GetWord() *word.Word {
	return l.word
}

func (l *Vocabulary) GetSource() source.Source {
	return l.source
}

func (l *Vocabulary) init(word *word.Word, source source.Source) *Vocabulary {
	l.word = word
	l.source = source
	return l
}

func NewVocabulary(word *word.Word, source source.Source) *Vocabulary {
	return (&Vocabulary{}).init(word, source)
}
