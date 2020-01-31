package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/core/tree/phrase_types"

	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/library/question"
)

const (
	HowManyCharactor = "多少"

	HowManyName string = "word.how_many"
)

var (
	HowManyFunc *variable.Function = nil
)

func init() {
	HowManyFunc = variable.NewFunction(nil)
	HowManyFunc.AddParamName(QuestionParam)
	HowManyFunc.Body().AddStep(
		expression.NewReturn(
			QuestionResult,
			expression.NewParamGet(
				expression.NewCall(
					index.NewConstIndex(question.HowMany),
					expression.NewNewParamWithInit(map[string]concept.Index{
						question.HowManyContent: index.NewBubbleIndex(QuestionParam),
					}),
				),
				question.HowManyContent,
			),
		),
	)

}

type HowMany struct {
	adaptor.SourceAdaptor
}

func (p *HowMany) GetName() string {
	return HowManyName
}

func (p *HowMany) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(HowManyCharactor),
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
						return index.NewConstIndex(HowManyFunc)
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
