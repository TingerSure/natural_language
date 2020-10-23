package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	PronounPersonal = tree.NewPhraseType("pronoun.personal", []*tree.PhraseType{
		Noun,
	})
)
