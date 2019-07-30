package tree

import (
	"fmt"
)

type PhraseStructAdaptor struct {
	size     int
	children []Phrase
	types    string
}

func (p *PhraseStructAdaptor) Types() string {
	return p.types
}

func (p *PhraseStructAdaptor) Size() int {
	return p.size
}

func (p *PhraseStructAdaptor) Copy() Phrase {
	substitute := NewPhraseStructAdaptor(p.size, p.types)
	for index, child := range p.children {
		substitute.SetChild(index, child.Copy())
	}
	return substitute
}

func (p *PhraseStructAdaptor) GetContent() *Vocabulary {
	return nil
}

func (p *PhraseStructAdaptor) GetChild(index int) Phrase {
	if index < 0 || index >= p.size {
		return nil
	}
	return p.children[index]
}

func (p *PhraseStructAdaptor) SetChild(index int, child Phrase) Phrase {
	if index < 0 || index >= p.size {
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
	info := fmt.Sprintf("%v%v (\n", space, p.types)
	for i := 0; i < len(p.children); i++ {
		info += p.GetChild(i).ToStringOffset(index + 1)
	}
	info = fmt.Sprintf("%v%v)\n", info, space)
	return info
}

func NewPhraseStructAdaptor(size int, types string) *PhraseStructAdaptor {
	return &PhraseStructAdaptor{
		size:     size,
		types:    types,
		children: make([]Phrase, size),
	}
}
