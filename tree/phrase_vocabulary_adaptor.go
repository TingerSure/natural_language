package tree

import (
	"fmt"
)

type PhraseVocabularyAdaptor struct {
	types   string
	content *Vocabulary
}

func (p *PhraseVocabularyAdaptor) Types() string {
	return p.types
}

func (p *PhraseVocabularyAdaptor) Copy() Phrase {
	return NewPhraseVocabularyAdaptor(p.content, p.types)
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
	var info = ""
	if index > 0 {
		for i := 0; i < index-1; i++ {
			info += "\t"
		}
		info += "|---"
	}
	return fmt.Sprintf("%v%v\n", info, p.content.ToString())
}

func NewPhraseVocabularyAdaptor(content *Vocabulary, types string) *PhraseVocabularyAdaptor {
	return &PhraseVocabularyAdaptor{
		content: content,
		types:   types,
	}
}
