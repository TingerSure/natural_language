package word

import (
	"github.com/TingerSure/natural_language/library/nl_string"
)

type Word struct {
	context string
	types   int
}

func (w *Word) GetContext() string {
	return w.context
}

func (w *Word) GetTypes() int {
	return w.types
}

func (w *Word) StartFor(sentence string) bool {
	return 0 == nl_string.Index(sentence, w.context)
}

func (w *Word) StartWith(first string) bool {
	return 0 == nl_string.Index(w.context, first)

}

func (w *Word) Len() int {
	return nl_string.Len(w.context)
}

func NewWord(context string, types int) *Word {
	return (&Word{
		context: context,
		types:   types,
	})

}

func NewUnknownWord(context string) *Word {
	return NewWord(context, UnknownType)
}

func WordsFilter(words []*Word, firstCharacter string) []*Word {
	var targets []*Word

	for _, word := range words {
		if word.StartWith(firstCharacter) {
			targets = append(targets, word)
		}
	}
	return targets
}
