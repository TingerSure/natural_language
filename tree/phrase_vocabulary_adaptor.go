package tree

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type PhraseVocabularyAdaptor struct {
	types   string
	content *Vocabulary
	index   func() concept.Index
}

func (p *PhraseVocabularyAdaptor) Index() concept.Index {
	return p.index()
}

func (p *PhraseVocabularyAdaptor) Types() string {
	return p.types
}

func (p *PhraseVocabularyAdaptor) Copy() Phrase {
	return NewPhraseVocabularyAdaptor(p.index, p.content, p.types)
}
func (p *PhraseVocabularyAdaptor) Size() int {
	return 0
}
func (p *PhraseVocabularyAdaptor) GetContent() *Vocabulary {
	return p.content
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
	return fmt.Sprintf("%v%v ( %v )\n", space, p.types, p.content.ToString())
}

func NewPhraseVocabularyAdaptor(index func() concept.Index, content *Vocabulary, types string) *PhraseVocabularyAdaptor {
	return &PhraseVocabularyAdaptor{
		content: content,
		index:   index,
		types:   types,
	}
}
