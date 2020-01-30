package brackets

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	RightCharactor        = ")"
	RightName      string = "word.brackets.right"
	RightType      int    = word_types.AuxiliaryBrackets
)

var (
	RightWord  []*tree.Word = []*tree.Word{tree.NewWord(RightCharactor, RightType)}
	RightIndex              = index.NewConstIndex(variable.NewString(RightCharactor))
)

type BracketsRight struct {
	adaptor.SourceAdaptor
}

func (s *BracketsRight) GetName() string {
	return RightName
}

func (s *BracketsRight) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(RightWord, sentence)
}
func (s *BracketsRight) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == s
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return RightIndex
					},
					Content: treasure,
					Types:   phrase_types.BracketsRight,
					From:    s.GetName(),
				})
			}, From: s.GetName(),
		}),
	}
}

func NewBracketsRight() *BracketsRight {
	return (&BracketsRight{})
}