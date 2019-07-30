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

const (
	sentenceFromTargetActionTarget = 3
)

func (s *Set) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{
		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(sentenceFromTargetActionTarget, phrase_types.Event)
		}, sentenceFromTargetActionTarget, []string{
			phrase_types.Target,
			phrase_types.Action,
			phrase_types.Target,
		}, s.GetName()),
	}
}
func NewSet() *Set {
	return (&Set{})
}
