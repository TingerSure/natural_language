package operator

import (
	"github.com/TingerSure/natural_language/library/operator"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	DivisionName      string = "word.operator.division"
	divisionType      int    = word_types.Operator
	divisionCharactor        = "/"
)

var (
	divisionWords []*tree.Word = []*tree.Word{tree.NewWord(divisionCharactor, divisionType)}
)

type Division struct {
	adaptor.SourceAdaptor
}

func (p *Division) GetName() string {
	return DivisionName
}

func (p *Division) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(divisionWords, sentence)
}

func (p *Division) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return index.NewConstIndex(operator.DivisionFunc)
			}, treasure, phrase_types.Operator, p.GetName())
		}, p.GetName()),
	}
}

func NewDivision() *Division {
	return (&Division{})
}
