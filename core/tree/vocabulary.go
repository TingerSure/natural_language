package tree

import (
	"strings"
	"unicode/utf8"
)

type Vocabulary struct {
	context string
}

func (w *Vocabulary) GetContext() string {
	return w.context
}

func (l *Vocabulary) init(context string) *Vocabulary {
	l.context = context
	return l
}

func (l *Vocabulary) ToString() string {
	return l.context
}

func (w *Vocabulary) StartFor(sentence string) bool {
	return 0 == strings.Index(sentence, w.context)
}

func (w *Vocabulary) Len() int {
	return utf8.RuneCountInString(w.context)
}

func NewVocabulary(context string) *Vocabulary {
	return (&Vocabulary{
		context: context,
	})
}
