package unknown

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	unknownName string = "word.unknown"
)

type Unknown struct {
	*adaptor.SourceAdaptor
}

func (p *Unknown) GetName() string {
	return unknownName
}

func (p *Unknown) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return nl_interface.IsNil(treasure.GetSource())
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
					Index: func() concept.Index {
						return p.Libs.Sandbox.Index.ConstIndex.New(p.Libs.Sandbox.Variable.String.New(treasure.GetContext()))
					},
					Content: treasure,
					Types:   phrase_type.UnknownName,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewUnknown(param *adaptor.SourceAdaptorParam) *Unknown {
	return (&Unknown{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
