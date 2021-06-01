package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type PhrasePriority struct {
	values []Phrase
}

func (p *PhrasePriority) Index() concept.Pipe {
	return p.values[0].Index()
}

func (p *PhrasePriority) Types() string {
	return p.values[0].Types()
}

func (p *PhrasePriority) SetTypes(types string) {
	p.values[0].SetTypes(types)
}

func (p *PhrasePriority) Copy() Phrase {
	return NewPhrasePriority(p.values)
}

func (p *PhrasePriority) Size() int {
	return p.values[0].Size()
}

func (p *PhrasePriority) ContentSize() int {
	return p.values[0].ContentSize()
}

func (p *PhrasePriority) GetContent() *Vocabulary {
	return p.values[0].GetContent()
}

func (p *PhrasePriority) GetChild(index int) Phrase {
	return p.values[0].GetChild(index)
}

func (p *PhrasePriority) SetChild(index int, child Phrase) Phrase {
	return p.values[0].SetChild(index, child)
}

func (p *PhrasePriority) ToContent() string {
	subContents := []string{}
	for _, value := range p.values {
		subContents = append(subContents, value.ToContent())
	}
	return fmt.Sprintf("[%v]", strings.Join(subContents, " , "))
}

func (p *PhrasePriority) ToString() string {
	return p.ToStringOffset(0)
}

func (p *PhrasePriority) ToStringOffset(index int) string {
	var space = strings.Repeat("\t", index)
	info := fmt.Sprintf("%v%v [\n", space, p.Types())
	for i := 0; i < p.ValueSize(); i++ {
		info += p.GetValue(i).ToStringOffset(index + 1)
	}
	return fmt.Sprintf("%v%v]\n", info, space)
}

func (p *PhrasePriority) From() string {
	return p.values[0].From()
}

func (p *PhrasePriority) GetValue(index int) Phrase {
	return p.values[index]
}

func (p *PhrasePriority) SetValue(index int, value Phrase) {
	p.values[index] = value
}

func (p *PhrasePriority) SetValues(values []Phrase) {
	p.values = values
}

func (p *PhrasePriority) ValueSize() int {
	return len(p.values)
}

func (p *PhrasePriority) AddValue(value Phrase) {
	p.values = append(p.values, value)
}

func (p *PhrasePriority) AllValues() []Phrase {
	return p.values
}

func (p *PhrasePriority) RemoveValue(index int) {
	p.values = append(p.values[:index], p.values[index+1:]...)
}

func (p *PhrasePriority) HasPriority() bool {
	return true
}

func (p *PhrasePriority) DependencyCheckValue() Phrase {
	return p
}

func NewPhrasePriority(values []Phrase) *PhrasePriority {
	instance := &PhrasePriority{
		values: make([]Phrase, len(values)),
	}
	copy(instance.values, values)
	return instance
}
