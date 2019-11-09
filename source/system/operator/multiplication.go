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
	multiplicationName string = "system.operator.multiplication"
	multiplicationType int    = word_types.Operator
)

var (
	multiplicationCharactor = "*"

	multiplicationWords []*tree.Word = []*tree.Word{tree.NewWord(multiplicationCharactor, multiplicationType)}

	multiplicationFuncs *variable.Function = nil
)

func init() {
	multiplicationFuncs = variable.NewFunction(nil)
	multiplicationFuncs.AddParamName(phrase_types.Operator_Left)
	multiplicationFuncs.AddParamName(phrase_types.Operator_Right)
	multiplicationFuncs.Body().AddStep(
		expression.NewReturn(
			phrase_types.Operator_Result,
			expression.NewMultiplication(
				index.NewLocalIndex(phrase_types.Operator_Left),
				index.NewLocalIndex(phrase_types.Operator_Right),
			),
		),
	)
}

type Multiplication struct {
	adaptor.SourceAdaptor
}

func (p *Multiplication) GetName() string {
	return multiplicationName
}

func (p *Multiplication) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(multiplicationWords, sentence)
}

func (p *Multiplication) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return index.NewConstIndex(multiplicationFuncs)
			}, treasure, phrase_types.Operator)
		}, p.GetName()),
	}
}

func NewMultiplication() *Multiplication {
	return (&Multiplication{})
}
