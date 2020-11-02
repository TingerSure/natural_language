package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	BoolFromBoolLogicalBoolName string = "structs.bool.bool_logical_bool"
)

var (
	BoolFromBoolLogicalBoolList []string = []string{
		phrase_type.BoolName,
		phrase_type.OperatorLogicalName,
		phrase_type.BoolName,
	}
)

type BoolFromBoolLogicalBool struct {
	*adaptor.SourceAdaptor
}

func (p *BoolFromBoolLogicalBool) GetStructRules() []*tree.StructRule {
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
					Size:  len(BoolFromBoolLogicalBoolList),
					Types: phrase_type.BoolName,
					From:  p.GetName(),
				})
			},
			Types: BoolFromBoolLogicalBoolList,
			From:  p.GetName(),
		}),
	}
}

func (p *BoolFromBoolLogicalBool) GetName() string {
	return BoolFromBoolLogicalBoolName
}

func NewBoolFromBoolLogicalBool(param *adaptor.SourceAdaptorParam) *BoolFromBoolLogicalBool {
	return (&BoolFromBoolLogicalBool{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
}
