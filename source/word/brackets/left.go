package brackets

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/sandbox/variable"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	LeftCharactor        = "("
	LeftName      string = "word.brackets.left"
	LeftType      int    = word_types.AuxiliaryBrackets
)

var (
	LeftWord  []*tree.Word = []*tree.Word{tree.NewWord(LeftCharactor, LeftType)}
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
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == s
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return LeftIndex
			}, treasure, phrase_types.BracketsLeft, s.GetName())
		}, s.GetName()),
	}
}

func NewBracketsLeft() *BracketsLeft {
	return (&BracketsLeft{})
}
