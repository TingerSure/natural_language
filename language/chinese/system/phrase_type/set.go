package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	Set = tree.NewPhraseType("set", []*tree.PhraseType{
		Any,
	})
)
