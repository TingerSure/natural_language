package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Phrase interface {
	Copy() Phrase
	Size() int
	Types() *PhraseType
	GetContent() *Vocabulary
	GetChild(index int) Phrase
	SetChild(index int, child Phrase) Phrase
	ToString() string
	ToStringOffset(index int) string
	Index() concept.Index
	From() string
}
