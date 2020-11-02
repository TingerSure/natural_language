package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type PhrasePackageParam struct {
	Index func(Phrase) concept.Index
	Types string
	From  string
}

type PhrasePackage struct {
	value Phrase
	param *PhrasePackageParam
}

func (p *PhrasePackage) Index() concept.Index {
	return p.param.Index(p.value)
}

func (p *PhrasePackage) Types() string {
	return p.param.Types
}

func (p *PhrasePackage) Size() int {
	return p.value.Size()
}

func (p *PhrasePackage) ContentSize() int {
	return p.value.ContentSize()
}

func (p *PhrasePackage) Copy() Phrase {
	substitute := NewPhrasePackage(p.param)
	substitute.SetValue(p.value.Copy())
	return substitute
}

func (p *PhrasePackage) GetContent() *Vocabulary {
	return p.value.GetContent()
}

func (p *PhrasePackage) GetChild(index int) Phrase {
	return p.value.GetChild(index)
}

func (p *PhrasePackage) SetChild(index int, child Phrase) Phrase {
	p.value.SetChild(index, child)
	return p
}

func (p *PhrasePackage) ToContent() string {
	return p.value.ToContent()
}

func (p *PhrasePackage) ToString() string {
	return p.ToStringOffset(0)
}

func (p *PhrasePackage) ToStringOffset(index int) string {
	var space = strings.Repeat("\t", index)
	return fmt.Sprintf("%v%v (\n %v%v)\n", space, p.Types(), p.value.ToStringOffset(index+1), space)
}

func (p *PhrasePackage) From() string {
	return p.param.From
}

func (p *PhrasePackage) HasPriority() bool {
	return p.value.HasPriority()
}

func (p *PhrasePackage) SetValue(value Phrase) *PhrasePackage {
	p.value = value
	return p
}

func NewPhrasePackage(param *PhrasePackageParam) *PhrasePackage {
	return &PhrasePackage{
		param: param,
	}
}
