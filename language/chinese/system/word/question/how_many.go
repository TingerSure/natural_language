package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
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
	instances   []*tree.Vocabulary
}

func (p *HowMany) GetName() string {
	return HowManyName
}

func (p *HowMany) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(p.instances, sentence)
}

func (p *HowMany) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
					Index: func() concept.Index {
						return p.Libs.Sandbox.Index.ConstIndex.New(p.HowManyFunc)
					},
					Content: treasure,
					Types:   phrase_type.PronounInterrogativeName,
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
	instance.HowManyFunc = page.GetFunction(instance.Libs.Sandbox.Variable.String.New("HowMany"))
	instance.instances = []*tree.Vocabulary{
		tree.NewVocabulary(HowManyCharactor, instance),
	}
	return instance
}
