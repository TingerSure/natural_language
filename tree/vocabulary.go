package tree

import (
	"fmt"
)

type Vocabulary struct {
	word   *Word
	source Source
}

func (l *Vocabulary) GetWord() *Word {
	return l.word
}

func (l *Vocabulary) GetSource() Source {
	return l.source
}

func (l *Vocabulary) init(word *Word, source Source) *Vocabulary {
	l.word = word
	l.source = source
	return l
}

func (l *Vocabulary) ToString() string {
	sourceName := "unknown"
	if l.source != nil {
		sourceName = l.source.GetName()

	}
	return fmt.Sprintf("%v (%v)", l.word.GetContext(), sourceName)
}

func NewVocabulary(word *Word, source Source) *Vocabulary {
	return (&Vocabulary{}).init(word, source)
}
