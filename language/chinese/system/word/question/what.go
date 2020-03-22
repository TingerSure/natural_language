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
	WhatCharactor        = "什么"
	WhatName      string = "word.what"
)

type What struct {
	adaptor.SourceAdaptor
	parent         *Question
	libs           *tree.LibraryManager
	libWhatContent concept.String
	libWhatFunc    concept.Function
	WhatFunc       *variable.Function
}

func (p *What) init() *What {
	p.WhatFunc = variable.NewFunction(nil)
	p.WhatFunc.AddParamName(p.parent.QuestionParam)
	p.WhatFunc.Body().AddStep(
		expression.NewReturn(
			p.parent.QuestionResult,
			expression.NewParamGet(
				expression.NewCall(
					index.NewConstIndex(p.libWhatFunc),
					expression.NewNewParamWithInit(map[concept.String]concept.Index{
						p.libWhatContent: index.NewBubbleIndex(p.parent.QuestionParam),
					}),
				),
				p.libWhatContent,
			),
		),
	)
	return p
}

func (p *What) GetName() string {
	return WhatName
}

func (p *What) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(WhatCharactor),
	}, sentence)
}

func (p *What) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewConstIndex(p.WhatFunc)
					},
					Content: treasure,
					Types:   phrase_type.Question,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewWhat(libs *tree.LibraryManager, parent *Question) *What {
	page := libs.GetLibraryPage("system", "question")
	return (&What{
		libs:           libs,
		parent:         parent,
		libWhatContent: page.GetConst(variable.NewString("WhatContent")),
		libWhatFunc:    page.GetFunction(variable.NewString("What")),
	}).init()
}
