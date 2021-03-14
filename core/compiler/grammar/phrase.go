package grammar

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

type Phrase interface {
	PhraseType() int
	Type() int
	Size() int
	AddChild(children ...Phrase)
	SetChild(int, Phrase)
	GetChild(int) Phrase
	GetToken() *lexer.Token
	GetRule() *Rule
	ToString(string) string
}
