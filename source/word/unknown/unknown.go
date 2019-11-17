package unknown

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	unknownName string = "word.unknown"
)

type Unknown struct {
	adaptor.SourceAdaptor
}

func (p *Unknown) GetName() string {
	return unknownName
}

func (p *Unknown) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetWord().GetTypes() == word_types.Unknown
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return nil
						//TODO
					},
					Content: treasure,
					Types:   phrase_types.Unknown,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewUnknown() *Unknown {
	return (&Unknown{})
}
