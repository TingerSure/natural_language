package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	Unknown = tree.NewPhraseType("unknown", []*tree.PhraseType{
		Any,
	})
)
