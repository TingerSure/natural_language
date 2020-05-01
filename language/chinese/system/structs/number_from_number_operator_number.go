package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
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
	*adaptor.SourceAdaptor
	operatorLeft   concept.String
	operatorRight  concept.String
	operatorResult concept.String
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
								expression.NewNewParamWithInit(map[concept.String]concept.Index{
									p.operatorLeft:  phrase[0].Index(),
									p.operatorRight: phrase[2].Index(),
								}),
							),
							p.operatorResult,
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
	page := Libs.GetLibraryPage("system", "operator")
	return (&NumberFromNumberOperatorNumber{
		SourceAdaptor:  adaptor.NewSourceAdaptor(param),
		operatorLeft:   page.GetConst(variable.NewString("Left")),
		operatorRight:  page.GetConst(variable.NewString("Right")),
		operatorResult: page.GetConst(variable.NewString("Result")),
	})
}
