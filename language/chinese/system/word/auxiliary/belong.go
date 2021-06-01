package auxiliary

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	belongAuxiliaryName string = "word.auxiliary.belong"
)

const (
	BelongTo string = "的"
)

type Belong struct {
	*adaptor.SourceAdaptor
	instances []*tree.Vocabulary
}

func (p *Belong) GetName() string {
	return belongAuxiliaryName
}

func (p *Belong) GetWords(sentence string) []*tree.Vocabulary {
	return tree.VocabularysFilter(p.instances, sentence)
}

func (p *Belong) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
					Index: func() concept.Pipe {
						return p.Libs.Sandbox.Index.ConstIndex.New(p.Libs.Sandbox.Variable.String.New(BelongTo))
					},
					Content: treasure,
					Types:   phrase_type.AuxiliaryBelongName,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewBelong(param *adaptor.SourceAdaptorParam) *Belong {
	belong := (&Belong{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})

	belong.instances = []*tree.Vocabulary{
		tree.NewVocabulary(BelongTo, belong),
	}

	return belong
}
