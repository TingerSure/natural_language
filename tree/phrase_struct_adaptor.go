package tree

import ()

type PhraseStructAdaptor struct {
	size     int
	children []Phrase
}

func (p *PhraseStructAdaptor) Size() int {
	return p.size
}

func (p *PhraseStructAdaptor) Copy() Phrase {
	substitute := NewPhraseStructAdaptor(p.size)
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

func (p *PhraseStructAdaptor) SetChild(index int, child Phrase) {
	if index < 0 || index >= p.size {
		panic("error index when set child")
	}
	p.children[index] = child
}

func (p *PhraseStructAdaptor) ToString() string {
	return p.ToStringOffset(0)
}

func (p *PhraseStructAdaptor) ToStringOffset(index int) string {
	var info = ""
	if index > 0 {
		for i := 0; i < index-1; i++ {
			info += "\t"
		}
		info += "|---"
	}
	for i := 0; i < len(p.children); i++ {
		info += p.GetChild(i).ToStringOffset(index + 1)
	}
	return info
}

func NewPhraseStructAdaptor(size int) *PhraseStructAdaptor {
	return &PhraseStructAdaptor{
		size:     size,
		children: make([]Phrase, size),
	}
}
