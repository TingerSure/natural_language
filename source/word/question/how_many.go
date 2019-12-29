package question

import (
	"github.com/TingerSure/natural_language/library/question"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	HowManyCharactor        = "多少"
	HowManyType      int    = word_types.Question
	HowManyName      string = "word.how_many"
)

type HowMany struct {
	adaptor.SourceAdaptor
}

func (p *HowMany) GetName() string {
	return HowManyName
}

func (p *HowMany) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(HowManyCharactor, HowManyType),
	}, sentence)
}

func (p *HowMany) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewConstIndex(question.HowMany)
					},
					Content: treasure,
					Types:   phrase_types.Question,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewHowMany() *HowMany {
	return (&HowMany{})
}
