package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/source/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/source/phrase_type"
	"github.com/TingerSure/natural_language/library/system/operator"
)

const (
	NumberFromNumberOperatorNumberName string = "structs.number.number_operator_number"
)

var (
	NumberFromNumberOperatorNumberList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Number,
		phrase_type.Operator,
		phrase_type.Number,
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
					Types: phrase_type.Number,
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
