package source

import (
	"github.com/TingerSure/natural_language/word"
)

type Source interface {
	GetName() string
	GetWords(firstCharacter string) []*word.Word
}
