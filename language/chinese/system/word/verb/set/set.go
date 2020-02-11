package set

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/source/phrase_type"

	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
)

const (
	Is    = "是"
	Equal = "等于"

	SetName string = "word.verb.set"
)

var (
	SetCharactors = []*tree.Word{
		tree.NewWord(Is),
		tree.NewWord(Equal),
	}
)

type Set struct {
	adaptor.SourceAdaptor
}

func (s *Set) GetName() string {
	return SetName
}

func (s *Set) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(SetCharactors, sentence)
}
func (s *Set) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == s
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewConstIndex(variable.NewString(treasure.GetWord().GetContext()))
					},
					Content: treasure,
					Types:   phrase_type.Set,
					From:    s.GetName(),
				})
			}, From: s.GetName(),
		}),
	}
}

func NewSet() *Set {
	return (&Set{})
}
