package tree

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_string"
)

type Word struct {
	context string
}

func (w *Word) GetContext() string {
	return w.context
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

func NewWord(context string) *Word {
	return (&Word{
		context: context,
	})

}

func WordsFilter(words []*Word, sentence string) []*Word {
	var targets []*Word

	for _, word := range words {
		if word.StartFor(sentence) {
			targets = append(targets, word)
		}
	}
	return targets
}
