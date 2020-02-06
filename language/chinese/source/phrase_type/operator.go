package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	Operator = tree.NewPhraseType("operator", []*tree.PhraseType{
		Any,
	})
)
