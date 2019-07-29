package pronoun

import (
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/word"
)

const (
	targetPronounName string = "system.pronoun.target"
	targetType        int    = word.Pronoun
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
	targetPronounWords []*word.Word = []*word.Word{
		word.NewWord(He, targetType),
		word.NewWord(She, targetType),
		word.NewWord(It, targetType),
		word.NewWord(You, targetType),
		word.NewWord(I, targetType),
	}
)

type Target struct {
}

func (p *Target) GetName() string {
	return targetPronounName
}

func (p *Target) GetWords(firstCharacter string) []*word.Word {
	return word.WordsFilter(targetPronounWords, firstCharacter)
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
