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
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewConstIndex(operator.DivisionFunc)
					},
					Content: treasure,
					Types:   phrase_types.Operator,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewDivision() *Division {
	return (&Division{})
}
