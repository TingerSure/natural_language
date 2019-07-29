package pronoun

import (
	"github.com/TingerSure/natural_language/tree"
)

const (
	targetPronounName string = "system.pronoun.target"
	targetType        int    = tree.Pronoun
)

const (
	targetRuleType string = "rule.target"
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
			return tree.NewPhraseVocabularyAdaptor(treasure)
		}, targetPronounName),
	}
}

func (p *Target) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{}
}

func NewTarget() *Target {
	return (&Target{})
}
