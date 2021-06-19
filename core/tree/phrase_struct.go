package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type PhraseStructParam struct {
	Index        func([]Phrase) concept.Pipe
	Size         int
	Types        string
	DynamicTypes func([]Phrase) string
	From         string
}

type PhraseStruct struct {
	contentSize int
	children    []Phrase
	param       *PhraseStructParam
	types       string
}

func (p *PhraseStruct) Index() concept.Pipe {
	return p.param.Index(p.children)
}

func (p *PhraseStruct) Types() string {
	if p.types != "" {
		return p.types
	}
	if p.param.DynamicTypes != nil {
		p.types = p.param.DynamicTypes(p.children)
		return p.types
	}
	return p.param.Types
}

func (p *PhraseStruct) SetTypes(types string) {
	p.types = types
}

func (p *PhraseStruct) Size() int {
	return p.param.Size
}

func (p *PhraseStruct) updateContentSize() {
	size := 0
	for _, child := range p.children {
		if !nl_interface.IsNil(child) {
			size += child.ContentSize()
		}
	}
	p.contentSize = size
}

func (p *PhraseStruct) ContentSize() int {
	return p.contentSize
}

func (p *PhraseStruct) Copy() Phrase {
	substitute := NewPhraseStruct(p.param)
	for index, child := range p.children {
		substitute.SetChild(index, child.Copy())
	}
	return substitute
}

func (p *PhraseStruct) GetContent() string {
	return ""
}

func (p *PhraseStruct) GetChild(index int) Phrase {
	if index < 0 || index >= p.param.Size {
		return nil
	}
	return p.children[index]
}

func (p *PhraseStruct) SetChild(index int, child Phrase) Phrase {
	if index < 0 || index >= p.param.Size {
		panic("error index when set child")
	}
	p.children[index] = child
	p.updateContentSize()
	return p
}

func (p *PhraseStruct) ToContent() string {
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

func (p *PhraseStruct) ToString() string {
	return p.ToStringOffset(0)
}

func (p *PhraseStruct) ToStringOffset(index int) string {
	var space = strings.Repeat("\t", index)
	info := fmt.Sprintf("%v%v (\n", space, p.Types())
	for i := 0; i < len(p.children); i++ {
		info += p.GetChild(i).ToStringOffset(index + 1)
	}
	info = fmt.Sprintf("%v%v)\n", info, space)
	return info
}

func (p *PhraseStruct) From() string {
	return p.param.From
}

func (p *PhraseStruct) HasPriority() bool {
	for _, child := range p.children {
		if child.HasPriority() {
			return true
		}
	}
	return false
}

func (p *PhraseStruct) DependencyCheckValue() Phrase {
	return p
}

func NewPhraseStruct(param *PhraseStructParam) *PhraseStruct {
	return &PhraseStruct{
		param:       param,
		contentSize: 0,
		children:    make([]Phrase, param.Size),
	}
}
