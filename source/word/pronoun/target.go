package pronoun

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	targetPronounName string = "word.pronoun.target"
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
	adaptor.SourceAdaptor
}

func (p *Target) GetName() string {
	return targetPronounName
}

func (p *Target) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(targetPronounWords, sentence)
}

func (p *Target) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return nil
				//TODO
			}, treasure, phrase_types.Target, p.GetName())
		}, p.GetName()),
	}
}

func NewTarget() *Target {
	return (&Target{})
}
