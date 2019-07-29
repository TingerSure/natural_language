package tree

import (
	"github.com/TingerSure/natural_language/word"
)

type Source interface {
	GetName() string
	GetWords(firstCharacter string) []*word.Word
	GetVocabularyRules() []*VocabularyRule
	GetStructRules() []*StructRule
}
