package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/word"
)

type Vocabulary struct {
	word   *word.Word
	source Source
}

func (l *Vocabulary) GetWord() *word.Word {
	return l.word
}

func (l *Vocabulary) GetSource() Source {
	return l.source
}

func (l *Vocabulary) init(word *word.Word, source Source) *Vocabulary {
	l.word = word
	l.source = source
	return l
}

func (l *Vocabulary) ToString() string {
	return fmt.Sprintf("%v ( %v )", l.word.GetContext(), l.source.GetName())
}

func NewVocabulary(word *word.Word, source Source) *Vocabulary {
	return (&Vocabulary{}).init(word, source)
}
