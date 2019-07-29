package pronoun

import (
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	targetPronounName string = "system.pronoun.target"
	targetType        int    = word_types.Pronoun
)

const (
	He  string = "他"
	She string = "她"
	It  string = "它"
	You string = "你"
	I   string = "我"
)

var (
	targetPronounWords []*tree.Word = []*tree.Word{
		tree.NewWord(He, targetType),
		tree.NewWord(She, targetType),
		tree.NewWord(It, targetType),
		tree.NewWord(You, targetType),
		tree.NewWord(I, targetType),
	}
)

type Target struct {
}

func (p *Target) GetName() string {
	return targetPronounName
}

func (p *Target) GetWords(firstCharacter string) []*tree.Word {
	return tree.WordsFilter(targetPronounWords, firstCharacter)
}

func (p *Target) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) tree.Phrase {
			if treasure.GetSource() != p {
				return nil
			}
			return tree.NewPhraseVocabularyAdaptor(treasure, phrase_types.Target)
		}, p.GetName()),
	}
}

const (
	targetFromTwoTargetsSize    = 2
	targetFromTargetUnknownSize = 2
	targetFromUnknownTargetSize = 2
)

func (p *Target) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{
		tree.NewStructRule(func(treasures []tree.Phrase) tree.Phrase {
			first := treasures[0]
			second := treasures[1]
			if first.Types() == phrase_types.Target &&
				second.Types() == phrase_types.Target {
				return tree.NewPhraseStructAdaptor(targetFromTwoTargetsSize, phrase_types.Target).SetChild(0, first).SetChild(1, second)
			}
			return nil
		}, targetFromTwoTargetsSize, p.GetName()),
		tree.NewStructRule(func(treasures []tree.Phrase) tree.Phrase {
			first := treasures[0]
			second := treasures[1]
			if first.Types() == phrase_types.Target &&
				second.Types() == phrase_types.Unknown {
				return tree.NewPhraseStructAdaptor(targetFromTargetUnknownSize, phrase_types.Target).SetChild(0, first).SetChild(1, second)
			}
			return nil
		}, targetFromTargetUnknownSize, p.GetName()),
		tree.NewStructRule(func(treasures []tree.Phrase) tree.Phrase {
			first := treasures[0]
			second := treasures[1]
			if first.Types() == phrase_types.Unknown &&
				second.Types() == phrase_types.Target {
				return tree.NewPhraseStructAdaptor(targetFromUnknownTargetSize, phrase_types.Target).SetChild(0, first).SetChild(1, second)
			}
			return nil
		}, targetFromUnknownTargetSize, p.GetName()),
	}
}

func NewTarget() *Target {
	return (&Target{})
}
