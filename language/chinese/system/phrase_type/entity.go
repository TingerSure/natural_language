package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	Entity = tree.NewPhraseType("entity", []*tree.PhraseType{
		Any,
	})
)
