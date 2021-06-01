package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	BoolFromNumberRelationalNumberName string = "structs.bool.number_relational_number"
)

var (
	BoolFromNumberRelationalNumberList []string = []string{
		phrase_type.NumberName,
		phrase_type.OperatorRelationalName,
		phrase_type.NumberName,
	}
)

type BoolFromNumberRelationalNumber struct {
	*adaptor.SourceAdaptor
}

func (p *BoolFromNumberRelationalNumber) GetStructRules() []*tree.StructRule {
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
					Size:  len(BoolFromNumberRelationalNumberList),
					Types: phrase_type.BoolName,
					From:  p.GetName(),
				})
			},
			Types: BoolFromNumberRelationalNumberList,
			From:  p.GetName(),
		}),
	}
}

func (p *BoolFromNumberRelationalNumber) GetName() string {
	return BoolFromNumberRelationalNumberName
}

func NewBoolFromNumberRelationalNumber(param *adaptor.SourceAdaptorParam) *BoolFromNumberRelationalNumber {
	return (&BoolFromNumberRelationalNumber{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
