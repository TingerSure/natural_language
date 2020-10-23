package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	Bool = tree.NewPhraseType("bool", []*tree.PhraseType{
		Any,
	})
)
