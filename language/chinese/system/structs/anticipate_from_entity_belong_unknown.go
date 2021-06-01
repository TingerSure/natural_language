package structs

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
)

const (
	AnticipateFromEntityBelongUnknownName string = "structs.anticipate.entity_belong_unknown"
)

var (
	AnticipateFromEntityBelongUnknownList []string = []string{
		phrase_type.EntityName,
		phrase_type.AuxiliaryBelongName,
		phrase_type.UnknownName,
	}
)

type AnticipateFromEntityBelongUnknown struct {
	*adaptor.SourceAdaptor
	GetFieldValue   concept.String
	GetFieldKey     concept.String
	GetFieldContent concept.String
	GetField        concept.Function
}

func (p *AnticipateFromEntityBelongUnknown) GetStructRules() []*tree.StructRule {
	return []*tree.StructRule{
		tree.NewStructRule(&tree.StructRuleParam{
			Create: func() tree.Phrase {
				return tree.NewPhraseStruct(&tree.PhraseStructParam{
					Index: func(phrase []tree.Phrase) concept.Pipe {
						return p.Libs.Sandbox.Expression.ParamGet.New(
							p.Libs.Sandbox.Expression.Call.New(
								p.Libs.Sandbox.Index.ConstIndex.New(p.GetField),
								p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Pipe{
									p.GetFieldContent: phrase[0].Index(),
									p.GetFieldKey:     phrase[2].Index(),
								}),
							),
							p.GetFieldValue,
						)
					},
					Size: len(AnticipateFromEntityBelongUnknownList),
					From: p.GetName(),
				})
			},
			Types: AnticipateFromEntityBelongUnknownList,
			From:  p.GetName(),
		}),
	}
}

func (p *AnticipateFromEntityBelongUnknown) GetName() string {
	return AnticipateFromEntityBelongUnknownName
}

func NewAnticipateFromEntityBelongUnknown(param *adaptor.SourceAdaptorParam) *AnticipateFromEntityBelongUnknown {
	instance := (&AnticipateFromEntityBelongUnknown{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	})
	libObject := instance.Libs.GetLibraryPage("system", "object")

	instance.GetFieldValue = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldValue"))
	instance.GetFieldKey = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldKey"))
	instance.GetFieldContent = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("GetFieldContent"))
	instance.GetField = libObject.GetFunction(instance.Libs.Sandbox.Variable.String.New("GetField"))

	return instance
}
