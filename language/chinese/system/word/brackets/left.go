package brackets

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"

	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	LeftCharactor        = "("
	LeftName      string = "word.brackets.left"
)

var (
	LeftWord  []*tree.Word = []*tree.Word{tree.NewWord(LeftCharactor)}
	LeftIndex              = index.NewConstIndex(variable.NewString(LeftCharactor))
)

type BracketsLeft struct {
	adaptor.SourceAdaptor
}

func (s *BracketsLeft) GetName() string {
	return LeftName
}

func (s *BracketsLeft) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(LeftWord, sentence)
}
func (s *BracketsLeft) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == s
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return LeftIndex
					},
					Content: treasure,
					Types:   phrase_type.BracketsLeft,
					From:    s.GetName(),
				})
			}, From: s.GetName(),
		}),
	}
}

func NewBracketsLeft(libs *tree.LibraryManager) *BracketsLeft {
	return (&BracketsLeft{})
}
