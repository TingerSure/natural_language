package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	BoolFromLogicalBoolName string = "structs.bool.bool_logical_bool"
)

var (
	BoolFromLogicalBoolList []string = []string{
		phrase_type.OperatorLogicalUnaryName,
		phrase_type.BoolName,
	}
)

type BoolFromLogicalBool struct {
	*adaptor.SourceAdaptor
}

func (p *BoolFromLogicalBool) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{

		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStruct(&tree.PhraseStructParam{
					Index: func(phrase []tree.Phrase) concept.Pipe {
						return p.Libs.Sandbox.Expression.ParamGet.New(
							p.Libs.Sandbox.Expression.Call.New(
								phrase[0].Index(),
								p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Pipe{
									p.Libs.Sandbox.Variable.String.New(ItemRight): phrase[1].Index(),
								}),
							),
							p.Libs.Sandbox.Variable.String.New(ItemResult),
						)
					},
					Size:  len(BoolFromLogicalBoolList),
					Types: phrase_type.BoolName,
					From:  p.GetName(),
				})
			},
			Types: BoolFromLogicalBoolList,
			From:  p.GetName(),
		}),
	}
}

func (p *BoolFromLogicalBool) GetName() string {
	return BoolFromLogicalBoolName
}

func NewBoolFromLogicalBool(param *adaptor.SourceAdaptorParam) *BoolFromLogicalBool {
	return (&BoolFromLogicalBool{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
