package phrase_type

import (
	"github.com/TingerSure/natural_language/core/tree"
)

var (
	BracketsLeft = tree.NewPhraseType("brackets.left", []*tree.PhraseType{
		Any,
	})

	BracketsRight = tree.NewPhraseType("brackets.right", []*tree.PhraseType{
		Any,
	})
)
