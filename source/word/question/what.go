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
	WhatCharactor        = "什么"
	WhatType      int    = word_types.Question
	WhatName      string = "word.what"
)

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
						return index.NewConstIndex(question.What)
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
