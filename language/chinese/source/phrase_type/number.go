package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	Number = tree.NewPhraseType("number", []*tree.PhraseType{
		Any,
	})
)
