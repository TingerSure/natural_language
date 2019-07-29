package unknown

import (
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
	return []*tree.Word{}
}

func (p *Unknown) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) tree.Phrase {
			if treasure.GetWord().GetTypes() != word_types.Unknown {
				return nil
			}
			return tree.NewPhraseVocabularyAdaptor(treasure, phrase_types.Unknown)
		}, p.GetName()),
	}
}

func (p *Unknown) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{}
}

func NewUnknown() *Unknown {
	return (&Unknown{})
}
