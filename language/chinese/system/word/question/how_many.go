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
	parent            *Question
	libs              *tree.LibraryManager
	libHowManyContent concept.String
	libHowManyFunc    concept.Function
	HowManyFunc       *variable.Function
}

func (p *HowMany) init() *HowMany {
	p.HowManyFunc = variable.NewFunction(nil)
	p.HowManyFunc.AddParamName(p.parent.QuestionParam)
	p.HowManyFunc.Body().AddStep(
		expression.NewReturn(
			p.parent.QuestionResult,
			expression.NewParamGet(
				expression.NewCall(
					index.NewConstIndex(p.libHowManyFunc),
					expression.NewNewParamWithInit(map[concept.String]concept.Index{
						p.libHowManyContent: index.NewBubbleIndex(p.parent.QuestionParam),
					}),
				),
				p.libHowManyContent,
			),
		),
	)
	return p
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

func NewHowMany(libs *tree.LibraryManager, parent *Question) *HowMany {
	page := libs.GetLibraryPage("system", "question")
	return (&HowMany{
		libs:              libs,
		parent:            parent,
		libHowManyContent: page.GetConst(variable.NewString("HowManyContent")),
		libHowManyFunc:    page.GetFunction(variable.NewString("HowMany")),
	}).init()
}
