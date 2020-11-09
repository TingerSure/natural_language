package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type PhraseStructAdaptorParam struct {
	Index        func([]Phrase) concept.Index
	Size         int
	Types        string
	DynamicTypes func([]Phrase) string
	From         string
}

type PhraseStructAdaptor struct {
	contentSize int
	children    []Phrase
	param       *PhraseStructAdaptorParam
	types       string
}

func (p *PhraseStructAdaptor) Index() concept.Index {
	return p.param.Index(p.children)
}

func (p *PhraseStructAdaptor) Types() string {
	if p.types != "" {
		return p.types
	}
	if p.param.DynamicTypes != nil {
		p.types = p.param.DynamicTypes(p.children)
		return p.types
	}
	return p.param.Types
}

func (p *PhraseStructAdaptor) Size() int {
	return p.param.Size
}

func (p *PhraseStructAdaptor) updateContentSize() {
	size := 0
	for _, child := range p.children {
		if !nl_interface.IsNil(child) {
			size += child.ContentSize()
		}
	}
	p.contentSize = size
}

func (p *PhraseStructAdaptor) ContentSize() int {
	return p.contentSize
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
	p.updateContentSize()
	return p
}

func (p *PhraseStructAdaptor) ToContent() string {
	subContents := []string{}
	for _, child := range p.children {
		if nl_interface.IsNil(child) {
			subContents = append(subContents, "null")
			continue
		}
		subContents = append(subContents, child.ToContent())
	}

	return fmt.Sprintf("(%v)", strings.Join(subContents, " "))
}

func (p *PhraseStructAdaptor) ToString() string {
	return p.ToStringOffset(0)
}

func (p *PhraseStructAdaptor) ToStringOffset(index int) string {
	var space = strings.Repeat("\t", index)
	info := fmt.Sprintf("%v%v (\n", space, p.Types())
	for i := 0; i < len(p.children); i++ {
		info += p.GetChild(i).ToStringOffset(index + 1)
	}
	info = fmt.Sprintf("%v%v)\n", info, space)
	return info
}

func (p *PhraseStructAdaptor) From() string {
	return p.param.From
}

func (p *PhraseStructAdaptor) HasPriority() bool {
	for _, child := range p.children {
		if child.HasPriority() {
			return true
		}
	}
	return false
}

func NewPhraseStructAdaptor(param *PhraseStructAdaptorParam) *PhraseStructAdaptor {
	return &PhraseStructAdaptor{
		param:       param,
		contentSize: 0,
		children:    make([]Phrase, param.Size),
	}
}
