package auxiliary

import (
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

func (p *Belong) GetWords(firstCharacter string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(BelongTo, belongType),
	}, firstCharacter)
}

func (p *Belong) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) tree.Phrase {
			if treasure.GetSource() != p {
				return nil
			}
			return tree.NewPhraseVocabularyAdaptor(treasure, phrase_types.AuxiliaryBelong)
		}, p.GetName()),
	}
}

const (
	targetFromTargetBelongUnknownSize = 3
)

func (p *Belong) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(targetFromTargetBelongUnknownSize, phrase_types.Target)
		}, targetFromTargetBelongUnknownSize, []string{
			phrase_types.Target,
			phrase_types.AuxiliaryBelong,
			phrase_types.Target,
		}, p.GetName()),
	}
}

func NewBelong() *Belong {
	return (&Belong{})
}
