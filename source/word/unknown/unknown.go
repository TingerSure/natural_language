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
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetWord().GetTypes() == word_types.Unknown
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return nil
				//TODO
			}, treasure, phrase_types.Unknown, p.GetName())
		}, p.GetName()),
	}
}

func NewUnknown() *Unknown {
	return (&Unknown{})
}
