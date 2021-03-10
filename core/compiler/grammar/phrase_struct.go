package grammar

import (
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

type PhraseStruct struct {
	size     int
	children []Phrase
	rule     *Rule
}

const (
	TypeStruct = 1
)

func (p *PhraseStruct) Size() int {
	return p.size
}

func (p *PhraseStruct) PhraseType() int {
	return TypeStruct
}

func (p *PhraseStruct) SetChild(index int, child Phrase) {
	p.children[index] = child
}

func (p *PhraseStruct) GetChild(index int) Phrase {
	return p.children[index]
}

func (p *PhraseStruct) GetToken() *lexer.Token {
	return nil
}

func (p *PhraseStruct) GetRule() *Rule {
	return p.rule
}

func NewPhraseStruct(rule *Rule) *PhraseStruct {
	return &PhraseStruct{
		rule: rule,
	}
}
