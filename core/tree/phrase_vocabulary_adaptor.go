package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type PhraseVocabularyAdaptorParam struct {
	Index   func() concept.Index
	Content *Vocabulary
	Types   *PhraseType
	From    string
}

type PhraseVocabularyAdaptor struct {
	param *PhraseVocabularyAdaptorParam
}

func (p *PhraseVocabularyAdaptor) Index() concept.Index {
	return p.param.Index()
}

func (p *PhraseVocabularyAdaptor) Types() *PhraseType {
	return p.param.Types
}

func (p *PhraseVocabularyAdaptor) Copy() Phrase {
	return NewPhraseVocabularyAdaptor(p.param)
}

func (p *PhraseVocabularyAdaptor) Size() int {
	return 0
}

func (p *PhraseVocabularyAdaptor) ContentSize() int {
	return p.param.Content.Len()
}

func (p *PhraseVocabularyAdaptor) GetContent() *Vocabulary {
	return p.param.Content
}

func (p *PhraseVocabularyAdaptor) GetChild(index int) Phrase {
	return nil
}

func (p *PhraseVocabularyAdaptor) SetChild(index int, child Phrase) Phrase {
	panic("This phrase can not set child")
	return p
}

func (p *PhraseVocabularyAdaptor) ToString() string {
	return p.ToStringOffset(0)
}

func (p *PhraseVocabularyAdaptor) ToStringOffset(index int) string {
	var space = ""
	for i := 0; i < index; i++ {
		space += "\t"
	}
	return fmt.Sprintf("%v%v ( %v )\n", space, p.param.Types.Name(), p.param.Content.ToString())
}

func (p *PhraseVocabularyAdaptor) From() string {
	return p.param.From
}

func (p *PhraseVocabularyAdaptor) HasPriority() bool {
	return false
}

func NewPhraseVocabularyAdaptor(param *PhraseVocabularyAdaptorParam) *PhraseVocabularyAdaptor {
	return &PhraseVocabularyAdaptor{
		param: param,
	}
}
