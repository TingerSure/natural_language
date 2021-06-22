package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
	"unicode/utf8"
)

type PhraseVocabularyParam struct {
	Index   func() concept.Function
	Content string
	Types   string
	From    string
}

type PhraseVocabulary struct {
	param *PhraseVocabularyParam
	types string
}

func (p *PhraseVocabulary) Index() concept.Function {
	return p.param.Index()
}

func (p *PhraseVocabulary) Types() string {
	if p.types != "" {
		return p.types
	}
	return p.param.Types
}

func (p *PhraseVocabulary) SetTypes(types string) {
	p.types = types
}

func (p *PhraseVocabulary) Copy() Phrase {
	return NewPhraseVocabulary(p.param)
}

func (p *PhraseVocabulary) Size() int {
	return 0
}

func (p *PhraseVocabulary) ContentSize() int {
	return utf8.RuneCountInString(p.param.Content)
}

func (p *PhraseVocabulary) GetContent() string {
	return p.param.Content
}

func (p *PhraseVocabulary) GetChild(index int) Phrase {
	return nil
}

func (p *PhraseVocabulary) SetChild(index int, child Phrase) Phrase {
	panic("This phrase can not set child")
	return p
}

func (p *PhraseVocabulary) ToString() string {
	return p.ToStringOffset(0)
}

func (p *PhraseVocabulary) ToContent() string {
	content := p.param.Content
	content = strings.Replace(content, "\\", "\\\\", -1)
	content = strings.Replace(content, ",", "\\,", -1)
	content = strings.Replace(content, "(", "\\(", -1)
	content = strings.Replace(content, ")", "\\)", -1)
	content = strings.Replace(content, "[", "\\[", -1)
	content = strings.Replace(content, "]", "\\]", -1)
	return content
}

func (p *PhraseVocabulary) ToStringOffset(index int) string {
	return fmt.Sprintf("%v%v ( %v )\n", strings.Repeat("\t", index), p.param.Types, p.param.Content)
}

func (p *PhraseVocabulary) From() string {
	return p.param.From
}

func (p *PhraseVocabulary) HasPriority() bool {
	return false
}

func (p *PhraseVocabulary) DependencyCheckValue() Phrase {
	return p
}

func NewPhraseVocabulary(param *PhraseVocabularyParam) *PhraseVocabulary {
	return &PhraseVocabulary{
		param: param,
	}
}
