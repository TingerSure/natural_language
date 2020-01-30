package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/core/tree/phrase_types"
	"github.com/TingerSure/natural_language/core/tree/word_types"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/library/question"
)

const (
	WhatCharactor        = "什么"
	WhatType      int    = word_types.Question
	WhatName      string = "word.what"
)

var (
	WhatFunc *variable.Function = nil
)

func init() {
	WhatFunc = variable.NewFunction(nil)
	WhatFunc.AddParamName(QuestionParam)
	WhatFunc.Body().AddStep(
		expression.NewReturn(
			QuestionResult,
			expression.NewParamGet(
				expression.NewCall(
					index.NewConstIndex(question.What),
					expression.NewNewParamWithInit(map[string]concept.Index{
						question.WhatContent: index.NewBubbleIndex(QuestionParam),
					}),
				),
				question.WhatContent,
			),
		),
	)

}

type What struct {
	adaptor.SourceAdaptor
}

func (p *What) GetName() string {
	return WhatName
}

func (p *What) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(WhatCharactor, WhatType),
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
						return index.NewConstIndex(WhatFunc)
					},
					Content: treasure,
					Types:   phrase_types.Question,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewWhat() *What {
	return (&What{})
}
