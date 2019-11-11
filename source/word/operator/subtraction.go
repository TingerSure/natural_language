package operator

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/sandbox/variable"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
	"github.com/TingerSure/natural_language/tree/word_types"
)

const (
	SubtractionName string = "word.operator.subtraction"
	subtractionType int    = word_types.Operator
)

var (
	subtractionCharactor = "-"

	subtractionWords []*tree.Word = []*tree.Word{tree.NewWord(subtractionCharactor, subtractionType)}

	subtractionFuncs *variable.Function = nil
)

func init() {
	subtractionFuncs = variable.NewFunction(nil)
	subtractionFuncs.AddParamName(phrase_types.Operator_Left)
	subtractionFuncs.AddParamName(phrase_types.Operator_Right)
	subtractionFuncs.Body().AddStep(
		expression.NewReturn(
			phrase_types.Operator_Result,
			expression.NewSubtraction(
				index.NewLocalIndex(phrase_types.Operator_Left),
				index.NewLocalIndex(phrase_types.Operator_Right),
			),
		),
	)
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
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return index.NewConstIndex(subtractionFuncs)
			}, treasure, phrase_types.Operator, p.GetName())
		}, p.GetName()),
	}
}

func NewSubtraction() *Subtraction {
	return (&Subtraction{})
}
