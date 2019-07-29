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
	targetFromTargetBelongTargetSize  = 3
)

func (p *Belong) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{
		tree.NewStructRule(func(treasures []tree.Phrase) tree.Phrase {
			target := treasures[0]
			belong := treasures[1]
			unknown := treasures[2]
			if target.Types() == phrase_types.Target &&
				belong.Types() == phrase_types.AuxiliaryBelong &&
				unknown.Types() == phrase_types.Unknown {
				return tree.NewPhraseStructAdaptor(targetFromTargetBelongUnknownSize, phrase_types.Target).SetChild(0, target).SetChild(1, belong).SetChild(2, unknown)
			}
			return nil
		}, targetFromTargetBelongUnknownSize, p.GetName()),
		tree.NewStructRule(func(treasures []tree.Phrase) tree.Phrase {
			first := treasures[0]
			belong := treasures[1]
			second := treasures[2]
			if first.Types() == phrase_types.Target &&
				belong.Types() == phrase_types.AuxiliaryBelong &&
				second.Types() == phrase_types.Target {
				return tree.NewPhraseStructAdaptor(targetFromTargetBelongUnknownSize, phrase_types.Target).SetChild(0, first).SetChild(1, belong).SetChild(2, second)
			}
			return nil
		}, targetFromTargetBelongUnknownSize, p.GetName()),
	}
}

func NewBelong() *Belong {
	return (&Belong{})
}
