package structs

import (
	"github.com/TingerSure/natural_language/library/operator"
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

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return expression.NewParamGet(
							expression.NewCall(
								phrase[1].Index(),
								expression.NewNewParamWithInit(map[string]concept.Index{
									operator.Left:  phrase[0].Index(),
									operator.Right: phrase[2].Index(),
								}),
							),
							operator.Result,
						)
					},
					Size:  len(NumberFromNumberOperatorNumberList),
					Types: phrase_types.Number,
					From:  p.GetName(),
				})
			},
			Types: NumberFromNumberOperatorNumberList,
			From:  p.GetName(),
		}),
	}
}

func (p *NumberFromNumberOperatorNumber) GetName() string {
	return NumberFromNumberOperatorNumberName
}

func NewNumberFromNumberOperatorNumber() *NumberFromNumberOperatorNumber {
	return (&NumberFromNumberOperatorNumber{})
}
