package tree

import (
	"fmt"
)

type Phrase struct {
	types    string
	content  *Vocabulary
	size     int
	children []*Phrase
}

func (p *Phrase) Copy() *Phrase {
	substitute := NewPhrase(p.types, p.content, p.size)
	for index, child := range p.children {
		substitute.SetChild(index, child.Copy())
	}
	return substitute
}

func (p *Phrase) GetType() string {
	return p.types
}

func (p *Phrase) GetContent() *Vocabulary {
	return p.content
}

func (p *Phrase) GetChild(index int) *Phrase {
	if index < 0 || index >= p.size {
		return nil
	}
	return p.children[index]
}

func (p *Phrase) SetChild(index int, child *Phrase) {
	if index < 0 || index >= p.size {
		panic("error index when set child")
	}
	p.children[index] = child
}

func (p *Phrase) ToString() string {
	return p.ToStringOffset(0)
}

func (p *Phrase) ToStringOffset(index int) string {
	var info = ""
	if index > 0 {
		for i := 0; i < index-1; i++ {
			info += "\t"
		}
		info += "|---"
	}
	info = fmt.Sprintf("%v%v\n", info, p.GetType())
	for i := 0; i < len(p.children); i++ {
		info += p.GetChild(i).ToStringOffset(index + 1)
	}
	return info
}

func NewPhrase(types string, content *Vocabulary, size int) *Phrase {
	return &Phrase{
		types:    types,
		content:  content,
		size:     size,
		children: make([]*Phrase, size),
	}
}
