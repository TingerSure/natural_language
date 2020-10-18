package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	PronounInterrogative = tree.NewPhraseType("pronoun.interrogative", []*tree.PhraseType{
		Any,
	})
)
