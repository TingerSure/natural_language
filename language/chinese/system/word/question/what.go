package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
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
	*adaptor.SourceAdaptor
	WhatFunc concept.Function
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
						return libs.Sandbox.Index.ConstIndex.New(p.WhatFunc)
					},
					Content: treasure,
					Types:   phrase_type.Question,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewWhat(param *adaptor.SourceAdaptorParam) *What {
	instance := (&What{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	page := instance.Libs.GetLibraryPage("system", "question")
	instance.WhatFunc = page.GetFunction(libs.Sandbox.Variable.String.New("What"))

	return instance
}
