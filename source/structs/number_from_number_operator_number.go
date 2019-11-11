package structs

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression"
	"github.com/TingerSure/natural_language/source/adaptor"
	"github.com/TingerSure/natural_language/tree"
	"github.com/TingerSure/natural_language/tree/phrase_types"
)

const (
	NumberFromNumberOperatorNumberName string = "structs.number.number_operator_number"
)

var (
	NumberFromNumberOperatorNumberList []string = []string{
		phrase_types.Number,
		phrase_types.Operator,
		phrase_types.Number,
	}
)

type NumberFromNumberOperatorNumber struct {
	adaptor.SourceAdaptor
}

func (p *NumberFromNumberOperatorNumber) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(func() tree.Phrase {
			return tree.NewPhraseStructAdaptor(func(phrase []tree.Phrase) concept.Index {
				return expression.NewParamGet(
					expression.NewCall(
						phrase[1].Index(),
						expression.NewNewParamWithInit(map[string]concept.Index{
							phrase_types.Operator_Left:  phrase[0].Index(),
							phrase_types.Operator_Right: phrase[2].Index(),
						}),
					),
					phrase_types.Operator_Result,
				)
			}, len(NumberFromNumberOperatorNumberList), phrase_types.Number, p.GetName())
		}, NumberFromNumberOperatorNumberList, p.GetName()),
	}
}

func (p *NumberFromNumberOperatorNumber) GetName() string {
	return NumberFromNumberOperatorNumberName
}

func NewNumberFromNumberOperatorNumber() *NumberFromNumberOperatorNumber {
	return (&NumberFromNumberOperatorNumber{})
}
