package grammar

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

const (
	TypeToken = 0
)

type PhraseToken struct {
	token *lexer.Token
}

func (p *PhraseToken) SetType(types int) {
	panic("complier.PhraseToken cannot set Type")
}

func (p *PhraseToken) Type() int {
	return p.token.Type()
}

func (p *PhraseToken) PhraseType() int {
	return TypeToken
}

func (p *PhraseToken) Size() int {
	return 0
}

func (p *PhraseToken) SetChild(int, Phrase) {
	panic("complier.PhraseToken cannot set child")
}

func (p *PhraseToken) GetChild(int) Phrase {
	panic("complier.PhraseToken cannot get child")
}

func (p *PhraseToken) GetToken() *lexer.Token {
	return p.token
}

func NewPhraseToken(token *lexer.Token) *PhraseToken {
	return &PhraseToken{
		token: token,
	}
}
