package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	HowManyCharactor        = "多少"
	HowManyName      string = "word.how_many"
)

type HowMany struct {
	adaptor.SourceAdaptor
	libs              *tree.LibraryManager
	libHowManyContent string
	libHowManyFunc    concept.Function
	HowManyFunc       *variable.Function
}

func (p *HowMany) init() {
	p.HowManyFunc = variable.NewFunction(nil)
	p.HowManyFunc.AddParamName(QuestionParam)
	p.HowManyFunc.Body().AddStep(
		expression.NewReturn(
			QuestionResult,
			expression.NewParamGet(
				expression.NewCall(
					index.NewConstIndex(p.libHowManyFunc),
					expression.NewNewParamWithInit(map[string]concept.Index{
						p.libHowManyContent: index.NewBubbleIndex(QuestionParam),
					}),
				),
				p.libHowManyContent,
			),
		),
	)
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
						return index.NewConstIndex(p.HowManyFunc)
					},
					Content: treasure,
					Types:   phrase_type.Question,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewHowMany(libs *tree.LibraryManager) *HowMany {
	page := libs.GetLibraryPage("system", "question")
	return (&HowMany{
		libs:              libs,
		libHowManyContent: page.GetConst("HowManyContent"),
		libHowManyFunc:    page.GetFunction("HowManyFunc"),
	})
}
