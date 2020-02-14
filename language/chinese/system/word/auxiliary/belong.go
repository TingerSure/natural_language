package auxiliary

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"

	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	belongAuxiliaryName string = "word.auxiliary.belong"
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
		tree.NewWord(BelongTo),
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
						return index.NewConstIndex(variable.NewString(treasure.GetWord().GetContext()))
					},
					Content: treasure,
					Types:   phrase_type.AuxiliaryBelong,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewBelong(libs *tree.LibraryManager) *Belong {
	return (&Belong{})
}
