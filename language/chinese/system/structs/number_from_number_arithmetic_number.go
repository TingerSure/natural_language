package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	NumberFromNumberArithmeticNumberName string = "structs.number.number_arithmetic_number"
)

const (
	ItemLeft   = "left"
	ItemRight  = "right"
	ItemResult = "result"
)

var (
	NumberFromNumberArithmeticNumberList []string = []string{
		phrase_type.NumberName,
		phrase_type.OperatorArithmeticName,
		phrase_type.NumberName,
	}
)

type NumberFromNumberArithmeticNumber struct {
	*adaptor.SourceAdaptor
}

func (p *NumberFromNumberArithmeticNumber) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStruct(&tree.PhraseStructParam{
					Index: func(phrase []tree.Phrase) concept.Pipe {
						return p.Libs.Sandbox.Expression.ParamGet.New(
							p.Libs.Sandbox.Expression.Call.New(
								phrase[1].Index(),
								p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Pipe{
									p.Libs.Sandbox.Variable.String.New(ItemLeft):  phrase[0].Index(),
									p.Libs.Sandbox.Variable.String.New(ItemRight): phrase[2].Index(),
								}),
							),
							p.Libs.Sandbox.Variable.String.New(ItemResult),
						)
					},
					Size:  len(NumberFromNumberArithmeticNumberList),
					Types: phrase_type.NumberName,
					From:  p.GetName(),
				})
			},
			Types: NumberFromNumberArithmeticNumberList,
			From:  p.GetName(),
		}),
	}
}

func (p *NumberFromNumberArithmeticNumber) GetName() string {
	return NumberFromNumberArithmeticNumberName
}

func NewNumberFromNumberArithmeticNumber(param *adaptor.SourceAdaptorParam) *NumberFromNumberArithmeticNumber {
	return (&NumberFromNumberArithmeticNumber{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
