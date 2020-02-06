package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	Question = tree.NewPhraseType("question", []*tree.PhraseType{
		Any,
	})
)
