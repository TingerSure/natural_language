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
	additionName string = "system.operator.addition"
	additionType int    = word_types.Operator
)

var (
	additionCharactor = "+"

	additionWords []*tree.Word = []*tree.Word{tree.NewWord(additionCharactor, additionType)}

	additionFuncs *variable.Function = nil
)

func init() {
	additionFuncs = variable.NewFunction(nil)
	additionFuncs.AddParamName(phrase_types.Operator_Left)
	additionFuncs.AddParamName(phrase_types.Operator_Right)
	additionFuncs.Body().AddStep(
		expression.NewReturn(
			phrase_types.Operator_Result,
			expression.NewAddition(
				index.NewLocalIndex(phrase_types.Operator_Left),
				index.NewLocalIndex(phrase_types.Operator_Right),
			),
		),
	)
}

type Addition struct {
	adaptor.Adaptor
}

func (p *Addition) GetName() string {
	return additionName
}

func (p *Addition) GetWords(sentence string) []*tree.Word {
	return tree.WordsFilter(additionWords, sentence)
}

func (p *Addition) GetVocabularyRules() []*tree.VocabularyRule {
	return []*tree.VocabularyRule{
		tree.NewVocabularyRule(func(treasure *tree.Vocabulary) bool {
			return treasure.GetSource() == p
		}, func(treasure *tree.Vocabulary) tree.Phrase {
			return tree.NewPhraseVocabularyAdaptor(func() concept.Index {
				return index.NewConstIndex(additionFuncs)
			}, treasure, phrase_types.Operator)
		}, p.GetName()),
	}
}

func NewAddition() *Addition {
	return (&Addition{})
}
