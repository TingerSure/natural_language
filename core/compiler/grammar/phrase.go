package grammar

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

type Phrase interface {
	PhraseType() int
	SetType(types int)
	Type() int
	Size() int
	SetChild(int, Phrase)
	GetChild(int) Phrase
	GetToken() *lexer.Token
}
