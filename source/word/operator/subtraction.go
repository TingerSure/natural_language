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
	SubtractionName      string = "word.operator.subtraction"
	subtractionType      int    = word_types.Operator
	subtractionCharactor        = "-"
)

var (
	subtractionWords []*tree.Word = []*tree.Word{tree.NewWord(subtractionCharactor, subtractionType)}
)

func init() {

}

type Subtraction struct {
	adaptor.SourceAdaptor
}

func (p *Subtraction) GetName() string {
	return SubtractionName
}

func (p *Subtraction) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(subtractionWords, sentence)
}

func (p *Subtraction) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(&tree.VocabularyRuleParam{
			Match: func(treasure *tree.Vocabulary) bool {
				return treasure.GetSource() == p
			},
			Create: func(treasure *tree.Vocabulary) tree.Phrase {
				return tree.NewPhraseVocabularyAdaptor(&tree.PhraseVocabularyAdaptorParam{
					Index: func() concept.Index {
						return index.NewConstIndex(operator.SubtractionFunc)
					},
					Content: treasure,
					Types:   phrase_types.Operator,
					From:    p.GetName(),
				})
			}, From: p.GetName(),
		}),
	}
}

func NewSubtraction() *Subtraction {
	return (&Subtraction{})
}
