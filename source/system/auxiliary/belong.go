package auxiliary

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	belongAuxiliaryName string = "system.auxiliary.belong"
	belongType          int    = word_types.AuxiliaryBelong
)

const (
	BelongTo string = "的"
)

type Belong struct {
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
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return nil
				//TODO
			}, treasure, phrase_types.AuxiliaryBelong)
		}, p.GetName()),
	}
}

func (p *Belong) GetStructRules() []*tree.StructRule {
	return nil
}

func NewBelong() *Belong {
	return (&Belong{})
}
