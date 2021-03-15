package grammar

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"strings"
)

type PhraseStruct struct {
	children []Phrase
	rule     *Rule
}

const (
	PhraseTypeStruct = 1
)

func (p *PhraseStruct) Type() int {
	return p.rule.GetResult().Type()
}

func (p *PhraseStruct) Size() int {
	return len(p.children)
}

func (p *PhraseStruct) PhraseType() int {
	return PhraseTypeStruct
}

func (p *PhraseStruct) AddChild(children ...Phrase) {
	p.children = append(p.children, children...)
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

func (p *PhraseStruct) ToString(prefix string) string {
	subPrefix := prefix + "\t"
	subs := []string{}
	for index := 0; index < p.Size(); index++ {
		child := p.GetChild(index)
		subs = append(subs, fmt.Sprintf("%v%v", subPrefix, child.ToString(subPrefix)))
	}
	return fmt.Sprintf("{%v\n%v\n%v\n%v}", p.rule.ToString(), subPrefix, strings.Join(subs, "\n"), prefix)
}

func NewPhraseStruct(rule *Rule) *PhraseStruct {
	return &PhraseStruct{
		rule: rule,
	}
}
