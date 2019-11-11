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
	DivisionName string = "word.operator.division"
	divisionType int    = word_types.Operator
)

var (
	divisionCharactor = "/"

	divisionWords []*tree.Word = []*tree.Word{tree.NewWord(divisionCharactor, divisionType)}

	divisionFuncs *variable.Function = nil
)

func init() {
	divisionFuncs = variable.NewFunction(nil)
	divisionFuncs.AddParamName(phrase_types.Operator_Left)
	divisionFuncs.AddParamName(phrase_types.Operator_Right)
	divisionFuncs.Body().AddStep(
		expression.NewReturn(
			phrase_types.Operator_Result,
			expression.NewDivision(
				index.NewLocalIndex(phrase_types.Operator_Left),
				index.NewLocalIndex(phrase_types.Operator_Right),
			),
		),
	)
}

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
				return index.NewConstIndex(divisionFuncs)
			}, treasure, phrase_types.Operator, p.GetName())
		}, p.GetName()),
	}
}

func NewDivision() *Division {
	return (&Division{})
}
