package grammar

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

const (
	PhraseTypeToken = 0
)

type PhraseToken struct {
	token *lexer.Token
}

func (p *PhraseToken) Type() int {
	return p.token.Type()
}

func (p *PhraseToken) PhraseType() int {
	return PhraseTypeToken
}

func (p *PhraseToken) Size() int {
	return 0
}

func (p *PhraseToken) AddChild(children ...Phrase) {
	panic("complier.PhraseToken cannot add child")
}

func (p *PhraseToken) SetChild(int, Phrase) {
	panic("complier.PhraseToken cannot set child")
}

func (p *PhraseToken) GetChild(int) Phrase {
	panic("complier.PhraseToken cannot get child")
}

func (p *PhraseToken) SetStartLine(*lexer.Line) {
	//Do nothing
}

func (p *PhraseToken) GetLine() *lexer.Line {
	return p.token.Line()
}

func (p *PhraseToken) GetToken() *lexer.Token {
	return p.token
}

func (p *PhraseToken) GetRule() *Rule {
	return nil
}

func (p *PhraseToken) ToString(prefix string) string {
	return fmt.Sprintf("(%v) <%v>", p.GetToken().Value(), p.GetToken().Name())
}

func NewPhraseToken(token *lexer.Token) *PhraseToken {
	return &PhraseToken{
		token: token,
	}
}
