package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	NumberFromNumberOperatorNumberName string = "structs.number.number_operator_number"
)

const (
	ItemLeft   = "left"
	ItemRight  = "right"
	ItemResult = "result"
)

var (
	NumberFromNumberOperatorNumberList []*tree.PhraseType = []*tree.PhraseType{
		phrase_type.Number,
		phrase_type.Operator,
		phrase_type.Number,
	}
)

type NumberFromNumberOperatorNumber struct {
	*adaptor.SourceAdaptor
}

func (p *NumberFromNumberOperatorNumber) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStructAdaptor(&tree.PhraseStructAdaptorParam{
					Index: func(phrase []tree.Phrase) concept.Index {
						return p.Libs.Sandbox.Expression.ParamGet.New(
							p.Libs.Sandbox.Expression.Call.New(
								phrase[1].Index(),
								p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Index{
									p.Libs.Sandbox.Variable.String.New(ItemLeft):  phrase[0].Index(),
									p.Libs.Sandbox.Variable.String.New(ItemRight): phrase[2].Index(),
								}),
							),
							p.Libs.Sandbox.Variable.String.New(ItemResult),
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

func NewNumberFromNumberOperatorNumber(param *adaptor.SourceAdaptorParam) *NumberFromNumberOperatorNumber {
	return (&NumberFromNumberOperatorNumber{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
