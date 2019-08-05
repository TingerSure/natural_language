package verb

import (
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	Is = "是"
)

const (
	setName string = "system.verb.set"
	setType int    = word_types.Verb
)

type Set struct {
}

func (s *Set) GetName() string {
	return setName
}

func (s *Set) GetWords(firstCharacter string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(Is, setType),
	}, firstCharacter)
}
func (s *Set) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) tree.Phrase {
			if treasure.GetSource() != s {
				return nil
			}
			return tree.NewPhraseVocabularyAdaptor(treasure, phrase_types.Action)
		}, s.GetName()),
	}
}

func (s *Set) GetStructRules() []*tree.StructRule {
	return nil
}
func NewSet() *Set {
	return (&Set{})
}
