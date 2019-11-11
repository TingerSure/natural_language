package verb

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	Is = "是"
)

const (
	setName string = "word.verb.set"
	setType int    = word_types.Verb
)

type Set struct {
	adaptor.SourceAdaptor
}

func (s *Set) GetName() string {
	return setName
}

func (s *Set) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter([]*tree.Word{
		tree.NewWord(Is, setType),
	}, sentence)
}
func (s *Set) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == s
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return nil
				//TODO
			}, treasure, phrase_types.Action, s.GetName())
		}, s.GetName()),
	}
}

func NewSet() *Set {
	return (&Set{})
}
