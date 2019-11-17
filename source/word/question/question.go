package question

import (
	"github.com/TingerSure/natural_language/library/std"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	WhatCharactor        = "什么"
	QuestionType  int    = word_types.Question
	QuestionName  string = "word.question"
)

type Question struct {
	adaptor.SourceAdaptor
}

func (p *Question) GetName() string {
	return QuestionName
}

func (p *Question) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(WhatCharactor, QuestionType),
	}, sentence)
}

func (p *Question) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewConstIndex(std.Print)
					},
					Content: treasure,
					Types:   phrase_types.Question,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewQuestion() *Question {
	return (&Question{})
}
