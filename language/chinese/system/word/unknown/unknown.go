package unknown

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
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
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.String.New(treasure.GetWord().GetContext()))
					},
					Content: treasure,
					Types:   phrase_type.Unknown,
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
