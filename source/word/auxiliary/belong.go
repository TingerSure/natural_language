package auxiliary

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	belongAuxiliaryName string = "word.auxiliary.belong"
	belongType          int    = word_types.AuxiliaryBelong
)

const (
	BelongTo string = "的"
)

type Belong struct {
	adaptor.SourceAdaptor
}

func (p *Belong) GetName() string {
	return belongAuxiliaryName
}

func (p *Belong) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(BelongTo, belongType),
	}, sentence)
}

func (p *Belong) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return nil
						//TODO
					},
					Content: treasure,
					Types:   phrase_types.AuxiliaryBelong,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewBelong() *Belong {
	return (&Belong{})
}
