package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	Noun = tree.NewPhraseType("noun", []*tree.PhraseType{
		Any,
	})
)
