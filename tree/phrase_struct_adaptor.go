package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type PhraseStructAdaptorParam struct {
	Index func([]Phrase) concept.Index
	Size  int
	Types string
	From  string
}

type PhraseStructAdaptor struct {
	children []Phrase
	param    *PhraseStructAdaptorParam
}

func (p *PhraseStructAdaptor) Index() concept.Index {
	return p.param.Index(p.children)
}

func (p *PhraseStructAdaptor) Types() string {
	return p.param.Types
}

func (p *PhraseStructAdaptor) Size() int {
	return p.param.Size
}

func (p *PhraseStructAdaptor) Copy() Phrase {
	substitute := NewPhraseStructAdaptor(p.param)
	for index, child := range p.children {
		substitute.SetChild(index, child.Copy())
	}
	return substitute
}

func (p *PhraseStructAdaptor) GetContent() *Vocabulary {
	return nil
}

func (p *PhraseStructAdaptor) GetChild(index int) Phrase {
	if index < 0 || index >= p.param.Size {
		return nil
	}
	return p.children[index]
}

func (p *PhraseStructAdaptor) SetChild(index int, child Phrase) Phrase {
	if index < 0 || index >= p.param.Size {
		panic("error index when set child")
	}
	p.children[index] = child
	return p
}

func (p *PhraseStructAdaptor) ToString() string {
	return p.ToStringOffset(0)
}

func (p *PhraseStructAdaptor) ToStringOffset(index int) string {
	var space = ""
	for i := 0; i < index; i++ {
		space += "\t"
	}
	info := fmt.Sprintf("%v%v (\n", space, p.param.Types)
	for i := 0; i < len(p.children); i++ {
		info += p.GetChild(i).ToStringOffset(index + 1)
	}
	info = fmt.Sprintf("%v%v)\n", info, space)
	return info
}

func (p *PhraseStructAdaptor) From() string {
	return p.param.From
}

func NewPhraseStructAdaptor(param *PhraseStructAdaptorParam) *PhraseStructAdaptor {
	return &PhraseStructAdaptor{
		param:    param,
		children: make([]Phrase, param.Size),
	}
}
