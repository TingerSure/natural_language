package operator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/core/tree/phrase_types"

	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/library/operator"
)

const (
	SubtractionName string = "word.operator.subtraction"

	subtractionCharactor = "-"
)

var (
	subtractionWords []*tree.Word = []*tree.Word{tree.NewWord(subtractionCharactor)}
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
