package unknown

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	unknownName string = "system.unknown"
)

type Unknown struct {
}

func (p *Unknown) GetName() string {
	return unknownName
}

func (p *Unknown) GetWords(firstCharacter string) []*tree.Word {
	return nil
}

func (p *Unknown) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetWord().GetTypes() == word_types.Unknown
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return nil
				//TODO
			}, treasure, phrase_types.Unknown)
		}, p.GetName()),
	}
}

func (p *Unknown) GetStructRules() []*tree.StructRule {
	return nil
}

func NewUnknown() *Unknown {
	return (&Unknown{})
}
