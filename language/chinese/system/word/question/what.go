package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
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
	WhatFunc  concept.Function
	instances []*tree.Vocabulary
}

func (p *What) GetName() string {
	return WhatName
}

func (p *What) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(p.instances, sentence)
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
						return p.Libs.Sandbox.Index.ConstIndex.New(p.WhatFunc)
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
	instance.WhatFunc = page.GetFunction(instance.Libs.Sandbox.Variable.String.New("What"))
	instance.instances = []*tree.Vocabulary{
		tree.NewVocabulary(WhatCharactor, instance),
	}
	return instance
}
