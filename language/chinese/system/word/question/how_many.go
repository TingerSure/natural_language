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
	HowManyCharactor        = "多少"
	HowManyName      string = "word.how_many"
)

type HowMany struct {
	*adaptor.SourceAdaptor
	HowManyFunc concept.Function
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
						return libs.Sandbox.Index.ConstIndex.New(p.HowManyFunc)
					},
					Content: treasure,
					Types:   phrase_type.Question,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewHowMany(param *adaptor.SourceAdaptorParam) *HowMany {
	instance := (&HowMany{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	page := instance.Libs.GetLibraryPage("system", "question")
	instance.HowManyFunc = page.GetFunction(libs.Sandbox.Variable.String.New("HowMany"))

	return instance
}
